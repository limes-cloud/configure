package service

import (
	"regexp"
	"sort"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/configure/internal/types"
)

// templateKeyRegex 模板占位符（与 internal/pkg/template.go 保持一致；
// 注意使用 [\w.]+ 字符类而非 (\w|\.)+ 重复组——后者在 FindAllStringSubmatch 中只保留最后一次迭代内容）
var templateKeyRegex = regexp.MustCompile(`\$\{([\w.]+)}`)

// Settings 系统配置服务：将当前模板引用的业务变量解析为可视化配置，保存即重新渲染并广播
type Settings struct {
	business  repository.Business // 业务变量定义与值
	template  repository.Template // 当前模板（解析引用的变量）
	configure *Configure          // 配置渲染/保存/广播（组合复用）
}

func NewSettings(
	business repository.Business,
	template repository.Template,
	configure *Configure,
) *Settings {
	return &Settings{
		business:  business,
		template:  template,
		configure: configure,
	}
}

// GetSettings 获取系统配置：当前模板引用的业务变量定义 + 指定环境的值
func (s *Settings) GetSettings(ctx core.Context, app string, envId uint32) (*types.GetSettingsReply, error) {
	// 1. 当前生效模板（多命名空间），解析引用的业务变量 key（剔除 ${resource.xxx} 资源变量）
	templates, err := s.template.ListCurrentTemplate(ctx, app)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	keys := s.templateBusinessKeys(templates)

	// 2. 业务变量定义（app 下全部，按模板引用过滤）
	businessList, _, err := s.business.ListBusiness(ctx, &types.ListBusinessRequest{
		Page:     1,
		PageSize: 10000,
		App:      app,
	})
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	businessMap := map[string]*entity.Business{}
	for _, item := range businessList {
		if _, ok := keys[item.Keyword]; ok {
			businessMap[item.Keyword] = item
		}
	}
	ctx.Logger().Warnw("msg", "settings debug", "app", app, "envId", envId, "templates", len(templates), "keys", len(keys), "businessList", len(businessList), "businessMap", len(businessMap))

	// 3. 当前值（按 app + env 查）
	values, err := s.business.ListBusinessValue(ctx, &types.ListBusinessValueRequest{
		App:   &app,
		EnvId: &envId,
	})
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	valueMap := map[uint32]string{}
	for _, item := range values {
		valueMap[item.BusinessId] = item.Value
	}

	// 4. 组装（保持模板解析顺序确定性；未设置值的变量 value 为空，由前端回退 extra.default）
	reply := &types.GetSettingsReply{App: app, EnvId: envId}
	if len(templates) > 0 {
		reply.Format = templates[0].Format
	}
	for _, keyword := range sortedKeys(businessMap) {
		business := businessMap[keyword]
		description := ""
		if business.Description != nil {
			description = *business.Description
		}
		reply.Items = append(reply.Items, &types.SettingsItem{
			Keyword:     business.Keyword,
			Type:        business.Type,
			Description: description,
			Extra:       business.Extra,
			Value:       valueMap[business.Id],
		})
	}
	return reply, nil
}

// UpdateSettings 保存系统配置：类型校验 + 落库 + 重新渲染 + 广播（保存即生效）
func (s *Settings) UpdateSettings(ctx core.Context, app string, envId uint32, items []*types.UpdateSettingsItem) error {
	// 1. 业务变量定义（app 下全部），按 keyword 索引用于校验与回填
	businessList, _, err := s.business.ListBusiness(ctx, &types.ListBusinessRequest{
		Page:     1,
		PageSize: 10000,
		App:      app,
	})
	if err != nil {
		return errors.GetError(err.Error())
	}
	businessMap := map[string]*entity.Business{}
	for _, item := range businessList {
		businessMap[item.Keyword] = item
	}

	// 2. 类型校验并 upsert 值（业务变量固定必填：空值会导致模板渲染为残缺配置）
	var values []*entity.BusinessValue
	for _, item := range items {
		business, ok := businessMap[item.Keyword]
		if !ok {
			return errors.ParamsError("业务变量不存在:" + item.Keyword)
		}
		if item.Value == "" {
			return errors.ParamsError(business.Keyword + "为必填项")
		}
		bv := &entity.BusinessValue{BusinessId: business.Id, EnvId: envId, Value: item.Value}
		v, err := bv.MarshalValue(business.Type, business.Extra)
		if err != nil {
			return err
		}
		bv.Value = v
		values = append(values, bv)
	}
	if err := s.business.UpdateBusinessValues(ctx, values); err != nil {
		return errors.UpdateError(err.Error())
	}

	// 3. 重新渲染 + 落库 + 广播（复用配置同步链路）；内容未变化（改回原值）视为已同步
	if err := s.configure.UpdateConfigure(ctx, &entity.Configure{App: app, EnvId: envId}); err != nil {
		if errors.IsConfigureVersionExistError(err) {
			return nil
		}
		return err
	}
	return nil
}

// templateBusinessKeys 解析所有模板内容引用的业务变量 key（剔除含 . 的资源变量，如 ${resource.host}）
func (s *Settings) templateBusinessKeys(templates []*entity.Template) map[string]bool {
	keys := map[string]bool{}
	for _, template := range templates {
		if template == nil {
			continue
		}
		for _, match := range templateKeyRegex.FindAllStringSubmatch(template.Content, -1) {
			if len(match) < 2 {
				continue
			}
			key := match[1]
			// 资源变量格式为 resource.xxx（含 .），与业务变量区分
			if _, ok := keys[key]; !ok && !containsDot(key) {
				keys[key] = true
			}
		}
	}
	return keys
}

// containsDot 判断字符串是否包含 .（资源变量特征）
func containsDot(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return true
		}
	}
	return false
}

// sortedKeys 按 key 字典序输出（确定性排序，保证多次请求返回顺序一致）
func sortedKeys(m map[string]*entity.Business) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}