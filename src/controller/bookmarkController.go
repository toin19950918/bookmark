package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/robin019/bookmark/src/error"
	"github.com/robin019/bookmark/src/service"
	"net/http"
)

func CreateBookmark(ctx iris.Context) {
	type rule struct {
		UserID 	string `valid:"required"`
		Name  	string `valid:"required"`
		URL 	string `valid:"required"`
	}

	params := &rule{
		UserID: ctx.PostValue("UserID"),
		Name:  ctx.PostValue("Name"),
		URL:  ctx.PostValue("URL"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.CustomError(http.StatusBadRequest, err.Error()))
		return
	}



	if err := service.CreateBookmark(params); err != nil {
		failed(ctx, err)
		return
	}

	success(ctx, "success")
}


func GetBookmarks(ctx iris.Context) {

	type rule struct {
		UserID 	string
		Name  	string
		URL 	string
	}

	params := &rule{
		UserID: ctx.URLParam("UserID"),
		Name: ctx.URLParam("Name"),
		URL:  ctx.URLParam("URL"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.CustomError(http.StatusBadRequest, err.Error()))
		return
	}

	results, err := service.GetBookmarks(params)
	if err != nil {
		failed(ctx, err)
		return
	}
	success(ctx, results)
}

func DeleteBookmark(ctx iris.Context){

	type rule struct {
		UserID 	string `valid:"required"`
		Name  	string `valid:"-"`
		URL 	string `valid:"-"`
	}
	params := &rule{
		UserID: ctx.FormValue("UserID"),
		Name:  ctx.FormValue("Name"),
		URL:  ctx.FormValue("URL"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.CustomError(http.StatusBadRequest, err.Error()))
		return
	}

	if params.Name == "" && params.URL == ""{

		failed(ctx,error.CustomError(iris.StatusBadRequest, "url or name 必須填寫"))
		return
	}

	if err := service.DeleteBookmark(params); err != nil {
		failed(ctx, err)
		return
	}
	success(ctx, "delete success")

}