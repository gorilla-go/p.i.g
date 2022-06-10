package Routing

import (
	"errors"
	"net/url"
	"php-in-go/Include/Config"
	"php-in-go/Include/Container"
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Routing/Component"
	"sort"
	"strings"
)

type Router struct {
	routeMaps   []*Component.RouteMap
	routeConfig Config.Loader
}

func (r *Router) Initializer(routeMaps []*Component.RouteMap, routeConfig Config.Loader) {
	// set route maps.
	r.routeMaps = routeMaps
	r.routeConfig = routeConfig
}

func (r *Router) Resolve(request *Request.Request) *Target {
	requestUri := request.RequestURI

	// remove query
	routeFormatArr := strings.Split(requestUri, "?")
	if len(routeFormatArr) == 2 {
		requestUri = routeFormatArr[0]
	}

	// remove target.
	if v, ok := r.routeConfig["urlHtmlSuffix"]; v != "" && ok && strings.HasSuffix(requestUri, "."+v.(string)) {
		requestUri = requestUri[:len(requestUri)-len("."+v.(string))]
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
	var routeMapsArr []*Component.RouteMap

	for _, routeMap := range r.routeMaps {
		if Container.GetPackageClassName(routeMap.GetController()) == Container.GetPackageClassName(Controller) &&
			routeMap.GetMethod() == method {
			routeMapsArr = append(routeMapsArr, routeMap)
		}
	}

	sort.Slice(routeMapsArr, func(i, j int) bool {
		cur := routeMapsArr[i].GetUriFormat()
		splitArr := strings.Split(cur, "/")
		curCount := 0
		for _, s := range splitArr {
			if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
				curCount++
			}
		}
		next := routeMapsArr[j].GetUriFormat()
		splitNextArr := strings.Split(next, "/")
		nextCount := 0
		for _, sn := range splitNextArr {
			if strings.HasPrefix(sn, "{") && strings.HasSuffix(sn, "}") {
				nextCount++
			}
		}
		return curCount > nextCount
	})

	paramsMap := make(map[string]string)
	paramsQuery := params

nextLoop:
	for _, routeMap := range routeMapsArr {
		uriFormat := routeMap.GetUriFormat()
		splitArr := strings.Split(uriFormat, "/")
		for _, s := range splitArr {
			if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
				if _, exist := params[s[1:len(s)-1]]; !exist {
					paramsMap = make(map[string]string)
					paramsQuery = params
					continue nextLoop
				} else {
					paramsMap[s] = params[s[1:len(s)-1]]
					delete(paramsQuery, s[1:len(s)-1])
				}
			}
		}
		for old, newValue := range paramsMap {
			uriFormat = strings.Replace(uriFormat, old, newValue, 1)
		}

		// add suffix.
		if v, ok := r.routeConfig["urlHtmlSuffix"]; v != "" && ok && routeMap.GetUriFormat() != "/" {
			uriFormat += "." + r.routeConfig["urlHtmlSuffix"].(string)
		}

		uri, err := url.Parse(uriFormat)
		if err != nil {
			panic(err)
		}
		rawQuery := uri.Query()
		for k, v := range paramsQuery {
			rawQuery.Set(k, v)
		}
		if rawQuery.Encode() != "" {
			return uri.String() + "?" + rawQuery.Encode()
		}
		return uri.String()
	}
	panic(errors.New("no found route"))
}
