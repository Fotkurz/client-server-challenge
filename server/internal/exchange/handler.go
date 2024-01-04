package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Fotkurz/client-server-challenge/server/internal/db"
)

// TODO: Timeout of 10 ms is always being exceeded, should look into this
const databaseTimeout = 20

func GetBID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Incoming [%s] for %s\n", r.Method, r.URL.Path)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Millisecond*databaseTimeout,
	)
	defer cancel()

	rawExchange, err := GetExchange()
	if err != nil {
		panic(err)
	}

	exchange := NewFromRaw(rawExchange)

	dbConn := db.CONN.WithContext(ctx)
	if err := dbConn.Save(&exchange).Error; err != nil {
		w.WriteHeader(http.StatusRequestTimeout)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		fmt.Printf("failed to save to database, err: %s\n", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	ctx.Done()
	json.NewEncoder(w).Encode(&exchange)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Incoming [%s] for %s\n", r.Method, r.URL.Path)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Millisecond*databaseTimeout,
	)
	defer cancel()

	dbConn := db.CONN.WithContext(ctx)

	var all []Exchange
	if err := dbConn.Find(&all).Error; err != nil {
		log.Fatalf("failed to find all exchanges in database, err: %s", err.Error())
	}

	ctx.Done()
	json.NewEncoder(w).Encode(&all)
}
