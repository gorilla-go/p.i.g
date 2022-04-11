package Controller

import (
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	Controller.BaseController
}

func (t *Index) Index(request *Http.Request, response *Http.Response) {
	println(t.GetRouter().Url(&Index{}, "Index", map[string]string{"cc": "bb", "t": "aa"}))
}

func (t *Index) Name(response *Http.Response) {
	response.Echo(t.Session.GetSession("a"))
}
