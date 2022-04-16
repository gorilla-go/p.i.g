package Controller

import (
	"php-in-go/App/Http/Model"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	Controller.BaseController
}

func (t *Index) Index(response *Http.Response) {
	userModel := t.Resolve(Model.UserModel{}).(Model.UserModel)
	response.Echo(userModel.Table)
}

func (t *Index) Name(response *Http.Response) {
	t.Resolve(func(session Session.ISession, request *Http.Request) {
		response.Echo(session.GetSession(
			"a",
			request,
			response,
		))
	})
}
