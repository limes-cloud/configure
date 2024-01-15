package logic

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding"
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

type Template struct {
	conf *config.Config
}

func NewTemplate(conf *config.Config) *Template {
	return &Template{
		conf: conf,
	}
}

type value struct {
	value   string
	exclude bool
}

const (
	Regexp     = `\$\{(\w|\.)+}`
	FormatJson = "json"
	FormatYaml = "yaml"
)

// Current 查询指定服务的当前模板
func (t *Template) Current(ctx kratosx.Context, in *v1.CurrentTemplateRequest) (*v1.CurrentTemplateReply, error) {
	template := model.Template{}
	reply := v1.CurrentTemplateReply{}

	if template.Current(ctx, in.ServerId) != nil {
		return &reply, nil
	}
	if util.Transform(template, &reply.Template) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

// Get 查询指定模板
func (t *Template) Get(ctx kratosx.Context, in *v1.GetTemplateRequest) (*v1.GetTemplateReply, error) {
	template := model.Template{}
	if err := template.OneById(ctx, in.Id); err != nil {
		return nil, v1.DatabaseError()
	}
	reply := v1.GetTemplateReply{}
	if util.Transform(template, &reply) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

// Page 查询分页模板
func (t *Template) Page(ctx kratosx.Context, in *v1.PageTemplateRequest) (*v1.PageTemplateReply, error) {
	template := model.Template{}
	list, total, err := template.Page(ctx, &types.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			return db.Where("server_id = ?", in.ServerId)
		},
	})

	if err != nil {
		return nil, v1.DatabaseError()
	}
	reply := v1.PageTemplateReply{
		Total: total,
	}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

// Add 添加模板
func (t *Template) Add(ctx kratosx.Context, in *v1.AddTemplateRequest) (*emptypb.Empty, error) {
	template := model.Template{}
	if err := util.Transform(in, &template); err != nil {
		return nil, v1.TransformError()
	}

	otc := map[string]any{}
	oe := encoding.GetCodec(template.Format)
	if err := oe.Unmarshal([]byte(template.Content), &otc); err != nil {
		return nil, v1.ParseTemplateErrorFormat(err.Error())
	}
	// 查找当前的版本
	template.Version = strings.ToUpper(util.MD5([]byte(in.Content))[:12])
	temp := model.Template{}
	if temp.Current(ctx, in.ServerId) == nil {
		if template.Version == temp.Version {
			return nil, v1.VersionExistError()
		}
	}

	// 进行模板检查
	if err := t.checkTemplate(ctx, in.ServerId, in.Content); err != nil {
		return nil, v1.CheckTemplateErrorFormat(err.Error())
	}

	// 进行变更对比
	compare, err := t.Compare(ctx, &v1.CompareTemplateRequest{Id: temp.ID, Format: in.Format, Content: in.Content})
	if err != nil {
		return nil, v1.ParseTemplateErrorFormat(err.Error())
	}
	compareByte, _ := json.Marshal(compare.List)
	template.Compare = string(compareByte)
	template.IsUse = true
	return nil, template.Create(ctx)
}

// Compare 对比变更细节
func (t *Template) Compare(ctx kratosx.Context, in *v1.CompareTemplateRequest) (*v1.CompareTemplateReply, error) {
	// 查询当前版本
	template := model.Template{}
	if err := template.OneById(ctx, in.Id); err != nil {
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

	ce := encoding.GetCodec(in.Format)
	if err := ce.Unmarshal([]byte(in.Content), &ctc); err != nil {
		return nil, v1.ParseTemplateError()
	}
	for key := range ctc {
		curKeys = append(curKeys, key)
	}

	reply := v1.CompareTemplateReply{
		List: make([]*v1.CompareTemplateInfo, 0),
	}

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
		reply.List = append(reply.List, &v1.CompareTemplateInfo{Type: "add", Key: key, Cur: val})
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
		reply.List = append(reply.List, &v1.CompareTemplateInfo{Type: "del", Key: key, Old: val})
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

		reply.List = append(reply.List, &v1.CompareTemplateInfo{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return &reply, nil
}

// Switch 使用指定版本
func (t *Template) Switch(ctx kratosx.Context, in *v1.SwitchTemplateRequest) (*emptypb.Empty, error) {
	template := model.Template{}
	if util.Transform(in, &template) != nil {
		return nil, v1.TransformError()
	}

	return nil, template.UseVersionByID(ctx, in.ServerId, in.Id)
}

// ParsePreview 配置解决预览
func (t *Template) ParsePreview(ctx kratosx.Context, in *v1.ParseTemplatePreviewRequest) (*v1.ParseTemplatePreviewReply, error) {
	env := model.Environment{}
	if err := env.OneByKeyword(ctx, in.EnvKeyword); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取字段的配置值
	values, err := t.getVariableValue(ctx, env.ID, in.ServerId)
	if err != nil {
		return nil, v1.ResourceFormatValueErrorFormat(err.Error())
	}

	// 进行值替换
	reg := regexp.MustCompile(Regexp)
	tempKeys := reg.FindAllString(in.Content, -1)
	for _, key := range tempKeys {
		if val, ok := values[key]; !ok {
			return nil, fmt.Errorf("非法字段：%v", key)
		} else {
			if in.Format == FormatJson {
				in.Content = t.replaceByFormatJson(in.Content, key, val)
			} else {
				in.Content = t.replaceByFormatYaml(in.Content, key, val)
			}
		}
	}

	return &v1.ParseTemplatePreviewReply{
		Content: in.Content,
	}, nil
}

// Parse 使用指定版本配置
func (t *Template) Parse(ctx kratosx.Context, in *v1.ParseTemplateRequest) (*v1.ParseTemplateReply, error) {
	env := model.Environment{}
	if err := env.OneByKeyword(ctx, in.EnvKeyword); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取指定模板
	tp := model.Template{}
	if err := tp.Current(ctx, in.ServerId); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取字段的配置值
	values, err := t.getVariableValue(ctx, env.ID, tp.ServerID)
	if err != nil {
		return nil, v1.ResourceFormatValueErrorFormat(err.Error())
	}

	// 进行值替换
	reg := regexp.MustCompile(Regexp)
	tempKeys := reg.FindAllString(tp.Content, -1)
	for _, key := range tempKeys {
		if val, ok := values[key]; !ok {
			return nil, fmt.Errorf("非法字段：%v", key)
		} else {
			if tp.Format == FormatJson {
				tp.Content = t.replaceByFormatJson(tp.Content, key, val)
			} else {
				tp.Content = t.replaceByFormatYaml(tp.Content, key, val)
			}
		}
	}

	return &v1.ParseTemplateReply{
		Content: tp.Content,
		Format:  tp.Format,
	}, nil
}

// checkTemplate 校验数据模板数据是否合法
func (t *Template) checkTemplate(ctx kratosx.Context, srvId uint32, template string) error {
	// 获取指定服务的模板字段
	bs := model.Business{}
	bsList, err := bs.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("server_id=?", srvId)
	})
	if err != nil {
		return v1.DatabaseError()
	}

	// 获取指定服务的模板字段
	rs := model.ResourceServer{}
	rsList, _ := rs.All(ctx, func(db *gorm.DB) *gorm.DB {
		db.Preload("Resource")
		return db.Where("server_id=?", srvId)
	})

	// 获取公共的资源变量
	prs := model.Resource{}
	prsList, _ := prs.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("private=false")
	})

	// 组合模板和模板的key
	keys := map[string]bool{}
	for _, item := range rsList {
		fields := strings.Split(item.Resource.Fields, ",")
		for _, val := range fields {
			keys[t.fillKey(item.Resource.Keyword+"."+val)] = true
		}
	}

	for _, item := range prsList {
		fields := strings.Split(item.Fields, ",")
		for _, val := range fields {
			keys[t.fillKey(item.Keyword+"."+val)] = true
		}
	}

	for _, item := range bsList {
		keys[t.fillKey(item.Keyword)] = true
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

func (t *Template) fillKey(val string) string {
	return fmt.Sprintf(`${%s}`, val)
}

func (t *Template) getVariableValue(ctx kratosx.Context, envId, srvId uint32) (map[string]value, error) {
	// 获取全部业务
	bv := model.BusinessValue{}
	bvs, err := bv.AllByEnvAndServer(ctx, envId, srvId)
	if err != nil {
		return nil, err
	}

	// 获取全部资源
	rv := model.ResourceValue{}
	rvs, err := rv.AllByEnvAndServer(ctx, envId, srvId)
	if err != nil {
		return nil, err
	}

	result := make(map[string]value)

	// 解析业务字段
	for _, item := range bvs {
		result[t.fillKey(item.Business.Keyword)] = t.parseBusinessValue(item.Business.Type, item.Value)
	}

	// 解析资源字段
	for _, item := range rvs {
		fs := map[string]any{}
		_ = json.Unmarshal([]byte(item.Values), &fs)

		fields := strings.Split(item.Resource.Fields, ",")
		for _, key := range fields {
			k := fmt.Sprintf("%s.%s", item.Resource.Keyword, key)
			result[t.fillKey(k)] = t.parseResourceValue(fs[key])
		}
	}

	return result, nil
}

func (t *Template) parseBusinessValue(tp, input string) value {
	switch tp {
	case "string":
		return value{value: input, exclude: false}
	case "object":
		return value{value: input, exclude: true}
	default:
		return value{value: input, exclude: true}
	}
}

func (t *Template) parseResourceValue(val any) value {
	if val == nil {
		return value{value: "null", exclude: true}
	}
	switch val.(type) {
	case string:
		return value{value: val.(string), exclude: false}
	case map[string]any, []any:
		str, _ := json.MarshalToString(val)
		return value{value: str, exclude: true}
	default:
		return value{value: fmt.Sprint(val), exclude: true}
	}
}

func (t *Template) replaceByFormatJson(template, key string, val value) string {
	isQuote := strings.Contains(template, fmt.Sprintf(`"%v"`, key)) && strings.Contains(template, fmt.Sprintf(`'%v'`, key))
	if isQuote && val.exclude {
		template = strings.Replace(template, fmt.Sprintf(`"%v"`, key), val.value, 1)
		template = strings.Replace(template, fmt.Sprintf(`'%v'`, key), val.value, 1)
	} else {
		template = strings.Replace(template, key, fmt.Sprintf("%v", val.value), 1)
	}
	return template
}

func (t *Template) replaceByFormatYaml(template, key string, val value) string {
	return strings.Replace(template, key, fmt.Sprintf("%v", val.value), 1)
}
