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

func (t *Index) Index(response *Http.Response) {
	d := t.Resolve(BB{})
	fmt.Println(&d)
	dd := t.Container.Resolve(BB{}, nil, true)
	fmt.Println(&dd)
}

func (t *Index) Name(response *Http.Response) {
	paramVar, _ := t.Request.ParamVar("name")
	response.Echo(paramVar)
}
