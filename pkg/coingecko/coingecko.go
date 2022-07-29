package coingecko

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CoinGeckoItem struct {
	ID         string     `json:"id"`
	MarketData marketData `json:"market_data"`
}

type marketData struct {
	CurrentPrice struct {
		Usd float64 `json:"usd"`
	} `json:"current_price"`
}

func GetCoinData(id string) (*CoinGeckoItem, error) {
	response := &CoinGeckoItem{}
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/" + id)

	if err != nil {
		return response, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, response)

	return response, err

}
