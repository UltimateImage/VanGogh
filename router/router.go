package router

import (
	"VanGogh/handler/sc"
	"VanGogh/router/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Setup(handlers ...gin.HandlerFunc) *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(middlewares.Secure)
    g.Use(handlers...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Resource not found")
	})

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusAccepted, "Welcome! (test)")
	})

	check := g.Group("/sc")
    check.GET("/health", sc.HealthCheck)

	return g

}
