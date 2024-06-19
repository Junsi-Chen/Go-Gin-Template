package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
	"template/tool/log"
)

var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans() (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() // 中文翻译器
		uni := ut.New(zhT, zhT)
		ok = false
		trans, ok = uni.GetTranslator("zh")
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", "zh")
		}
		// 注册翻译器
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
		return
	}
	return
}

// ValidParams 验证参数是否合法
func ValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindBodyWith(params, binding.JSON); err != nil {
		// 获取validator.ValidationErrors类型的errors
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			log.Logger.Error(err.Error())
			return err
		}
		var sliceErrs []string
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ";"))
	}
	return nil
}
