package App

import (
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Http/Response"
)

type IApp interface {
	Initializer()
	Handle(request *Request.Request, response *Response.Response)
	Close()
}
