package webapp

import "github.com/gin-gonic/gin"

type MyApiController interface {
	GetApi(c *gin.Context)
}
type MyApiItem struct {
	ID      string      `json:"id"`
	Content interface{} `json:"content,omitempty"`
	Partial bool        `json:"partial"`
}
