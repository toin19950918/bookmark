package service

import (
	"fmt"
	"github.com/kataras/iris"
	gormdao "github.com/robin019/bookmark/persistance"
	userdao "github.com/robin019/bookmark/persistance/gorm/userDao"
	"github.com/robin019/bookmark/src/assembler"
	"github.com/robin019/bookmark/src/error"
	"reflect"
)

func CreateUser(params interface{}) (e *error.Error) {
	value := reflect.ValueOf(params).Elem()

	tx := gormdao.DB()

	if userdao.Get(tx, &userdao.QueryModel{
		UserAccount: value.FieldByName("Account").String(),
	}) != nil {
		return error.CustomError(iris.StatusBadRequest, fmt.Sprintf("會員帳號 %v 已存在", value.FieldByName("Account").String()))
	}

	userdao.New(tx, &userdao.Model{
		UserAccount: value.FieldByName("Account").String(),
		Gender:      value.FieldByName("Gender").String(),
	})

	return nil
}

func GetUsers(params interface{}) (results []map[string]interface{}, e *error.Error) {
	value := reflect.ValueOf(params).Elem()

	tx := gormdao.DB()

	users := userdao.Gets(tx, &userdao.QueryModel{
		UserAccount: value.FieldByName("Account").String(),
		Gender:      value.FieldByName("Gender").String(),
	})

	return assembler.User(users), nil
}

func DeleteUser(params interface{}) (e *error.Error) {

	value := reflect.ValueOf(params).Elem()
	tx := gormdao.DB()

	if userdao.Get(tx, &userdao.QueryModel{
		UserAccount: value.FieldByName("Account").String(),
	}) == nil {
		return error.CustomError(iris.StatusBadRequest, fmt.Sprintf("會員帳號 %v 不存在", value.FieldByName("Account").String()))
	}else{

		userdao.Delete(tx, &userdao.QueryModel{
			UserAccount: 	value.FieldByName("Account").String(),
			Gender:			value.FieldByName("Gender").String(),
		})
	}

	return nil
}

