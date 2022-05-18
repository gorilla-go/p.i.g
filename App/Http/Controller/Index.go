package Controller

import (
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	Controller.BaseController
}

func (t *Index) Index() {
	t.Response.Echo(t.Request.RequestURI)
}

func (t *Index) Name(response *Http.Response) {

}
