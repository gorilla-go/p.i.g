package Log

import (
	"php-in-go/Include/Http/Request"
	Http2 "php-in-go/Include/Http/Response"
)

type ILog interface {
	StartLogManager()
	Log(request *Request.Request, response *Http2.Response)
	CloseLogManager()
}
