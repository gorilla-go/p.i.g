package Http

import (
	Container2 "php-in-go/Include/Container"
	Http2 "php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Contracts/Http/Controller"
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

// ServicesRegister init request instance bean.
func (k *Kernel) ServicesRegister() {}

// GetRequestContainer get current request container.
func (k *Kernel) GetRequestContainer() *Container2.Container {
	return k.requestContainer
}

// GetApp get global application.
func (k *Kernel) GetApp() Http2.IApp {
	return k.app
}

// Handle each request kernel.
func (k *Kernel) Handle(request *Http.Request, response *Http.Response) *Http.Response {
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

	// init context.
	k.ServicesRegister()

	// init background.
	actionTarget := k.app.GetRouter().Resolve(request)

	// call method.
	return k.dispatch(actionTarget, request, response)
}

// dispatch call correct controller method.
func (k *Kernel) dispatch(target *Routing.Target, request *Http.Request, response *Http.Response) *Http.Response {
	// page no found
	if target == nil {
		baseController := &Controller2.BaseController{
			Request:  request,
			Response: response,
		}
		return baseController.NoFound()
	}

	// resolve controller params.
	targetController := k.GetRequestContainer().Resolve(target.Controller, nil, true).(Controller.IController)

	// resolve target method
	targetMethod := reflect.ValueOf(targetController).MethodByName(target.Method)

	// no found method ? to NoFound method in base controller.
	if targetMethod.IsValid() == false {
		noFoundMethod := reflect.ValueOf(targetController).MethodByName("NoFound")
		responseArr := noFoundMethod.Call([]reflect.Value{})
		return responseArr[0].Interface().(*Http.Response)
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
	responseArr := targetMethod.Call(paramsArr)
	return responseArr[0].Interface().(*Http.Response)
}

func (k *Kernel) containerFoundation(request *Http.Request, response *Http.Response) {
	requestContainer := k.GetRequestContainer()

	// binding request.
	requestContainer.Singleton(request, "request")

	// binding response
	requestContainer.Singleton(response, "response")
}
