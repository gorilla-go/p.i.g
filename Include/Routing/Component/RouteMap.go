package Component

import (
	"php-in-go/Include/Contracts/Http/Controller"
	"strings"
)

type RouteMap struct {
	uriFormat     string
	controllerPkg interface{}
	methodName    string
	requestMethod []RequestMethod
}

// NewRouteMap create new router map.
func NewRouteMap(path string, pkg Controller.IController, methodName string) *RouteMap {
	return &RouteMap{
		uriFormat:     path,
		controllerPkg: pkg,
		methodName:    methodName,
	}
}

func (m *RouteMap) GetUriFormat() string {
	return m.uriFormat
}

func (m *RouteMap) GetRequestMethods() []RequestMethod {
	return m.requestMethod
}

func (m *RouteMap) GetController() interface{} {
	return m.controllerPkg
}

// GetMethod get controller action method.
func (m *RouteMap) GetMethod() string {
	return m.methodName
}

func (m *RouteMap) SetRequestMethods(method string) *RouteMap {
	methodArr := strings.Split(method, "|")
	for _, methodStr := range methodArr {
		methodFormat := strings.ToUpper(strings.TrimSpace(methodStr))
		switch methodFormat {
		case "GET":
			m.requestMethod = append(m.requestMethod, GET)
			continue
		case "POST":
			m.requestMethod = append(m.requestMethod, POST)
			continue
		case "PUT":
			m.requestMethod = append(m.requestMethod, PUT)
			continue
		case "DELETE":
			m.requestMethod = append(m.requestMethod, DELETE)
			continue
		}
	}
	return m
}
