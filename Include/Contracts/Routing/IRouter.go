package Routing

import (
	"php-in-go/Include/Http"
	"php-in-go/Include/Routing"
	"php-in-go/Include/Routing/Component"
)

type IRouter interface {
	Initializer(routeMaps []*Component.RouteMap, routeConfig map[string]interface{})
	Resolve(request *Http.Request) *Routing.Target
	Url(Controller interface{}, method string, params map[string]string) string
}
