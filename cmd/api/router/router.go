package router

import (
	"challenge/v1/internal/webapp"

	"github.com/gin-gonic/gin"
)

func CreateRouter(g *gin.Engine, myApiControler webapp.MyApiController) {
	g.GET("/myapi", myApiControler.GetApi)
}
