package Controller

import (
	"net/http"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Container"
	Http2 "php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Contracts/Routing"
	"php-in-go/Include/Http"
)

type BaseController struct {
	App       Http2.IApp
	Container Container.IContainer
	Request   *Http.Request
	Response  *Http.Response
	Session   Session.ISession
	Cache     Cache.ICache
}

//NoFound no found action.
func (c *BaseController) NoFound() {
	c.Response.HtmlWithCode("<h2>404</h2><h4>page no found! </h4>", http.StatusNotFound)
}

// GetRouter get router.
func (c *BaseController) GetRouter() Routing.IRouter {
	return c.App.GetRouter()
}
