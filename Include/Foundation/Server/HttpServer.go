package Server

import (
	"fmt"
	"net/http"
	"php-in-go/Bootstrap"
	Container2 "php-in-go/Include/Container"
	"php-in-go/Include/Contracts/Container"
	"php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Foundation/Config"
	Http2 "php-in-go/Include/Http"
)

type HttpServer struct {
	port      int
	bashPath  string
	app       Http.IApp
	container Container.IContainer
}

func (s *HttpServer) Initializer(config Config.HttpServer) {
	// set http server config
	s.port = config.Port
	s.bashPath = config.BasePath
	s.container = config.Container

	// set http kernel
	s.container.AddBinding((*Http.IApp)(nil), Container2.NewBindingImpl(&Bootstrap.App{}))
	s.app = s.container.GetInstanceByAbstract((*Http.IApp)(nil)).(Http.IApp)
}

func (s *HttpServer) GetContainer() Container.IContainer {
	return s.container
}

func (s *HttpServer) Start() {
	// set global container to http kernel
	s.app.Initializer(s)

	// handle logo.
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte(""))
	})

	// start server.
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		go func() {
			s.app.Handle(Http2.BuildRequest(request))
		}()
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
	fmt.Println(err)
}
