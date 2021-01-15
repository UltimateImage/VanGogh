package errs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	//InvalidImage status
	InvalidImage = &Error{Code: INVALIDIMAGE, Message: "Image uploaded is invalid"}
	//NotFound
	NotFound = &Error{Code: http.StatusNotFound, Message: "Resource Not Found"}
)

//Error represents customized errs
type Error struct {
	Code    int
	Message string
	Err     error
}

//NewError create a customized errs
func NewError(code int, message string, err error) *Error {
	return &Error{Code: code, Message: message, Err: err}
}

func OK(data interface{}) *gin.H {
	return &gin.H{"code": http.StatusOK, "data": data}
}
