package Routes

import (
	"php-in-go/App/Http/Controller"
	"php-in-go/Include/Routing/Component"
)

func Route() (maps []*Component.RouteMap) {
	// route maps.
	maps = []*Component.RouteMap{
		// user routes.
		Component.NewRouteMap("/", &Controller.Index{}, "Indexx"),

		Component.NewRouteMap("/{name}/get_name", &Controller.Index{}, "Index"),

		// base routes.
		Component.NewRouteMap("/500", &Controller.Common{}, "PageError"),
		Component.NewRouteMap("/404", &Controller.Common{}, "PageNoFound"),
	}
	return
}
