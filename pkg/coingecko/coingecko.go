package coingecko

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CoinGeckoItem struct {
	ID         string `json:"id"`
	MarketData struct {
		CurrentPrice struct {
			Usd float64 `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
	Partial bool `json:"partial"`
}

func GetCoinData(id string) *CoinGeckoItem {
	response := &CoinGeckoItem{Partial: true}
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/" + id)

	if err != nil {
		return response
	}

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, response)

	response.Partial = false
	return response

}
