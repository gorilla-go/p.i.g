package Controller

import (
	"fmt"
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	Controller.BaseController
}

type BB struct {
	A int
}

func (t *Index) Index() {
	t.Resolve(func(request *Http.Request) {
		fmt.Println(request.RequestURI)
	})
}

func (t *Index) Name(response *Http.Response) {
	paramVar, _ := t.Request.ParamVar("name")
	response.Echo(paramVar)
}
