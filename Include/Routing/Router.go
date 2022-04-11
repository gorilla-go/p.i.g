package Routing

import (
	"php-in-go/Include/Http"
	"php-in-go/Include/Routing/Component"
	"strings"
)

type Router struct {
	routeMaps   []*Component.RouteMap
	routeConfig map[string]interface{}
}

func (r *Router) Initializer(routeMaps []*Component.RouteMap, routeConfig map[string]interface{}) {
	// set route maps.
	r.routeMaps = routeMaps
	r.routeConfig = routeConfig
}

func (r *Router) Resolve(request *Http.Request) *Target {
	requestUri := request.RequestURI

	// remove target.
	if v, ok := r.routeConfig["urlHtmlSuffix"]; v != "" && ok && strings.HasSuffix(requestUri, "."+v.(string)) {
		requestUri = requestUri[:len(requestUri)-len("."+v.(string))]
	}

	// remove query
	routeFormatArr := strings.Split(requestUri, "?")
	if len(routeFormatArr) == 2 {
		requestUri = routeFormatArr[0]
	}

	// split.
	routePathArr := strings.Split(requestUri, "/")

	// set query params
	values := request.URL.Query()
	request.Params = &values

	// resolve route path.
routeLoop:
	for _, route := range r.routeMaps {
		uriFormat := route.GetUriFormat()
		uriFormatArr := strings.Split(uriFormat, "/")

		// check struct.
		if len(uriFormatArr) != len(routePathArr) {
			continue
		}

		// check request method.
		if len(route.GetRequestMethods()) != 0 {
			var inArray = false
			for _, allowRequestMethod := range route.GetRequestMethods() {
				if strings.ToUpper(request.Method) == allowRequestMethod.ToString() {
					inArray = true
				}
			}
			if inArray == false {
				return nil
			}
		}

		// pattern item
		for key, pathItem := range uriFormatArr {
			if len(pathItem) > 2 && pathItem[0:1] == "{" && pathItem[len(pathItem)-1:] == "}" {
				k := pathItem[1 : len(pathItem)-1]
				if request.Params.Has(k) {
					request.Params.Set(k, routePathArr[key])
					continue
				}
				request.Params.Add(k, routePathArr[key])
				continue
			}

			if strings.ToLower(pathItem) != strings.ToLower(routePathArr[key]) {
				continue routeLoop
			}
		}

		return &Target{
			Controller: route.GetController(),
			Method:     route.GetMethod(),
		}
	}

	// no found, nil.
	return nil
}

// Url resolve controller action to format url.
func (r *Router) Url(Controller interface{}, method string, params map[string]string) string {
	return ""
}
