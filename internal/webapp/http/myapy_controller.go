package http

import (
	"challenge/v1/internal/webapp"
	"challenge/v1/pkg/coingecko"

	"github.com/gin-gonic/gin"
)

type myApiController struct {
}

func NewMyApiController() webapp.MyApiController {
	return &myApiController{}
}

func (api *myApiController) GetApi(c *gin.Context) {
	id := "bitcoin"
	dto := &webapp.MyApiDto{}
	response := coingecko.GetCoinData(id)

	if response.Partial {
		dto.Partial = true
		dto.ID = id
		c.JSON(206, dto)
	} else {
		buildResponse(response, dto)
		c.JSON(200, dto)
	}

}

func buildResponse(resp *coingecko.CoinGeckoItem, dto *webapp.MyApiDto) {
	dto.ID = resp.ID
	dto.Partial = resp.Partial
	dto.Content.Currency = "USD"
	dto.Content.Price = resp.MarketData.CurrentPrice.Usd
}
