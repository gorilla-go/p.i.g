package Controller

import (
	"fmt"
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	Controller.BaseController
}

func (t *Index) Index(request *Http.Request, response *Http.Response) {
	fmt.Println(request.PostVar("name"))
	fmt.Println(request.ParamVar("a"))
	fmt.Println(request.IsAjax())
}

func (t *Index) Name(response *Http.Response) {
	response.Echo(t.Session.GetSession("a"))
}
