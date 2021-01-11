package router

import (
	"VanGogh/handler/sc"
	"VanGogh/router/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Load default router setting
func Load(g *gin.Engine, handlers ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middlewares.Secure)
	g.Use(handlers...)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Resource not found")
	})
	check := g.Group("/sc")
	{
		check.GET("/health", sc.HealthCheck)
	}
	return g
}
