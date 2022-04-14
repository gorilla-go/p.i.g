package Server

import (
	"fmt"
	"net/http"
	"php-in-go/Include/Contracts/Http/App"
	"php-in-go/Include/Foundation/Config"
	Http2 "php-in-go/Include/Http"
)

type HttpServer struct {
	port     int
	bashPath string
	app      App.IApp
}

func (s *HttpServer) Initializer(config Config.HttpServer) {
	// set http server config
	s.port = config.Port
	s.bashPath = config.BasePath

	// set http kernel
	s.app = config.App
}

func (s *HttpServer) Start() {
	// set global container to http kernel
	s.app.Initializer()

	// start server.
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// start to dispatch.
		s.app.Handle(
			Http2.BuildRequest(
				request,
				s.app.GetConfigs(),
			),
			Http2.BuildResponse(
				writer,
				s.app.GetConfigs(),
			),
		)
	})

	// error?
	panic(http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil))
}
