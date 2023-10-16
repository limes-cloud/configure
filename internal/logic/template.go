package logic

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/limes-cloud/configure/pkg/md"

	json "github.com/json-iterator/go"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type Template struct {
	conf *config.Config
}

func NewTemplate(conf *config.Config) *Template {
	return &Template{
		conf: conf,
	}
}

// Current 查询指定服务的当前模板
func (l *Template) Current(ctx kratos.Context, in *v1.CurrentTemplateRequest) (*v1.CurrentTemplateReply, error) {
	template := model.Template{}
	if err := template.Current(ctx, in.ServerId); err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.CurrentTemplateReply{}
	if util.Transform(template, &reply) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Get 查询指定模板
func (l *Template) Get(ctx kratos.Context, in *v1.GetTemplateRequest) (*v1.GetTemplateReply, error) {
	template := model.Template{}
	if err := template.OneById(ctx, in.Id); err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.GetTemplateReply{}
	if util.Transform(template, &reply) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Page 查询分页模板
func (l *Template) Page(ctx kratos.Context, in *v1.PageTemplateRequest) (*v1.PageTemplateReply, error) {
	template := model.Template{}
	list, total, err := template.Page(ctx, &model.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			return db.Where("server_id = ?", in.ServerId)
		},
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.PageTemplateReply{
		Total: total,
	}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Add 添加模板
func (l *Template) Add(ctx kratos.Context, in *v1.AddTemplateRequest) (*emptypb.Empty, error) {
	template := model.Template{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if err := util.Transform(in, &template); err != nil {
		return nil, v1.ErrorTransform()
	}
	template.Version = util.MD5([]byte(in.Content))

	// 查找当前的版本
	temp := model.Template{}
	if temp.Current(ctx, in.ServerId) == nil {
		// 判断两个版本之前是否有差异
		if template.Version == temp.Version {
			return nil, v1.ErrorVersionExist()
		}
	}

	if err := l.checkTemplate(ctx, in.ServerId, in.Content); err != nil {
		return nil, v1.ErrorCheckTemplateFormat(err.Error())
	}
	template.IsUse = true
	return nil, template.Create(ctx)
}

// UseVersion 使用指定版本
func (l *Template) UseVersion(ctx kratos.Context, in *v1.UseTemplateVersionRequest) (*emptypb.Empty, error) {
	template := model.Template{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if util.Transform(in, &template) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, template.UseVersionByID(ctx, in.ServerId, in.Id)
}

// Parse 使用指定版本
func (l *Template) Parse(ctx kratos.Context, in *v1.ParseTemplateRequest) (*v1.ParseTemplateReply, error) {
	// 获取指定服务
	server := model.Server{}
	if err := server.OneByKeyword(ctx, in.SrvKeyword); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}
	// 获取指定模板
	template := model.Template{}
	if err := template.Current(ctx, server.ID); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 获取环境ID
	env := model.Environment{}
	if err := env.OneByKeyword(ctx, in.EnvKeyword); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 获取字段的配置值
	values, err := l.getVariableValue(ctx, env.ID, template.ServerID)
	if err != nil {
		return nil, v1.ErrorResourceFormatValueFormat(err.Error())
	}

	// 进行值替换
	reg := regexp.MustCompile(`"\{\{(\w|\.)+}}"`)
	tempKeys := reg.FindAllString(template.Content, -1)
	for _, key := range tempKeys {
		if val, ok := values[key]; ok {
			template.Content = strings.Replace(template.Content, key, val, 1)
		} else {
			template.Content = strings.Replace(template.Content, key, "null", 1)
		}
	}

	return &v1.ParseTemplateReply{
		Content: template.Content,
	}, nil
}

func (l *Template) fillKey(val string) string {
	return fmt.Sprintf(`"{{%s}}"`, val)
}

// checkTemplate 校验数据模板数据是否合法
func (l *Template) checkTemplate(ctx kratos.Context, srvId int64, template string) error {
	//获取指定服务的模板字段
	bs := model.Business{}
	bsList, err := bs.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("server_id = ?", srvId)
	})
	if err != nil {
		return v1.ErrorDatabase()
	}

	// 获取指定服务的模板字段
	rs := model.ResourceServer{}
	rsList, _ := rs.All(ctx, func(db *gorm.DB) *gorm.DB {
		db.Preload("Resource")
		return db.Where("server_id =?", srvId)
	})

	//组合模板和模板的key
	keys := map[string]bool{}
	for _, item := range rsList {
		var fields []string
		_ = json.Unmarshal([]byte(item.Resource.Fields), &fields)
		for _, val := range fields {
			keys[l.fillKey(item.Resource.Keyword+"."+val)] = true
		}
	}
	for _, item := range bsList {
		keys[l.fillKey(item.Keyword)] = true
	}

	//进行增则匹配
	reg := regexp.MustCompile(`"\{\{(\w|\.)+}}"`)
	tempKeys := reg.FindAllString(template, -1)
	// 进行参数判断
	for _, key := range tempKeys {
		if !keys[key] {
			return fmt.Errorf("非法字段：%v", key)
		}
	}
	return nil
}

func (l *Template) getVariableValue(ctx kratos.Context, envId, srvId int64) (map[string]string, error) {
	// 查找模板
	bv := model.BusinessValue{}
	bvs, err := bv.AllByEnvAndServer(ctx, envId, srvId)
	if err != nil {
		return nil, err
	}

	rv := model.ResourceValue{}
	rvs, err := rv.AllByEnvAndServer(ctx, envId, srvId)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)

	for _, item := range bvs {
		result[l.fillKey(item.Business.Keyword)] = l.parseBusinessValue(item.Value)
	}

	for _, item := range rvs {
		fs := map[string]any{}
		_ = json.Unmarshal([]byte(item.Values), &fs)
		for key, val := range fs {
			k := fmt.Sprintf("%v.%v", item.Resource.Keyword, key)
			result[l.fillKey(k)] = l.parseResourceValue(val)
		}
	}

	return result, nil
}

func (l *Template) parseBusinessValue(input string) string {
	input = strings.ReplaceAll(input, `\"`, `"`)
	input = strings.Trim(input, " ")

	// 以"*"格式的默认为字符串
	if input[0] == '"' && input[len(input)-1] == '"' {
		return input
	}

	// 不为数字，也不为json
	if _, err := strconv.Atoi(input); err != nil && input[0] != '{' {
		input = fmt.Sprintf(`"%s"`, input)
	}

	return input
}

func (l *Template) parseResourceValue(val any) string {
	switch val.(type) {
	//case int64, int32, int16, int8, uint64, uint32, uint16, uint8, int, float32, float64, bool:
	//	return fmt.Sprint(val)
	case string:
		return fmt.Sprintf(`"%s"`, val)
	case map[string]interface{}, []interface{}:
		str, _ := json.MarshalToString(val)
		return str
	default:
		return fmt.Sprint(val)
	}
}
