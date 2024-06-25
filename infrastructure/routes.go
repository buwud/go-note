package infrastructure

import "github.com/gin-gonic/gin"

type GinRouter struct {
	Gin *gin.Engine
}

func NewGinRouter() GinRouter {
	httpRouter := gin.Default()
	httpRouter.G
}
