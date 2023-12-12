package main

import (
	"net/http"

	"github.com/Fotkurz/client-server-challenge/server/internal/exchange"
)

func main() {
	http.HandleFunc("/cotacao", exchange.GetBID)
	http.ListenAndServe(":8080", nil)
}
