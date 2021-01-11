package sc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HealthCheck checks server status
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Good")
}
