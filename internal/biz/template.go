package biz

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/limes-cloud/configure/pkg/util"

	"github.com/go-kratos/kratos/v2/encoding"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/biz/types"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"
)

const (
	Regexp     = `\$\{(\w|\.)+}`
	FormatJson = "json"
	FormatYaml = "yaml"
)

type Template struct {
	ktypes.BaseModel
	ServerID    uint32  `json:"server_id" gorm:"uniqueIndex:sv;not null;comment:模板id"`
	Content     string  `json:"content" gorm:"not null;type:text;comment:模板内容"`
	Version     string  `json:"version" gorm:"uniqueIndex:sv;not null;size:32;comment:模板版本"`
	IsUse       bool    `json:"is_use" gorm:"default:false;comment:是否使用"`
	Format      string  `json:"format" gorm:"not null;size:32;comment:模板格式"`
	Description string  `json:"description" gorm:"not null;size:128;comment:模板描述"`
	Compare     string  `json:"compare" gorm:"not null;type:text;comment:变更详情"`
	Server      *Server `json:"server" gorm:"constraint:onDelete:cascade"`
}

type TemplateRepo interface {
	Get(ctx kratosx.Context, id uint32) (*Template, error)
	GetByVersion(ctx kratosx.Context, version string) (*Template, error)
	Current(ctx kratosx.Context, srvId uint32) (*Template, error)
	Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*Template, uint32, error)
	All(ctx kratosx.Context, options ktypes.Scopes) ([]*Template, error)
	Create(ctx kratosx.Context, c *Template) (uint32, error)
	Update(ctx kratosx.Context, c *Template) error
	Use(ctx kratosx.Context, srvId, tpId uint32) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}

type TemplateUseCase struct {
	config *config.Config
	tpRepo TemplateRepo
	bsRepo BusinessRepo
	rsRepo ResourceRepo
}

func NewTemplateUseCase(config *config.Config, tpRepo TemplateRepo, rsRepo ResourceRepo, bsRepo BusinessRepo) *TemplateUseCase {
	return &TemplateUseCase{config: config, tpRepo: tpRepo, rsRepo: rsRepo, bsRepo: bsRepo}
}

// Current 获取当前版本的配置信息
func (t *TemplateUseCase) Current(ctx kratosx.Context, srvId uint32) (*Template, error) {
	template, err := t.tpRepo.Current(ctx, srvId)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return template, nil
}

// Get 获取指定模板信息
func (t *TemplateUseCase) Get(ctx kratosx.Context, id uint32) (*Template, error) {
	template, err := t.tpRepo.Get(ctx, id)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return template, nil
}

