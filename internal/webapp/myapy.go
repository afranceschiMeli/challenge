package webapp

import "github.com/gin-gonic/gin"

type MyApiController interface {
	GetApi(c *gin.Context)
}

type MyApiDto struct {
	Data string `json:"data"`
}
