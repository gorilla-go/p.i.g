package Bootstrap

import (
	"php-in-go/App/Exception"
	Http3 "php-in-go/App/Http"
	"php-in-go/Config"
	Container2 "php-in-go/Include/Container"
	"php-in-go/Include/Contracts/Container"
	"php-in-go/Include/Contracts/Debug"
	Http2 "php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Contracts/Routing"
	Server2 "php-in-go/Include/Contracts/Server"
	"php-in-go/Include/Http"
	Routing2 "php-in-go/Include/Routing"
	"php-in-go/Routes"
)

type App struct {
	container        Container.IContainer
	server           Server2.IServer
	router           Routing.IRouter
	exceptionHandler Debug.IExceptionHandler
	kernel           Http2.IKernel
}

func (a *App) Initializer(server Server2.IServer) {
	a.server = server
	a.container = server.(Container.IContainerAvailable).GetContainer()

	// set http kernel route
	a.container.AddBinding((*Routing.IRouter)(nil), Container2.NewBindingImpl(&Routing2.Router{}))
	a.router = a.container.GetSingletonByAbstract((*Routing.IRouter)(nil)).(Routing.IRouter)
	a.router.Initializer(Routes.Route(), Config.Route())

	// set http exception handler
	a.container.AddBinding((*Debug.IExceptionHandler)(nil), Container2.NewBindingImpl(&Exception.Handler{}))
	a.exceptionHandler = a.container.GetSingletonByAbstract((*Debug.IExceptionHandler)(nil)).(Debug.IExceptionHandler)

	// set http kernel
	a.container.AddBinding((*Http2.IKernel)(nil), Container2.NewBindingImpl(&Http3.Kernel{}))
	a.kernel = a.container.GetSingletonByAbstract((*Http2.IKernel)(nil)).(Http2.IKernel)

	a.kernel.Bootstrap(a)
}

func (a *App) Handle(request *Http.Request, response *Http.Response) *Http.Response {
	return a.kernel.Handle(request, response)
}

func (a *App) GetContainer() Container.IContainer {
	return a.container
}

func (a *App) GetServer() Server2.IServer {
	return a.server
}

func (a *App) GetRouter() Routing.IRouter {
	return a.router
}

func (a *App) GetExceptionHandler() Debug.IExceptionHandler {
	return a.exceptionHandler
}
