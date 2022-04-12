package Http

import (
	"fmt"
	Http3 "php-in-go/App/Http"
	Container2 "php-in-go/Include/Container"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Container"
	"php-in-go/Include/Contracts/Debug"
	Http2 "php-in-go/Include/Contracts/Http"
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

// GetRequestContainer get current request container.
func (k *Kernel) GetRequestContainer() *Container2.Container {
	return k.requestContainer
}

// GetApp get global application.
func (k *Kernel) GetApp() Http2.IApp {
	return k.app
}

// Handle each request kernel.
func (k *Kernel) Handle(request *Http.Request, response *Http.Response) {
	// register error handle.
	defer func() {
		// catch exception.
		err := recover()
		if err != nil {
			exception := k.app.GetContainer().GetSingletonByAbstract((*Debug.IExceptionHandler)(nil)).(Debug.IExceptionHandler)
			v := make([]byte, 1024*2)
			runtime.Stack(v, true)
			exception.Handle(
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

	// init request container
	k.requestContainer = Container2.NewContainer()

	// resolve container foundation
	k.containerFoundation(request, response)

	// route resolve.
	actionTarget := k.app.GetRouter().Resolve(request)

	// middleware
	if k.middlewareHandler(actionTarget, request, response) == false {
		return
	}

	// call method.
	k.dispatch(actionTarget, request, response)
}

// dispatch call correct controller method.
func (k *Kernel) dispatch(target *Routing.Target, request *Http.Request, response *Http.Response) {
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
	targetController := k.GetRequestContainer().Resolve(target.Controller, nil, true).(Controller.IController)

	// resolve target method
	targetMethod := reflect.ValueOf(targetController).MethodByName(target.Method)

	// no found method ? to NoFound method in base controller.
	if targetMethod.IsValid() == false {
		panic(fmt.Sprintf("Controller method no found: %s", target.Method))
	}

	// call method.
	k.GetRequestContainer().Resolve(targetMethod.Interface(), nil, false)
}

func (k *Kernel) containerFoundation(request *Http.Request, response *Http.Response) {
	requestContainer := k.GetRequestContainer()

	// app.
	requestContainer.AddBinding((*Http2.IApp)(nil), Container2.NewBindingImpl(k.app))

	// request container.
	requestContainer.AddBinding((*Container.IContainer)(nil), Container2.NewBindingImpl(k.requestContainer))

	// binding request.
	requestContainer.Singleton(request, "request")

	// binding response
	requestContainer.Singleton(response, "response")

	// cache drive.
	requestContainer.AddBinding(
		(*Cache.ICache)(nil),
		Container2.NewBindingImpl(
			k.app.GetContainer().GetSingletonByAbstract((*Cache.ICache)(nil)),
		),
	)

	// session drive.
	requestContainer.AddBinding(
		(*Session.ISession)(nil),
		Container2.NewBindingImpl(
			k.app.GetContainer().GetSingletonByAbstract((*Session.ISession)(nil)),
		),
	)
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
