package exchange

import "gorm.io/gorm"

type Exchange struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
	gorm.Model
}

func NewFromRaw(rawExchange *GetExchangeResponse) *Exchange {
	return &Exchange{
		Code:       rawExchange.USDBRL.Code,
		Codein:     rawExchange.USDBRL.Codein,
		Name:       rawExchange.USDBRL.Name,
		High:       rawExchange.USDBRL.High,
		Low:        rawExchange.USDBRL.Low,
		VarBid:     rawExchange.USDBRL.VarBid,
		PctChange:  rawExchange.USDBRL.PctChange,
		Bid:        rawExchange.USDBRL.Bid,
		Ask:        rawExchange.USDBRL.Ask,
		Timestamp:  rawExchange.USDBRL.Timestamp,
		CreateDate: rawExchange.USDBRL.CreateDate,
	}
}
