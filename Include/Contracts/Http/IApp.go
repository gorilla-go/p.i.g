package Http

import (
	"php-in-go/Include/Contracts/Container"
	"php-in-go/Include/Contracts/Debug"
	"php-in-go/Include/Contracts/Routing"
	Server2 "php-in-go/Include/Contracts/Server"
	"php-in-go/Include/Http"
)

type IApp interface {
	Initializer(server Server2.IServer)
	Handle(request *Http.Request) *Http.Response
	GetContainer() Container.IContainer
	GetServer() Server2.IServer
	GetExceptionHandler() Debug.IExceptionHandler
	GetRouter() Routing.IRouter
}
