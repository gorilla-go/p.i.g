package Middleware

import (
	"php-in-go/Include/Http"
	"php-in-go/Include/Routing/Component"
)

type Middleware struct {
}

func (m *Middleware) HandleRequest(request *Http.Request, r *Component.RouteMap) {

}
