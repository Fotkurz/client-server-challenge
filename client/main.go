package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Exchange struct {
	USDBRL struct {
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
	} `json:"USDBRL"`
}

func (e *Exchange) String() string {
	return fmt.Sprintf(
		"Usd-Brl(Name=%s,High=%v,Low=%v,VarBidy=%v,Timestamp=%v)",
		e.USDBRL.Name,
		e.USDBRL.High,
		e.USDBRL.Low,
		e.USDBRL.VarBid,
		e.USDBRL.Timestamp,
	)
}

func main() {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data Exchange
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		panic(err)
	}

	fmt.Printf("Status: %v\n", res.Status)
	fmt.Printf("Body: %+v", data.String())
}
