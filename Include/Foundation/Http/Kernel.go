package Http

import (
	"fmt"
	Http3 "php-in-go/App/Http"
	Container2 "php-in-go/Include/Container"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Container"
	Http2 "php-in-go/Include/Contracts/Http/App"
	"php-in-go/Include/Contracts/Http/Controller"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Foundation/Exceptions"
	Controller2 "php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
	"php-in-go/Include/Routing"
	"reflect"
	"runtime"
)

type Kernel struct {
	app              Http2.IApp
	requestContainer *Container2.Container
}

func (k *Kernel) Bootstrap(app Http2.IApp) {
	k.app = app
}

// Handle each request kernel.
func (k *Kernel) Handle(request *Http.Request, response *Http.Response) {
	// register error handle.
	defer func() {
		// catch exception.
		err := recover()
		if err != nil {
			v := make([]byte, 1024*2)
			runtime.Stack(v, true)
			k.app.GetExceptionHandler().Handle(
				Exceptions.NewException(
					1,
					fmt.Sprintf("%v\n\n%v", err, string(v)),
				),
				response,
			)

			// set runtime exception stack
			response.ErrorStack = string(v)
			response.ErrorMessage = fmt.Sprintf("%v", err)
		}

		// logger.
		k.app.GetLogger().Log(request, response)
	}()

	// resolve container foundation
	container := k.containerFoundation(request, response)

	// route resolve.
	actionTarget := k.app.GetRouter().Resolve(request)

	// middleware
	if k.middlewareHandler(actionTarget, request, response) == false {
		return
	}

	// call method.
	k.dispatch(actionTarget, request, response, container)
}

// dispatch call correct controller method.
func (k *Kernel) dispatch(
	target *Routing.Target,
	request *Http.Request,
	response *Http.Response,
	container Container.IContainer,
) {
	// page no found
	if target == nil {
		baseController := &Controller2.BaseController{
			Request:  request,
			Response: response,
		}
		baseController.NoFound()
		return
	}

	// resolve controller params.
	targetController := container.Resolve(target.Controller, nil, true).(Controller.IController)

	// resolve target method
	targetMethod := reflect.ValueOf(targetController).MethodByName(target.Method)

	// no found method ? to NoFound method in base controller.
	if targetMethod.IsValid() == false {
		panic(fmt.Sprintf("Controller method no found: %s", target.Method))
	}

	// call method.
	container.Resolve(targetMethod.Interface(), nil, true)
}

func (k *Kernel) containerFoundation(request *Http.Request, response *Http.Response) Container.IContainer {
	// init request container
	requestContainer := Container2.NewContainer()

	// app.
	requestContainer.AddBinding(
		(*Http2.IApp)(nil),
		Container2.NewBindingImpl(k.app).SetShared().SetAlias("app"),
	)

	// request container.
	requestContainer.AddBinding(
		(*Container.IContainer)(nil),
		Container2.NewBindingImpl(requestContainer).SetShared().SetAlias("container"),
	)

	// binding request.
	requestContainer.Singleton(request, "request")

	// binding response
	requestContainer.Singleton(response, "response")

	// cache drive.
	requestContainer.AddBinding(
		(*Cache.ICache)(nil),
		Container2.NewBindingImpl(
			k.app.GetCache(),
		).SetAlias("cache"),
	)

	// session drive.
	requestContainer.AddBinding(
		(*Session.ISession)(nil),
		Container2.NewBindingImpl(
			k.app.GetSession(),
		).SetShared().SetAlias("session"),
	)

	return requestContainer
}

func (k *Kernel) middlewareHandler(target *Routing.Target, request *Http.Request, response *Http.Response) bool {
	middlewares := Http3.Middlewares()
	if len(middlewares) > 0 {
		for _, middleware := range middlewares {
			if middleware.Handle(request, response, target) == false {
				return false
			}
		}
	}
	return true
}
