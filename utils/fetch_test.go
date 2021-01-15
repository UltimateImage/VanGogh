package utils

import "testing"

var IMG = "https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=3393595332,3048178123&fm=26&gp=0.jpg"

func TestGetBytesByURL(t *testing.T) {
	_, err := GetBytesByURL(IMG)
	if err != nil {
		t.Errorf("fetch error %v", err.Error())
	}
}

func TestGetImageByUrl(t *testing.T) {
	_, err := GetImageByUrl(IMG)
	if err != nil {
		t.Errorf("fetch error %v", err.Error())
	}
}
