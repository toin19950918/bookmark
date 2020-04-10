package error

// CustomError return a new CustomError instance
func CustomError(code int, msg string) *Error {
	return &Error{
		code:    code,
		message: msg,
	}
}
