package exchange

import (
	"encoding/json"
	"net/http"
)

func GetExchange() (map[string]interface{}, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(res.Body).Decode(&data)

	return data, nil
}
