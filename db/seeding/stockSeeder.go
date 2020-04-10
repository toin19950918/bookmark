package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	gormdao "github.com/auroratechit/rf-stock-gateway/src/persistance"
	stockdao "github.com/auroratechit/rf-stock-gateway/src/persistance/gorm/stockDao"
	"path/filepath"
	"runtime"
	"strings"
)

func stockSeeder() {
	fmt.Println("Seeding 股票列表 開始")
	tx := gormdao.DB().Begin()

	for _, file := range []string{"sse", "szse"} {
		excelPath := rootPath() + fmt.Sprintf("/%v.xlsx", file)
		f, err := excelize.OpenFile(excelPath)
		if err != nil {
			println(err.Error())
			return
		}

		rows := f.GetRows("Sheet1")
		for _, row := range rows {
			stock := &stockdao.Model{
				Symbol:         row[0],
				ExchangeSymbol: strings.ToUpper(file),
				Name:           row[1],
				TimeToMarket:   row[3],
				CategoryCode:   row[2],
				Shares:         100,
				Status:         "enabled",
				Type:           "normal",
			}
			if stockdao.Count(tx, &stockdao.QueryModel{
				Symbol: stock.Symbol,
			}) == 0 {
				stockdao.New(tx, stock)
			}
		}
	}
	tx.Commit()

	fmt.Println("Seeding 股票列表 結束")
}

func rootPath() string {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	return dir + "/../.."
}
