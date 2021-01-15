package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func BytesToImage(pixels []byte) (*image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(pixels))
	if err != nil {
		return nil, err
	}
	return &img, nil
}

func BytesToBase64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64ToBytes(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
