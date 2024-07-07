package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	Gin *gin.Engine
}

func NewGinRouter() GinRouter {
	httpRouter := gin.Default()
	//load template
	httpRouter.Static("/assets", "./public/assets")
	httpRouter.LoadHTMLGlob("views/*.html")

	// Define route for the index page
	httpRouter.GET("/", func(c *gin.Context) {
		// Render the index.html template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home"})
	})
	return GinRouter{
		Gin: httpRouter,
	}
}
