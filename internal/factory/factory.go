package factory

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/biz/business"
	"github.com/limes-cloud/configure/internal/biz/resource"
)

const (
	Regexp     = `\$\{(\w|\.)+}`
	FormatJson = "json"
	FormatYaml = "yaml"
)

type factory struct {
	bsRepo business.Repo
	rsRepo resource.Repo
}

type Factory interface {
	TemplateVersion(body []byte) string
	CheckTemplate(ctx kratosx.Context, srvId uint32, template string) error
	ParseByContent(ctx kratosx.Context, req *ParseByContentRequest) (string, error)
}

func New(bsRepo business.Repo, rsRepo resource.Repo) Factory {
	return &factory{
		bsRepo: bsRepo,
		rsRepo: rsRepo,
	}
}

// TemplateVersion 获取模板的版本信息
func (f *factory) TemplateVersion(body []byte) string {
	return util.MD5ToUpper(body)[:12]
}

// fillKey 获取真实的模板变量
func (f *factory) fillKey(val string) string {
	return fmt.Sprintf(`${%s}`, val)
}

// CheckTemplate 校验数据模板数据是否合法
func (f *factory) CheckTemplate(ctx kratosx.Context, srvId uint32, template string) error {
	// 组合模板和模板的key
	keys := map[string]bool{}

	bsKeys, _ := f.bsRepo.AllBusinessField(ctx, srvId)
	for _, key := range bsKeys {
		keys[f.fillKey(key)] = true
	}

	rsKeys, _ := f.rsRepo.AllResourceField(ctx, srvId)
	for _, key := range rsKeys {
		keys[f.fillKey(key)] = true
	}

	// 进行增则匹配
	reg := regexp.MustCompile(Regexp)
	tempKeys := reg.FindAllString(template, -1)
	// 进行参数判断
	for _, key := range tempKeys {
		if !keys[key] {
			return errors.CheckTemplateErrorFormat("非法字段：%v", key)
		}
	}
	return nil
}

func (f *factory) parseBusinessValue(tp, input string) TemplateValue {
	switch tp {
	case "string":
		return TemplateValue{Value: input, Exclude: false}
	case "object":
		return TemplateValue{Value: input, Exclude: true}
	default:
		return TemplateValue{Value: input, Exclude: true}
	}
}

func (f *factory) parseResourceValue(val any) TemplateValue {
	if val == nil {
		return TemplateValue{Value: "null", Exclude: true}
	}
	switch val.(type) {
	case string:
		return TemplateValue{Value: val.(string), Exclude: false}
	case map[string]any, []any:
		str, _ := json.Marshal(val)
		return TemplateValue{Value: string(str), Exclude: true}
	default:
		return TemplateValue{Value: fmt.Sprint(val), Exclude: true}
	}
}

func (f *factory) getVariableValue(ctx kratosx.Context, envId, srvId uint32) (map[string]TemplateValue, error) {
	// 获取全部业务
	bvs, err := f.bsRepo.AllBusinessValue(ctx, envId, srvId)
	if err != nil {
		return nil, err
	}

	// 获取全部资源
	rvs, err := f.rsRepo.AllResourceValue(ctx, envId, srvId)
	if err != nil {
		return nil, err
	}

	result := make(map[string]TemplateValue)

	// 解析业务字段
	for _, item := range bvs {
		result[f.fillKey(item.Business.Keyword)] = f.parseBusinessValue(item.Business.Type, item.Value)
	}

	// 解析资源字段
	for _, item := range rvs {
		fs := map[string]any{}
		_ = json.Unmarshal([]byte(item.Value), &fs)

		fields := strings.Split(item.Resource.Fields, ",")
		for _, key := range fields {
			k := fmt.Sprintf("%s.%s", item.Resource.Keyword, key)
			result[f.fillKey(k)] = f.parseResourceValue(fs[key])
		}
	}

	return result, nil
}

func (f *factory) replaceByFormatJson(template, key string, val TemplateValue) string {
	isQuote := strings.Contains(template, fmt.Sprintf(`"%v"`, key)) && strings.Contains(template, fmt.Sprintf(`'%v'`, key))
	if isQuote && val.Exclude {
		template = strings.Replace(template, fmt.Sprintf(`"%v"`, key), val.Value, 1)
		template = strings.Replace(template, fmt.Sprintf(`'%v'`, key), val.Value, 1)
	} else {
		template = strings.Replace(template, key, fmt.Sprintf("%v", val.Value), 1)
	}
	return template
}

func (f *factory) replaceByFormatYaml(template, key string, val TemplateValue) string {
	return strings.Replace(template, key, fmt.Sprintf("%v", val.Value), 1)
}

// ParseByContent 指定content生成配置预览
func (f *factory) ParseByContent(ctx kratosx.Context, req *ParseByContentRequest) (string, error) {
	// 获取字段的配置值
	values, err := f.getVariableValue(ctx, req.EnvId, req.ServerId)
	if err != nil {
		return "", errors.ResourceFormatValueErrorFormat(err.Error())
	}

	// 进行值替换
	reg := regexp.MustCompile(Regexp)
	tempKeys := reg.FindAllString(req.Content, -1)
	for _, key := range tempKeys {
		if val, ok := values[key]; !ok {
			return "", errors.ParseTemplateErrorFormat("非法字段：%v", key)
		} else {
			if req.Format == FormatJson {
				req.Content = f.replaceByFormatJson(req.Content, key, val)
			} else {
				req.Content = f.replaceByFormatYaml(req.Content, key, val)
			}
		}
	}
	return req.Content, nil
}
