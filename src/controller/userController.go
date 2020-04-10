package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/robin019/bookmark/src/error"
	"github.com/robin019/bookmark/src/service"
	"net/http"
)

func CreateUser(ctx iris.Context) {
	type rule struct {
		Account string `valid:"required"`
		Gender  string `valid:"required, in(female|male)"`
	}

	params := &rule{
		Account: ctx.FormValue("Account"),
		Gender:  ctx.FormValue("Gender"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.CustomError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := service.CreateUser(params); err != nil {
		failed(ctx, err)
		return
	}
	success(ctx, "success")
}

func GetUsers(ctx iris.Context) {
	type rule struct {
		Account string `valid:"-"`
		Gender  string `valid:"in(female|male)"`
	}
	params := &rule{
		Account: ctx.URLParam("Account"),
		Gender:  ctx.URLParam("Gender"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.CustomError(http.StatusBadRequest, err.Error()))
		return
	}
	results, err := service.GetUsers(params)
	if err != nil {
		failed(ctx, err)
		return
	}
	success(ctx, results)
}

func DeleteUser(ctx iris.Context){

	type rule struct {
		Account string `valid:"-"`
		Gender  string `valid:"in(female|male)"`
	}
	params := &rule{
		Account: ctx.URLParam("Account"),
		Gender:  ctx.URLParam("Gender"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.CustomError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := service.DeleteUser(params); err != nil {
		failed(ctx, err)
		return
	}
	success(ctx, "delete success")

}
