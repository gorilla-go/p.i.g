package Http

import (
	"context"
	"errors"
	"fmt"
	"log"
	Container2 "php-in-go/Include/Container"
	"php-in-go/Include/Contracts/Container"
	Http2 "php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Http"
	"reflect"
	"time"
)

type Kernel struct {
	app              Http2.IApp
	requestContainer *Container2.Container
}

func (k *Kernel) Bootstrap(app Http2.IApp) {
	k.app = app
}

// RequestInstance init request instance bean.
func (k *Kernel) RequestInstance() {}

// GetRequestInstance get current request container.
func (k *Kernel) GetRequestInstance() *Container2.Container {
	return k.requestContainer
}

// GetApp get global application.
func (k *Kernel) GetApp() Http2.IApp {
	return k.app
}

// Handle each request kernel.
func (k *Kernel) Handle(request *Http.Request) (response *Http.Response) {
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		v := make([]byte, 1024)
	//		runtime.Stack(v, true)
	//		fmt.Println(string(v))
	//	}
	//}()

	// init request container
	k.requestContainer = Container2.NewContainer()

	// resolve container foundation
	k.containerFoundation(request, response)

	// init context.
	k.RequestInstance()

	// init background.
	actionTarget := k.app.GetRouter().Resolve(request)

	// page no found
	if actionTarget == nil {
		fmt.Println("404. route not found.")
	}

	// resolve controller params.
	targetController := k.GetRequestInstance().Build(actionTarget.Controller, nil)
	fmt.Println(targetController)
	// controller check.
	if reflect.TypeOf(targetController).Implements(reflect.TypeOf((*Container.IContainer)(nil))) == false {
		fmt.Println(11)
		log.Fatalln("Invalid Controller.")
	}

	// resolve target method
	_ = reflect.ValueOf(targetController).MethodByName(actionTarget.Method)
	return
}

type test struct {
	A int8
}

func (t *test) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (t *test) Done() <-chan struct{} {
	return make(chan struct{})
}
func (t *test) Err() error {
	return errors.New("")
}
func (t *test) Value(key interface{}) interface{} {
	return ""
}

func (k *Kernel) containerFoundation(request *Http.Request, response *Http.Response) {
	requestContainer := k.GetRequestInstance()

	// binding context
	// ctx := context.Background()
	requestContainer.AddBinding(
		new(context.Context),
		Container2.NewBindingImpl(&test{A: 9}).SetShared().SetAlias("ctx"),
	)

	// binding request.
	requestContainer.Singleton(request, "request")

	// binding response
	requestContainer.Singleton(response, "response")
}
