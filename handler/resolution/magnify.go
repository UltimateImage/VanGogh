package resolution

import (
	"VanGogh/pkg/errors"

	"github.com/gin-gonic/gin"
)

//HandleMagnify magnifies a given picture
func HandleMagnify(c *gin.Context) {
	_, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(410, gin.H{"errors": errors.InvalidImage})
		return
	}
}
