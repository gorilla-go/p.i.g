package Middleware

import (
	"php-in-go/Include/Http"
	Routing2 "php-in-go/Include/Routing/Component"
)

type IMiddleware interface {
	HandleRequest(request *Http.Request, r *Routing2.RouteMap)
}
