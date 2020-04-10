package service

import (
	"fmt"
	"github.com/kataras/iris"
	gormdao "github.com/robin019/bookmark/persistance"
	bookmarkdao "github.com/robin019/bookmark/persistance/gorm/bookmarkDao"
	userdao "github.com/robin019/bookmark/persistance/gorm/userDao"
	"github.com/robin019/bookmark/src/assembler"
	"github.com/spf13/cast"
	"reflect"
	"github.com/robin019/bookmark/src/error"
)



func CreateBookmark(params interface{}) (e *error.Error) {
	value := reflect.ValueOf(params).Elem()

	tx := gormdao.DB()

	if userdao.Get(tx,&userdao.QueryModel{
		ID : cast.ToInt(value.FieldByName("UserID").String()),
	})== nil {
		return error.CustomError(iris.StatusBadRequest, fmt.Sprintf("此會員ID %v 不存在", value.FieldByName("UserID").String()))
	}

	if bookmarkdao.Get(tx, &bookmarkdao.QueryModel{
		UserID : cast.ToInt(value.FieldByName("UserID").String()),
		URL : value.FieldByName("URL").String(),

	}) != nil {

		return error.CustomError(iris.StatusBadRequest, fmt.Sprintf("此會員的書籤： %v 已存在", value.FieldByName("URL").String()))
	}

	bookmarkdao.New(tx, &bookmarkdao.Model{
		UserID: 	cast.ToInt(value.FieldByName("UserID").String()),
		Name:      	value.FieldByName("Name").String(),
		URL: 		value.FieldByName("URL").String(),
	})

	return nil
}

func GetBookmarks(params interface{}) (results []map[string]interface{}, e *error.Error) {
	value := reflect.ValueOf(params).Elem()

	tx := gormdao.DB()

	bookmarks := bookmarkdao.Gets(tx, &bookmarkdao.QueryModel{
		UserID: 	cast.ToInt(value.FieldByName("UserID").String()),
		Name:      	value.FieldByName("Name").String(),
		URL: 		value.FieldByName("URL").String(),
	})

	return assembler.Bookmark(bookmarks), nil
}

func DeleteBookmark(params interface{}) (e *error.Error) {

	value := reflect.ValueOf(params).Elem()

	tx := gormdao.DB()

	if userdao.Get(tx,&userdao.QueryModel{
		ID : cast.ToInt(value.FieldByName("UserID").String()),
	})== nil {
		return error.CustomError(iris.StatusBadRequest, fmt.Sprintf("此會員ID %v 不存在", value.FieldByName("UserID").String()))
	}

	if bookmarkdao.Get(tx, &bookmarkdao.QueryModel{
		UserID : cast.ToInt(value.FieldByName("UserID").String()),
		Name : value.FieldByName("Name").String(),
		URL : value.FieldByName("URL").String(),
	}) == nil {
		return error.CustomError(iris.StatusBadRequest, fmt.Sprintf("此會員不存在： %v 書籤", value.FieldByName("URL").String()))
	}


	bookmarkdao.Delete(tx, &bookmarkdao.QueryModel{
		UserID: 	cast.ToInt(value.FieldByName("UserID").String()),
		Name:      	value.FieldByName("Name").String(),
		URL: 		value.FieldByName("URL").String(),
	})

	return nil
}
