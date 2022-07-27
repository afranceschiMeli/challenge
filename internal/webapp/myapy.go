package webapp

import "github.com/gin-gonic/gin"

type MyApiController interface {
	GetApi(c *gin.Context)
}
type MyApiDto struct {
	ID      string `json:"id"`
	Content struct {
		Price    float64 `json:"price,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"content,omitempty"`
	Partial bool `json:"partial"`
}
