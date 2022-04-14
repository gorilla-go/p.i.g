package Controller

import (
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	Controller.BaseController
}

func (t *Index) Index(response *Http.Response) {
	t.Resolve(func(session Session.ISession) {
		session.SetSession("a", "b", 3600)
		response.Echo(session.GetSession("a"))
	})
}

func (t *Index) Name(response *Http.Response) {
	response.Echo(t.GetSession("k"))
}
