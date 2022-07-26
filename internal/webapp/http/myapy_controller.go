package http

import (
	"challenge/v1/internal/webapp"

	"github.com/gin-gonic/gin"
)

type myApiController struct {
}

func NewMyApiController() webapp.MyApiController {
	return &myApiController{}
}

func (api *myApiController) GetApi(c *gin.Context) {
	data, _ := c.GetQuery("data")

	response := &webapp.MyApiDto{Data: data}

	c.JSON(200, response)
}
