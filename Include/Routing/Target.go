package Routing

type Target struct {
	Controller  interface{}
	Method      string
	RouteParams map[string]string
}
