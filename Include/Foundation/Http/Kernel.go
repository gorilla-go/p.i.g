package Http

import (
	"fmt"
	Http3 "php-in-go/App/Http"
	"php-in-go/Include/Config"
	Container2 "php-in-go/Include/Container"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Container"
	"php-in-go/Include/Contracts/Debug"
	"php-in-go/Include/Contracts/Http/Controller"
	"php-in-go/Include/Contracts/Http/Log"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Contracts/Routing"
	"php-in-go/Include/Foundation/Exceptions"
	Controller2 "php-in-go/Include/Foundation/Http/Controller"
	HttpPipeline "php-in-go/Include/Foundation/Http/Pipeline"
	Session2 "php-in-go/Include/Foundation/Http/Session"
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Http/Response"
	"php-in-go/Routes"
	"reflect"
	"runtime"
)

type Kernel struct {
	router               Routing.IRouter
	httpExceptionHandler Debug.IExceptionHandler
	session              Session.ISession
	cache                Cache.ICache
	log                  Log.ILog
	config               Config.Loader
}

func (k *Kernel) Bootstrap(
	router Routing.IRouter,
	httpExceptionHandler Debug.IExceptionHandler,
	session Session.ISession,
	cache Cache.ICache,
	log Log.ILog,
	config Config.Loader,
) {
	k.config = config

	// set route.
	k.router = router
	k.router.Initializer(Routes.Route(), config.LoadPath("route"))

	// set cache.
	k.cache = cache
	k.cache.StartCacheManager()

	// set http base exception handler.
	k.httpExceptionHandler = httpExceptionHandler

	// init session
	k.session = session
	k.session.StartSessionManager(k.cache, Session2.Config{
		Expire: k.config["app.sessionExpire"].(int),
		Name:   k.config["app.sessionKey"].(string),
	})

	// set log server.
	k.log = log
	log.StartLogManager()
}

// Handle each request kernel.
func (k *Kernel) Handle(request *Request.Request, response *Response.Response) {
	// register error handle.
	defer func() {
		// handler exception
		k.exception(request, response)

		// set log.
		go k.doLog(request, response)
	}()

	// start session
	k.session.SessionStart(request, response)

	// fetch middleware
	var middlewareArr []func(*Request.Request, *Response.Response, func(*Request.Request, *Response.Response))
	for _, item := range Http3.Middlewares() {
		middlewareArr = append(middlewareArr, item.Handle)
	}

	// start pipe.
	HttpPipeline.NewHttpPipeline().Send(request, response).Through(middlewareArr...).Then(
		func(request *Request.Request, response *Response.Response) {
			k.dispatch(request, response)
		},
	)
}

func (k *Kernel) Close() {
	k.session.CloseSessionManager()
	k.cache.CloseCacheManager()
	k.log.CloseLogManager()
}

func (k *Kernel) exception(request *Request.Request, response *Response.Response) {
	// catch exception.
	err := recover()
	if err != nil {
		v := make([]byte, 1024*2)
		runtime.Stack(v, true)
		k.httpExceptionHandler.Handle(
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
}

func (k *Kernel) doLog(request *Request.Request, response *Response.Response) {
	k.log.Log(request, response)
}

// dispatch call correct controller method.
func (k *Kernel) dispatch(request *Request.Request, response *Response.Response) {
	// resolve container foundation
	container := k.requestContainer(request, response)

	// route resolve.
	target := k.router.Resolve(request)

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
	method := targetMethod.Interface()
	container.Resolve(method, nil, true)
}

func (k *Kernel) requestContainer(
	request *Request.Request,
	response *Response.Response,
) Container.IContainer {
	// init basic container
	basicContainer := Container2.NewContainer()

	// set base.
	basicContainer.Singleton(k.session, "session")
	basicContainer.Singleton(k.config, "config")
	basicContainer.Singleton(k.cache, "cache")

	// request container.
	basicContainer.AddBinding(
		(*Container.IContainer)(nil),
		Container2.NewBindingImpl(basicContainer).SetShared().SetAlias("container"),
	)

	// binding request.
	basicContainer.Singleton(request, "request")

	// binding response
	basicContainer.Singleton(response, "response")

	// return container.
	return basicContainer
}

func (k *Kernel) pageNoFoundHandle(request *Request.Request, response *Response.Response) {
	baseController := &Controller2.BaseController{
		Request:  request,
		Response: response,
	}
	baseController.NoFound()
}
