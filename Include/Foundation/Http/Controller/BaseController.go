package Controller

import (
	"fmt"
	"php-in-go/Include/Contracts/Container"
	Http2 "php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Http"
)

type BaseController struct {
	App       Http2.IApp
	Container Container.IContainer
	Request   *Http.Request
	Response  *Http.Response
	Session   Session.ISession
}

func (c *BaseController) NoFound() *Http.Response {
	fmt.Println("404.")
	return c.Response
}
