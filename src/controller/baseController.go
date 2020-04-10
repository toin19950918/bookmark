package controller

import (
	"regexp"
	"github.com/spf13/cast"
	"time"
	"github.com/asaskevich/govalidator"
	"github.com/robin019/bookmark/src/error"
	"github.com/kataras/iris"
)

// custome validator
func init() {
	// validate minimum
	govalidator.ParamTagMap["min"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		min := cast.ToFloat64(params[0])
		number := cast.ToFloat64(str)
		return number >= min
	})
	govalidator.ParamTagRegexMap["min"] = regexp.MustCompile("^min\\((\\w+)\\)$")

	// validate maximum
	govalidator.ParamTagMap["max"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		max := cast.ToFloat64(params[0])
		number := cast.ToFloat64(str)
		return number <= max
	})
	govalidator.ParamTagRegexMap["max"] = regexp.MustCompile("^max\\((\\w+)\\)$")

	// validate timestamp
	govalidator.CustomTypeTagMap.Set("timestamp", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		const timeStamp = "2006-01-02 15:04:05"
		_, err := time.Parse(timeStamp, i.(string))
		if err != nil {
			return false
		}
		return true
	}))

	govalidator.TagMap["no_special_chars"] = govalidator.Validator(func(str string) bool {
		return govalidator.Matches(str, "^[0-9a-zA-Z]+$")
	})

	govalidator.TagMap["product_list"] = govalidator.Validator(func(str string) bool {
		return govalidator.Matches(str, `^[0-9a-zA-Z\,]+$`)
	})
}

func success(ctx iris.Context, data interface{}) {
	ctx.StatusCode(200)
	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

func failed(ctx iris.Context, err *error.Error) {
	ctx.StatusCode(err.Code())
	ctx.JSON(iris.Map{
		"code":    err.Code(),
		"message": err.Error(),
		"data":    []string{},
	})
}