// Page 获取分页模板信息
func (t *TemplateUseCase) Page(ctx kratosx.Context, req *types.PageTemplateRequest) ([]*Template, uint32, error) {
	list, total, err := t.tpRepo.Page(ctx, &ktypes.PageOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			return db.Where("server_id=?", req.ServerID)
		},
	})
	if err != nil {
		return nil, 0, v1.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// Add 添加模板信息
func (t *TemplateUseCase) Add(ctx kratosx.Context, template *Template) (uint32, error) {
	// 序列化数据，判断格式是否正确
	otc := map[string]any{}
	oe := encoding.GetCodec(template.Format)
	if err := oe.Unmarshal([]byte(template.Content), &otc); err != nil {
		return 0, v1.ParseTemplateErrorFormat(err.Error())
	}

	// 当前的版本
	template.Version = t.getTemplateVersion([]byte(template.Content))

	// 查找是否存在相同的版本
	if _, err := t.tpRepo.GetByVersion(ctx, template.Version); err == nil {
		return 0, v1.VersionExistError()
	}

	// 进行模板检查是否正确
	if err := t.checkTemplate(ctx, template.ServerID, template.Content); err != nil {
		return 0, v1.CheckTemplateErrorFormat(err.Error())
	}

	// 获取当前的模板
	current, err := t.tpRepo.Current(ctx, template.ServerID)
	// 获取变更对比信息
	if err != nil {
		compare, err := t.Compare(ctx, &types.CompareTemplateRequest{
			Id:      current.ID,
			Format:  template.Format,
			Content: template.Content,
		})
		if err != nil {
			return 0, v1.ParseTemplateErrorFormat(err.Error())
		}
		compareByte, _ := json.Marshal(compare)
		template.Compare = string(compareByte)
	}

	// 数据入库
	template.IsUse = true
	id, err := t.tpRepo.Create(ctx, template)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Compare 对比变更细节
func (t *TemplateUseCase) Compare(ctx kratosx.Context, req *types.CompareTemplateRequest) ([]*types.CompareTemplateReply, error) {
	// 查询版本
	template, err := t.tpRepo.Get(ctx, req.Id)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 解析当前版本数据
	var oldKeys []string
	otc := map[string]any{}

	oe := encoding.GetCodec(template.Format)
	if err := oe.Unmarshal([]byte(template.Content), &otc); err != nil {
		return nil, v1.ParseTemplateError()
	}
	for key := range otc {
		oldKeys = append(oldKeys, key)
	}

	// 解析上传的模板
	var curKeys []string
	ctc := map[string]any{}

	ce := encoding.GetCodec(req.Format)
	if err := ce.Unmarshal([]byte(req.Content), &ctc); err != nil {
		return nil, v1.ParseTemplateError()
	}
	for key := range ctc {
		curKeys = append(curKeys, key)
	}

	var reply []*types.CompareTemplateReply

	// 新增字段
	addKeys := util.Diff(curKeys, oldKeys)
	for _, key := range addKeys {
		val := ""
		switch ctc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(ctc[key])
			val = string(b)
		default:
			val = fmt.Sprint(ctc[key])
		}
		reply = append(reply, &types.CompareTemplateReply{Type: "add", Key: key, Cur: val})
	}

	// 删除的字段
	delKeys := util.Diff(oldKeys, curKeys)
	for _, key := range delKeys {
		val := ""
		switch otc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(otc[key])
			val = string(b)
		default:
			val = fmt.Sprint(otc[key])
		}
		reply = append(reply, &types.CompareTemplateReply{Type: "del", Key: key, Old: val})
	}

	// 变更的字段
	for key, val := range ctc {
		if otc[key] == nil {
			continue
		}
		if fmt.Sprint(val) == fmt.Sprint(otc[key]) {
			continue
		}

		oldVal := ""
		switch otc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(otc[key])
			oldVal = string(b)
		default:
			oldVal = fmt.Sprint(otc[key])
		}

		curVal := ""
		switch ctc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(ctc[key])
			curVal = string(b)
		default:
			curVal = fmt.Sprint(ctc[key])
		}

		reply = append(reply, &types.CompareTemplateReply{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return reply, nil
}

// ParseByContent 指定content生成配置预览
func (t *TemplateUseCase) ParseByContent(ctx kratosx.Context, req *types.ParseByContentRequest) (string, error) {
	// 获取字段的配置值
	values, err := t.getVariableValue(ctx, req.EnvId, req.ServerId)
	if err != nil {
		return "", v1.ResourceFormatValueErrorFormat(err.Error())
	}

	// 进行值替换
	reg := regexp.MustCompile(Regexp)
	tempKeys := reg.FindAllString(req.Content, -1)
	for _, key := range tempKeys {
		if val, ok := values[key]; !ok {
			return "", v1.ParseTemplateErrorFormat("非法字段：%v", key)
		} else {
			if req.Format == FormatJson {
				req.Content = t.replaceByFormatJson(req.Content, key, val)
			} else {
				req.Content = t.replaceByFormatYaml(req.Content, key, val)
			}
		}
	}
	return req.Content, nil
}

// Parse 使用指定版本配置
func (t *TemplateUseCase) Parse(ctx kratosx.Context, req *types.ParseRequest) (*types.ParseReply, error) {
	// 获取当前模板
	tp, err := t.tpRepo.Current(ctx, req.ServerId)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := &types.ParseReply{
		Format: tp.Format,
	}

	reply.Content, err = t.ParseByContent(ctx, &types.ParseByContentRequest{
		EnvId:    req.EnvId,
		ServerId: req.ServerId,
		Format:   tp.Format,
		Content:  tp.Content,
	})

	return reply, err
}

// Switch 切换指定版本信息
func (t *TemplateUseCase) Switch(ctx kratosx.Context, srvId, tpId uint32) error {
	if err := t.tpRepo.Use(ctx, srvId, tpId); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Update 更新模板信息
func (t *TemplateUseCase) Update(ctx kratosx.Context, template *Template) error {
	if err := t.tpRepo.Update(ctx, template); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除模板信息
func (t *TemplateUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := t.tpRepo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// getTemplateVersion 获取模板的版本信息
func (t *TemplateUseCase) getTemplateVersion(body []byte) string {
	return strings.ToUpper(util.MD5(body)[:12])
}

// fillKey 获取真实的模板变量
func (t *TemplateUseCase) fillKey(val string) string {
	return fmt.Sprintf(`${%s}`, val)
}

// checkTemplate 校验数据模板数据是否合法
func (t *TemplateUseCase) checkTemplate(ctx kratosx.Context, srvId uint32, template string) error {
	// 组合模板和模板的key
	keys := map[string]bool{}

	bsList, _ := t.bsRepo.AllServerBusiness(ctx, srvId)
	for _, item := range bsList {
		keys[t.fillKey(item.Keyword)] = true
	}

	rsList, _ := t.rsRepo.AllServerResource(ctx, srvId)
	for _, item := range rsList {
		fields := strings.Split(item.Fields, ",")
		for _, val := range fields {
			keys[t.fillKey(item.Keyword+"."+val)] = true
		}
	}

	// 进行增则匹配
	reg := regexp.MustCompile(Regexp)
	tempKeys := reg.FindAllString(template, -1)
	// 进行参数判断
	for _, key := range tempKeys {
		if !keys[key] {
			return fmt.Errorf("非法字段：%v", key)
		}
	}
	return nil
}

func (t *TemplateUseCase) parseBusinessValue(tp, input string) types.TemplateValue {
	switch tp {
	case "string":
		return types.TemplateValue{Value: input, Exclude: false}
	case "object":
		return types.TemplateValue{Value: input, Exclude: true}
	default:
		return types.TemplateValue{Value: input, Exclude: true}
	}
}

func (t *TemplateUseCase) parseResourceValue(val any) types.TemplateValue {
	if val == nil {
		return types.TemplateValue{Value: "null", Exclude: true}
	}
	switch val.(type) {
	case string:
		return types.TemplateValue{Value: val.(string), Exclude: false}
	case map[string]any, []any:
		str, _ := json.Marshal(val)
		return types.TemplateValue{Value: string(str), Exclude: true}
	default:
		return types.TemplateValue{Value: fmt.Sprint(val), Exclude: true}
	}
}

func (t *TemplateUseCase) getVariableValue(ctx kratosx.Context, envId, srvId uint32) (map[string]types.TemplateValue, error) {
	// 获取全部业务
	bvs, err := t.bsRepo.GetEnvValues(ctx, envId, srvId)
	if err != nil {
		return nil, err
	}

	// 获取全部资源
	rvs, err := t.rsRepo.GetEnvValues(ctx, envId, srvId)
	if err != nil {
		return nil, err
	}

	result := make(map[string]types.TemplateValue)

	// 解析业务字段
	for _, item := range bvs {
		result[t.fillKey(item.Business.Keyword)] = t.parseBusinessValue(item.Business.Type, item.Value)
	}

	// 解析资源字段
	for _, item := range rvs {
		fs := map[string]any{}
		_ = json.Unmarshal([]byte(item.Value), &fs)

		fields := strings.Split(item.Resource.Fields, ",")
		for _, key := range fields {
			k := fmt.Sprintf("%s.%s", item.Resource.Keyword, key)
			result[t.fillKey(k)] = t.parseResourceValue(fs[key])
		}
	}

	return result, nil
}

func (t *TemplateUseCase) replaceByFormatJson(template, key string, val types.TemplateValue) string {
	isQuote := strings.Contains(template, fmt.Sprintf(`"%v"`, key)) && strings.Contains(template, fmt.Sprintf(`'%v'`, key))
	if isQuote && val.Exclude {
		template = strings.Replace(template, fmt.Sprintf(`"%v"`, key), val.Value, 1)
		template = strings.Replace(template, fmt.Sprintf(`'%v'`, key), val.Value, 1)
	} else {
		template = strings.Replace(template, key, fmt.Sprintf("%v", val.Value), 1)
	}
	return template
}

func (t *TemplateUseCase) replaceByFormatYaml(template, key string, val types.TemplateValue) string {
	return strings.Replace(template, key, fmt.Sprintf("%v", val.Value), 1)
}
