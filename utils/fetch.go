package utils

import (
	"io/ioutil"
	"net/http"
)

//GetImageByURL gets a image from a url
func GetImageByURL(url string) ([]byte, error) {
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
