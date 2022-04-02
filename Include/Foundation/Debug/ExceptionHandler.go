package Debug

import (
	"log"
	"php-in-go/Include/Contracts/Exception"
)

type ExceptionHandler struct {
}

func (h *ExceptionHandler) Handle(exception Exception.IException) {
	log.Println(exception.GetMessage())
}
