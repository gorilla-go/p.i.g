package Debug

import (
	"php-in-go/Include/Contracts/Exception"
	"php-in-go/Include/Http"
)

type IExceptionHandler interface {
	Handle(exception Exception.IException, response *Http.Response)
}
