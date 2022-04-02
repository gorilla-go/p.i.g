package Exceptions

type Exception struct {
	code    int
	message string
}

func NewException(code int, message string) *Exception {
	return &Exception{
		code:    code,
		message: message,
	}
}

func (e *Exception) GetCode() int {
	return e.code
}

func (e *Exception) GetMessage() string {
	return e.message
}
