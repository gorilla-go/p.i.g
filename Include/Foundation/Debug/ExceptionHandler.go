package Debug

import (
	"php-in-go/Config"
	"php-in-go/Include/Contracts/Exception"
	"php-in-go/Include/Http"
)

type ExceptionHandler struct {
}

func (h *ExceptionHandler) Handle(exception Exception.IException, response *Http.Response) {
	config := Config.App()
	if config["debug"].(bool) {
		response.EchoWithCode(exception.GetMessage(), 500)
		return
	}
	response.SetCode(500)
}
