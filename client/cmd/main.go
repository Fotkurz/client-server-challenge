package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Fotkurz/client-server-challenge/client/internal/exchange"
)

const contextTimeout = 300

func main() {
	setupFile()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*contextTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("failed to create a new request with context, err: %s", err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to request the server, err: %s", err.Error())
	}
	defer res.Body.Close()

	ctx.Done()

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusRequestTimeout:
			var data map[string]string
			json.NewDecoder(res.Body).Decode(&data)
			log.Fatalf("the operation timedout: [%v] %v", res.StatusCode, data["error"])
		}
		log.Fatalf("failed to request the server: [%v]", res.StatusCode)
	}

	var data exchange.Exchange
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatalf("failed to decode response body: %s", err.Error())
	}

	data.SaveToFile("./data/exchange.txt")

}

func setupFile() {
	err := os.MkdirAll("data", 0777)
	if err != nil {
		log.Fatalf("failed to create folder: %s", err.Error())
	}

	myFile, err := os.Create("data/exchange.txt")
	if err != nil {
		log.Fatalf("faield to create exchange file: %s", err.Error())
	}

	myFile.Close()
}
