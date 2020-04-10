package main

import (
	"fmt"
	gormdao "github.com/robin019/bookmark/persistance"
	userdao "github.com/robin019/bookmark/persistance/gorm/userDao"
)

func main() {
	tx := gormdao.DB()
	user := userdao.Get(tx, &userdao.QueryModel{
		Gender: "female",
	})
	fmt.Println(user)
}
