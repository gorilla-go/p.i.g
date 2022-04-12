package Controller

import (
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	Controller.BaseController
}

func (t *Index) Index() {
	t.BaseController.Container.Resolve(func(response *Http.Response) {
		response.Echo("1")
	}, nil, true)
}

func (t *Index) Name(response *Http.Response) {
	paramVar, _ := t.Request.ParamVar("name")
	response.Echo(paramVar)
}
