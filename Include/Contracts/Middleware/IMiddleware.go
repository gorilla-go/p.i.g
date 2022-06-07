package Middleware

import (
	"php-in-go/Include/Http"
)

type IMiddleware interface {
	Handle(
		*Http.Request,
		*Http.Response,
		func(request2 *Http.Request, response2 *Http.Response),
	)
}
