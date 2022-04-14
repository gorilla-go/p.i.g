package Controller

import (
	"net/http"
	"os"
	"php-in-go/Include/Contracts/Container"
	Http2 "php-in-go/Include/Contracts/Http/App"
	"php-in-go/Include/Http"
)

type BaseController struct {
	App       Http2.IApp
	Container Container.IContainer
	Request   *Http.Request
	Response  *Http.Response
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

func (c *BaseController) SetSession(k string, v string) {
	c.App.GetSession().SetSession(k, v, c.GetConfig("sessionExpire").(int))
}

func (c *BaseController) GetSession(k string) interface{} {
	return c.App.GetSession().GetSession(k)
}

func (c *BaseController) GetConfig(s string) interface{} {
	configs := c.App.GetConfigs()
	return configs[s]
}
