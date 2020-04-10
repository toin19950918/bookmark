package jsonformat

import (
	"time"
)

type QuoteJSON struct {
	StockSymbol  string    `json:"stockSymbol"`
	SettledPrice float64   `json:"settledPrice"`
	PriorSettle  float64   `json:"priorSettle"`
	Bid          float64   `json:"buyPrice"`
	Ask          float64   `json:"sellPrice"`
	Time         time.Time `json:"time"`
}

type CloseOrder struct {
	OrderNo       string  `json:"orderNo"`
	SettledPrice  float64 `json:"settledPrice"`
	SettledAmount int     `json:"settledAmount"`
	Status        string  `json:"status"`
}

type CancelOrder struct {
	OrderNo string `json:"orderNo"`
	Status  string `json:"status"`
}

type Quote struct {
	Time         time.Time `json:"time"`
	StockSymbol  string    `json:"stockSymbol"`
	PriorSettle  float64   `json:"priorSettle"`  //昨收
	OpenPrice    float64   `json:"openPrice"`    //今開
	HighPrice    float64   `json:"highPrice"`    //最高
	LowPrice     float64   `json:"lowPrice"`     //最低
	SettledPrice float64   `json:"settledPrice"` //成交
	Quantity     int       `json:"quantity"`     //成交量
	BuyPrice     float64   `json:"buyPrice"`     //買價
	BuyQuantity  int       `json:"buyQuantity"`
	SellPrice    float64   `json:"sellPrice"` //賣價
	SellQuantity int       `json:"sellQuantity"`
	SimulateTick bool      `json:"SimulateTick"` //模擬報價
}

type Order struct {
	OrderNo     string    `json:"orderNo"`
	Account     string    `json:"account"`
	StockSymbol string    `json:"stockSymbol"`
	OrderType   string    `json:"orderType"`
	Price       float64   `json:"price"`
	Amount      int       `json:"amount"`
	Time        time.Time `json:"time"`
	Status      string    `json:"status"`
	BuySell     string    `json:"buySell"`
}
