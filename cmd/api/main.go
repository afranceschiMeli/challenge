package main

import (
	"challenge/v1/cmd/api/router"
	"challenge/v1/internal/webapp/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	myApiController := http.NewMyApiController()

	router.CreateRouter(g, myApiController)

	g.Run()
}
