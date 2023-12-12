package exchange

import (
	"encoding/json"
	"net/http"
)

func GetBID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	data, err := GetExchange()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(&data)
}
