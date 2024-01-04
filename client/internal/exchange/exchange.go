package exchange

import (
	"fmt"
	"log"
	"os"
)

type Exchange struct {
	Bid string `json:"bid"`

	// USDBRL struct {
	// 	Code       string `json:"code"`
	// 	Codein     string `json:"codein"`
	// 	Name       string `json:"name"`
	// 	High       string `json:"high"`
	// 	Low        string `json:"low"`
	// 	VarBid     string `json:"varBid"`
	// 	PctChange  string `json:"pctChange"`
	// 	Ask        string `json:"ask"`
	// 	Timestamp  string `json:"timestamp"`
	// 	CreateDate string `json:"create_date"`
	// } `json:"USDBRL"`
}

func (e *Exchange) String() string {
	return fmt.Sprintf("Dólar: %s", e.Bid)
}

func (e *Exchange) SaveToFile(path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open file %s, err: %s", path, err.Error())
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("Dólar: %s\n", e.Bid)); err != nil {
		log.Fatalf("failed to write to file, err: %s", err.Error())
	}
}
