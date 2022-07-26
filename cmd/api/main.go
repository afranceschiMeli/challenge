package main

import (
	"challenge/v1/internal/myapi"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	a := myapi.NewMyApiController()

	g.GET("/myapi", a.GetApi)

	g.Run()
}
