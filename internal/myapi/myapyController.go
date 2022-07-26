package myapi

import (
	"github.com/gin-gonic/gin"
)

type MyApiController interface {
	GetApi(c *gin.Context)
}

type myApiController struct {
}

func NewMyApiController() MyApiController {
	return &myApiController{}
}

func (Api *myApiController) GetApi(c *gin.Context) {
	data, _ := c.GetQuery("data")

	response := &Dto{Data: data}

	c.JSON(200, response)
}
