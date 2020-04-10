package main

import (
	"fmt"
	gormdao "github.com/auroratechit/rf-stock-gateway/src/persistance"
	hotstockdao "github.com/auroratechit/rf-stock-gateway/src/persistance/gorm/hotStockDao"
)

func hotStockSeeder() {
	fmt.Println("Seeding 熱門股列表 開始")
	tx := gormdao.DB().Begin()

	hotList := []string{"601318", "600519", "600030", "000858", "300059", "000063", "000725", "000651", "600036", "600352"}
	for _, stockNo := range hotList {
		if hotstockdao.Count(tx, &hotstockdao.QueryModel{
			Symbol: stockNo,
		}) == 0 {
			hotstockdao.New(tx, &hotstockdao.Model{
				Symbol: stockNo,
			})
		}
	}
	tx.Commit()

	fmt.Println("Seeding 熱門股列表 結束")
}
