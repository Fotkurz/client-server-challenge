package main

import (
	"net/http"

	"github.com/Fotkurz/client-server-challenge/server/internal/db"
	"github.com/Fotkurz/client-server-challenge/server/internal/exchange"
)

func main() {
	db.StartDB()
	db.CONN.AutoMigrate(&exchange.Exchange{})

	http.HandleFunc("/cotacao", exchange.GetBID)
	http.HandleFunc("/all", exchange.GetAll)
	http.ListenAndServe(":8080", nil)
}
