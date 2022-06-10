package Middleware

import (
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Http/Response"
)

type IMiddleware interface {
	Handle(
		*Request.Request,
		*Response.Response,
		func(request2 *Request.Request, response2 *Response.Response),
	)
}
