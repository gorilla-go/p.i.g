package Middleware

import (
	"php-in-go/Include/Http"
	"php-in-go/Include/Routing"
)

type IMiddleware interface {
	Handle(request *Http.Request, response *Http.Response, target *Routing.Target) bool
}
