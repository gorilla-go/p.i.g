package Http

import (
	"php-in-go/Include/Contracts/Http/App"
	Http "php-in-go/Include/Http"
)

type IKernel interface {
	Bootstrap(container App.IApp)
	Handle(request *Http.Request, response *Http.Response)
}
