package utils

import (
	"image"
	"io/ioutil"
	"net/http"
)

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
