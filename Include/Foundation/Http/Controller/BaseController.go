package Controller

import (
	"fmt"
	"php-in-go/Include/Http"
)

type BaseController struct {
	Request  *Http.Request
	Response *Http.Response
}

func (c *BaseController) NoFound() *Http.Response {
	fmt.Println("404.")
	return c.Response
}
