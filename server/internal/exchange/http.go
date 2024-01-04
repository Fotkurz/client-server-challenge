package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GetExchangeResponse struct {
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

const getExchangeTimeout = 200
const exchangeAPIUrl = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

func GetExchange() (*GetExchangeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*getExchangeTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", exchangeAPIUrl, nil)
	if err != nil {
		return &GetExchangeResponse{}, fmt.Errorf("failed to build request, err: %s", err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &GetExchangeResponse{}, fmt.Errorf("failed to request the server, err: %s", err.Error())
	}
	defer res.Body.Close()

	var data GetExchangeResponse
	json.NewDecoder(res.Body).Decode(&data)

	return &data, nil
}
