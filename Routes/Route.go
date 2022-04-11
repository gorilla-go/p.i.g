package Routes

import (
	"php-in-go/App/Http/Controller"
	"php-in-go/Include/Routing/Component"
)

func Route() (maps []*Component.RouteMap) {
	// route maps.
	maps = []*Component.RouteMap{
		// user routes.
		Component.NewRouteMap("/", &Controller.Index{}, "Index"),
		Component.NewRouteMap("/name", &Controller.Index{}, "Name"),

		Component.NewRouteMap("/{name}/get_name", &Controller.Index{}, "Name"),
	}
	return
}
