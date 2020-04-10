package main

import (
	"fmt"
	gormdao "github.com/auroratechit/rf-stock-gateway/src/persistance"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/gorm"
	"github.com/tidwall/gjson"
	"regexp"
)

func stStockSeeder() {
	fmt.Println("Seeding ST 開始")
	tx := gormdao.DB().Begin()

	updateDB(tx, getSSE())
	updateDB(tx, getSZSE())

	tx.Commit()
	fmt.Println("Seeding ST 結束")
}

func updateDB(tx *gorm.DB, list []string) {
	tx.Exec("update stock set type = 'normal' where symbol != ?", "XXXX")
	tx.Exec("update stock set type = 'st' where symbol IN (?)", list)
}

func getSSE() (r []string) {
	client := resty.New()
	resp, _ := client.R().Get("http://www.sse.com.cn/disclosure/listedinfo/riskplate/")
	response := string(resp.Body())

	validStr := regexp.MustCompile(`COMPANY_CODE=[0-9]{6}">[0-9]`)
	results := validStr.FindAllStringSubmatch(response, -1)
	for _, result := range results {
		r = append(r, result[0][13:19])
	}
	return
}

func getSZSE() (r []string) {
	totalPage := int(gjson.Get(getPageResp(1), "0.metadata.pagecount").Int())
	for i := 1; i <= totalPage; i++ {
		result := gjson.Get(getPageResp(i), "0.data")
		result.ForEach(func(key, value gjson.Result) bool {
			r = append(r, gjson.Get(value.String(), "agdm").String())
			return true // keep iterating
		})
	}
	return
}

func getPageResp(pageNo int) string {
	client := resty.New()
	resp, _ := client.R().Get(fmt.Sprintf(`http://www.szse.cn/api/report/ShowReport/data?SHOWTYPE=JSON&CATALOGID=1110&TABKEY=tab1&PAGENO=%v&txtDMorJC=ST`, pageNo))
	return string(string(resp.Body()))
}
