package convert

import (
	"net/http"

	"github.com/UltimateImage/VanGogh/pkg/errs"
	"github.com/UltimateImage/VanGogh/utils"
	"github.com/gin-gonic/gin"
	"github.com/qeesung/image2ascii/convert"
)

func AsciiHandler(c *gin.Context) {
	img, err := utils.ReadImageFromCtx(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.NewError(errs.INVALIDIMAGE, "Invalid Image", err)})
		return
	}
	convertOptions := convert.DefaultOptions
	converter := convert.NewImageConverter()
	res := converter.Image2ASCIIString(*img, &convertOptions)
	c.JSON(http.StatusOK, *errs.OK(res))
}
