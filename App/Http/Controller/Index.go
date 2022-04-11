package Controller

import (
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	Controller.BaseController
}

func (t *Index) Index(request *Http.Request, response *Http.Response) {
	var m map[string]string
	m["A"] = "a"
	response.Echo("Dd")
}

func (t *Index) Name(response *Http.Response) {
	response.Echo(t.Session.GetSession(""))
}
