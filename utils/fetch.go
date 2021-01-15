package utils

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadImageFromCtx(c *gin.Context) (*image.Image, error) {
	url := c.Query("url")
	if len(url) > 0 {
		img, err := GetImageByUrl(url)
		if err != nil {
			return nil, err
		}
		return img, nil
	}

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, file); err != nil {
		return nil, err
	}
	img, err := BytesToImage(buff.Bytes())
	if err != nil {
		return nil, err
	}
	return img, err
}

//GetImageByUrl gets a image from given url
func GetImageByUrl(url string) (*image.Image, error) {
	pixels, err := GetBytesByURL(url)
	if err != nil {
		return nil, err
	}
	img, err := BytesToImage(pixels)
	if err != nil {
		return nil, err
	}
	return img, nil
}

//GetBytesByURL gets bytes from a url
func GetBytesByURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	pixels, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return pixels, err
}
