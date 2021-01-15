package error

var (
	//Success status
	Success = &Error{Code: OK, Message: "Success"}
)

//Error represents customized errors
type Error struct {
	Code    int
	Message string
	Err     error
}

//NewError create a customized error
func NewError(code int, message string, err error) *Error {
	return &Error{Code: code, Message: message, Err: err}
}
