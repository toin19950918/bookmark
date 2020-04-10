package main

import (
	"fmt"
	gormdao "github.com/auroratechit/rf-stock-gateway/src/persistance"
	exchangedao "github.com/auroratechit/rf-stock-gateway/src/persistance/gorm/exchangeDao"
)

func exchangeSeeder() {
	fmt.Println("Seeding 交易所 開始")

	tx := gormdao.DB().Begin()

	exchanges := []*exchangedao.Model{
		{
			ID:         "E0001",
			Category:   "stock",
			Name:       "上海证券交易所",
			Timezone:   8,
			Symbol:     "SSE",
			Status:     "enabled",
			NormalTime: `[ { "startDay":"1", "endDay":"5", "startTime":"0930", "endTime":"1130" },{ "startDay":"1", "endDay":"5", "startTime":"1300", "endTime":"1500" } ]`,
		}, {
			ID:         "E0002",
			Category:   "stock",
			Name:       "深圳证券交易所",
			Timezone:   8,
			Symbol:     "SZSE",
			Status:     "enabled",
			NormalTime: `[ { "startDay":"1", "endDay":"5", "startTime":"0930", "endTime":"1130" },{ "startDay":"1", "endDay":"5", "startTime":"1300", "endTime":"1500" } ]`,
		},
	}

	for _, exchange := range exchanges {
		if exchangedao.Count(tx, &exchangedao.QueryModel{
			Symbol: exchange.Symbol,
		}) == 0 {
			exchangedao.New(tx, exchange)
		}
	}

	tx.Commit()
	fmt.Println("Seeding 交易所 完成")
}
