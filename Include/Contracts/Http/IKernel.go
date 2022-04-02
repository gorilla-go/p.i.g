package Http

import (
	Http "php-in-go/Include/Http"
)

type IKernel interface {
	Bootstrap(container IApp)
	Handle(request *Http.Request) *Http.Response
	RequestInstance()
}
