package Routing

import (
	"php-in-go/Include/Config"
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Routing"
	"php-in-go/Include/Routing/Component"
)

type IRouter interface {
	Initializer(routeMaps []*Component.RouteMap, routeConfig Config.Loader)
	Resolve(request *Request.Request) *Routing.Target
	Url(Controller interface{}, method string, params map[string]string) string
}
