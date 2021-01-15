package resolution

import (
	"github.com/UltimateImage/VanGogh/pkg/errs"
	"github.com/gin-gonic/gin"
)

//MagnifyHandler magnifies a given picture
func MagnifyHandler(c *gin.Context) {
	_, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(410, gin.H{"errs": errs.INVALIDIMAGE})
		return
	}
}
