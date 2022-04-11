package Debug

import (
	"php-in-go/Include/Contracts/Exception"
	"php-in-go/Include/Http"
)

type ExceptionHandler struct {
}

func (h *ExceptionHandler) Handle(exception Exception.IException, response *Http.Response) {
	response.Echo(exception.GetMessage())
}
