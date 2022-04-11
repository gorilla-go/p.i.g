package Controller

import (
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
}

//NoFound no found action.
func (c *BaseController) NoFound() {
	c.Response.Html("<p>page no found </p>")
}

// GetRouter get router.
func (c *BaseController) GetRouter() Routing.IRouter {
	return c.App.GetRouter()
}
