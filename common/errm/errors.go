package errm

import "fmt"

var _ error = &Error{}

type Error struct {
	Code    ErrorCode
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", GetCodeText(e.Code), e.Message)
}

func New(code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func Newf(code ErrorCode, format string, a ...any) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
	}
}

func CommonError(message string) *Error {
	return New(CommonErrorCode, message)
}

func ParamsError(message string) *Error {
	return New(InvalidParamsErrorCode, message)
}

func DbError(message string) *Error {
	return New(DBErrorCode, message)
}
