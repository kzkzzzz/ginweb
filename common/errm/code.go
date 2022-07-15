package errm

type ErrorCode int

const (
	CommonErrorCode        ErrorCode = 1000
	InvalidParamsErrorCode ErrorCode = 1001
	DBErrorCode            ErrorCode = 1002
	ServiceErrorCode       ErrorCode = 1003
	NotFoundErrorCode      ErrorCode = 1004
	SystemErrorCode        ErrorCode = 9999
)

var (
	codeText = map[ErrorCode]string{
		CommonErrorCode:        "Error",
		InvalidParamsErrorCode: "Invalid Params",
		NotFoundErrorCode:      "Not Found",
		DBErrorCode:            "DB Error",
		ServiceErrorCode:       "Service Error",
		SystemErrorCode:        "System Error",
	}
)

func GetCodeText(code ErrorCode) string {
	if v, ok := codeText[code]; ok {
		return v
	} else {
		return "error"
	}
}
