package Debug

import (
	"php-in-go/Include/Contracts/Exception"
	"php-in-go/Include/Http/Response"
)

type ExceptionHandler struct {
}

func (h *ExceptionHandler) Handle(exception Exception.IException, response *Response.Response) {
	response.EchoWithCode(exception.GetMessage(), 500)
}
