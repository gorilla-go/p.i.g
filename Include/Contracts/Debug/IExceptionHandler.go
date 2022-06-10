package Debug

import (
	"php-in-go/Include/Contracts/Exception"
	"php-in-go/Include/Http/Response"
)

type IExceptionHandler interface {
	Handle(exception Exception.IException, response *Response.Response)
}
