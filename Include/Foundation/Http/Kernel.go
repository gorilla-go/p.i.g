package Http

import (
	"fmt"
	"php-in-go/App/Http/Middleware"
	Container2 "php-in-go/Include/Container"
	"php-in-go/Include/Contracts/Container"
	Http2 "php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Contracts/Http/Controller"
	"php-in-go/Include/Contracts/Http/Session"
	Controller2 "php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
	"php-in-go/Include/Routing"
	"reflect"
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
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		v := make([]byte, 1024)
	//		runtime.Stack(v, true)
	//		fmt.Println(string(v))
	//	}
	//}()
	// init response

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

	// resolve func params.
	var paramsArr []reflect.Value
	paramsNum := targetMethod.Type().NumIn()

	for i := 0; i < paramsNum; i++ {
		// get param item type.
		paramItemType := targetMethod.Type().In(i)

		// is interface?
		if paramItemType.Kind() == reflect.Interface {
			paramItemType = reflect.PtrTo(paramItemType)
		}

		// append to params arr.
		paramsArr = append(
			paramsArr,
			reflect.ValueOf(
				k.requestContainer.Resolve(reflect.New(paramItemType).Elem().Interface(), nil, true),
			),
		)
	}

	// try to call controller.
	targetMethod.Call(paramsArr)
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

	// session drive.
	requestContainer.AddBinding(
		(*Session.ISession)(nil),
		Container2.NewBindingImpl(
			k.app.GetContainer().GetSingletonByAbstract((*Session.ISession)(nil)),
		),
	)
}

func (k *Kernel) middlewareHandler(target *Routing.Target, request *Http.Request, response *Http.Response) bool {
	middlewares := Middleware.Middlewares()
	if len(middlewares) > 0 {
		for _, middleware := range middlewares {
			if middleware.Handle(request, response, target) == false {
				return false
			}
		}
	}
	return true
}
