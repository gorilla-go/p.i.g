package Controller

import (
	"net/http"
	"os"
	"php-in-go/Include/Config"
	"php-in-go/Include/Contracts/Container"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Http/Response"
)

type BaseController struct {
	Container Container.IContainer
	Request   *Request.Request
	Response  *Response.Response
}

//NoFound no found action.
func (c *BaseController) NoFound() {
	c.Response.HtmlWithCode("<h2>404</h2><h4>page no found! </h4>", http.StatusNotFound)
}

// GetRoot get root file path
func (c *BaseController) GetRoot() string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return root + "/"
}

// Resolve receive object from container.
func (c *BaseController) Resolve(abstract interface{}) interface{} {
	return c.Container.Resolve(abstract, nil, false)
}

// ResolveNew receive new object from container.
func (c *BaseController) ResolveNew(abstract interface{}) interface{} {
	return c.Container.Resolve(abstract, nil, true)
}

func (c *BaseController) SetSession(k string, v string) {
	c.Container.GetSingleton("session").(Session.ISession).SetSession(k, v, c.Request, c.Response)
}

func (c *BaseController) GetSession(k string) interface{} {
	return c.Container.GetSingleton("session").(Session.ISession).GetSession(k, c.Request, c.Response)
}

func (c *BaseController) GetConfig(s string) interface{} {
	return c.Container.GetSingleton("config").(Config.Loader).Load(s)
}
