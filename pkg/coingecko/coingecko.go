package coingecko

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

type CoinGeckoResponse struct {
	ID         string      `json:"id"`
	MarketData *marketData `json:"market_data"`
	Err        error
}

type marketData struct {
	CurrentPrice struct {
		Usd float64 `json:"usd"`
	} `json:"current_price"`
}

func getCoinData(id string) *CoinGeckoResponse {
	response := &CoinGeckoResponse{}
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/" + id)

	if err != nil {
		response.ID = id
		response.Err = err
		return response
	}

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, response)

	return response

}

func GetCoinsData(ids []string) []*CoinGeckoResponse {
	response := []*CoinGeckoResponse{}

	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)
	g, _ := errgroup.WithContext(ctx)

	ch := make(chan *CoinGeckoResponse)

	for _, id := range ids {
		id := id
		g.Go(func() error {
			data := getCoinData(id)
			ch <- data
			return nil
		})
	}

	go func() {
		g.Wait()
		close(ch)
	}()

	for resp := range ch {
		response = append(response, resp)
	}

	return response
}
