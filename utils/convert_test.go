package utils

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestBytesToBase64(t *testing.T) {
	data := []byte("123456")
	res, err := Base64ToBytes(BytesToBase64(data))
	if err != nil {
		t.Errorf("convert error! %v", err.Error())
	}
	assert.Equal(t, res, data)
}
