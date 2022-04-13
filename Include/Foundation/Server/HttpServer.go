package Server

import (
	"fmt"
	"net/http"
	"php-in-go/Bootstrap"
	"php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Foundation/Config"
	Http2 "php-in-go/Include/Http"
)

type HttpServer struct {
	port     int
	bashPath string
	app      Http.IApp
}

func (s *HttpServer) Initializer(config Config.HttpServer) {
	// set http server config
	s.port = config.Port
	s.bashPath = config.BasePath

	// set http kernel
	s.app = &Bootstrap.App{}
}

func (s *HttpServer) Start() {
	// set global container to http kernel
	s.app.Initializer(s)

	// start server.
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// start to dispatch.
		s.app.Handle(Http2.BuildRequest(request), Http2.BuildResponse(writer))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
	fmt.Println(err)
}
