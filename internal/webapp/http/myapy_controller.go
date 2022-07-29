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
	dto := &webapp.MyApiDto{ID: id, Partial: true}
	content, err := coingecko.GetCoinData(id)

	if err != nil {
		c.JSON(206, dto)
	} else {
		dto.Content = content
		dto.Partial = false
		c.JSON(200, dto)
	}

}
