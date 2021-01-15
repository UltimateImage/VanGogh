package errors

var (
	//InvalidImage status
	InvalidImage = &Error{Code: INVALIDIMAGE, Message: "Image uploaded is invalid"}
)

//Error represents customized errors
type Error struct {
	Code    int
	Message string
	Err     error
}

//NewError create a customized errors
func NewError(code int, message string, err error) *Error {
	return &Error{Code: code, Message: message, Err: err}
}
