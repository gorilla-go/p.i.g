package Controller

import (
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	Controller.BaseController
}

func (t *Index) Index(request *Http.Request, response *Http.Response) {
	t.Response.Download(t.GetRoot()+"main.go", "main.go")
}

func (t *Index) Name(response *Http.Response) {
	paramVar, _ := t.Request.ParamVar("name")
	response.Echo(paramVar)
}
