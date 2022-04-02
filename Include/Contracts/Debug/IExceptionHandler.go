package Debug

import "php-in-go/Include/Contracts/Exception"

type IExceptionHandler interface {
	Handle(exception Exception.IException)
}
