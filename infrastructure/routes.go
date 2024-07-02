package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	Gin *gin.Engine
}

func NewGinRouter() GinRouter {
	httpRouter := gin.Default()
	httpRouter.Use(static.Serve("/", static.LocalFile("./views", true)))

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "UP and RUNNING..."})
	})
	return GinRouter{
		Gin: httpRouter,
	}
}
