package router

import (
	"net/http"

	"github.com/UltimateImage/VanGogh/handler/convert"
	"github.com/UltimateImage/VanGogh/pkg/errs"

	"github.com/UltimateImage/VanGogh/handler/sc"
	"github.com/UltimateImage/VanGogh/router/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(middlewares.Secure)

	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": errs.NotFound})
	})

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusAccepted, "Welcome! (test)")
	})

	check := g.Group("/sc")
	check.GET("/health", sc.HealthCheck)

	v1 := g.Group("/v1")
	v1.GET("/ascii", convert.AsciiHandler)

	return g
}
