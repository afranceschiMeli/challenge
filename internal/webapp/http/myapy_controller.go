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
	ids := []string{"bitcoin", "ethereum", "dash"}
	dto := []webapp.MyApiItem{}

	response := coingecko.GetCoinsData(ids)
	withError := false

	for _, r := range response {
		data := webapp.MyApiItem{ID: r.ID, Partial: false, Content: r.MarketData}
		if r.Err != nil {
			data.Content = nil
			withError = true
			data.Partial = true
		}
		dto = append(dto, data)
	}

	if withError {
		c.JSON(206, dto)
	} else {

		c.JSON(200, dto)
	}

}
