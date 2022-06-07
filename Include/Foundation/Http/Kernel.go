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
	HttpPipeline "php-in-go/Include/Foundation/Http/Pipeline"
	"php-in-go/Include/Http"
	"reflect"
	"runtime"
	"time"
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
	defer k.exception(request, response)

	// fetch middleware
	var middlewareArr []func(*Http.Request, *Http.Response, func(*Http.Request, *Http.Response))
	for _, item := range Http3.Middlewares() {
		middlewareArr = append(middlewareArr, item.Handle)
	}

	// start pipe.
	HttpPipeline.NewHttpPipeline().Send(request, response).Through(middlewareArr...).Then(
		func(request *Http.Request, response *Http.Response) {
			k.dispatch(request, response)
		},
	)
}

func (k *Kernel) shellDump(request *Http.Request, response *Http.Response) {
	fmt.Printf(
		"%s [%d] %s %s %dms  %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		response.Code,
		request.Method,
		request.RequestURI,
		time.Now().Sub(request.StartTime).Microseconds(),
		response.ErrorMessage,
	)
}

func (k *Kernel) exception(request *Http.Request, response *Http.Response) {
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

	// dump into debug
	k.shellDump(request, response)

	// logger.
	k.app.GetLogger().Log(request, response)
}

// dispatch call correct controller method.
func (k *Kernel) dispatch(request *Http.Request, response *Http.Response) {
	// resolve container foundation
	container := k.containerFoundation(request, response)

	// route resolve.
	target := k.app.GetRouter().Resolve(request)

	// page no found
	if target == nil {
		k.pageNoFoundHandle(request, response)
		return
	}

	// resolve controller params.
	targetController := container.Resolve(target.Controller, nil, true).(Controller.IController)

	// resolve target method
	controllerRef := reflect.ValueOf(targetController)
	targetMethod := controllerRef.MethodByName(target.Method)

	// no found method ? to NoFound method in base controller.
	if targetMethod.IsValid() == false {
		k.pageNoFoundHandle(request, response)
		return
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

	// session.
	requestContainer.AddBinding(
		(*Session.ISession)(nil),
		Container2.NewBindingImpl(k.app.GetSession()).SetShared().SetAlias("session"),
	)

	// session.
	requestContainer.AddBinding(
		(*Cache.ICache)(nil),
		Container2.NewBindingImpl(k.app.GetCache()).SetShared().SetAlias("cache"),
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

	return requestContainer
}

func (k *Kernel) pageNoFoundHandle(request *Http.Request, response *Http.Response) {
	baseController := &Controller2.BaseController{
		Request:  request,
		Response: response,
	}
	baseController.NoFound()
}
