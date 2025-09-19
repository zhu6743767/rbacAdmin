package validate

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	// 初始化翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			label = field.Name
		}
		name := field.Tag.Get("json")
		return fmt.Sprintf("%s---%s", name, label)
	})
}

type ValidateErrorResponse struct {
	FiledMap map[string]any `json:"filed_map"`
	Msg      string         `json:"msg"`
}

func ValidateError(err error) (resp ValidateErrorResponse) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		resp.Msg = err.Error()
		return
	}
	m := make(map[string]any)
	var msgList []string
	for _, e := range errs {
		msg := e.Translate(trans)
		_list := strings.Split(msg, "---")
		m[_list[0]] = _list[1]
		msgList = append(msgList, _list[1])
	}
	resp.FiledMap = m
	resp.Msg = strings.Join(msgList, ";")
	return resp
}
