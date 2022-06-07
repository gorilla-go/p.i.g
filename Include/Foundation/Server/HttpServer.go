package Server

import (
	"fmt"
	"net/http"
	"php-in-go/Include/Contracts/Http/App"
	Http2 "php-in-go/Include/Http"
)

type HttpServer struct {
	App App.IApp
}

func (s *HttpServer) Start() {
	// set global container to http kernel
	s.App.Initializer()
	defer s.App.Close()

	// start server.
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// start to dispatch.
		s.App.Handle(
			Http2.BuildRequest(request),
			Http2.BuildResponse(writer),
		)
	})

	// error?
	globalConfig := s.App.GetConfigs()
	panic(http.ListenAndServe(fmt.Sprintf(":%d", globalConfig["app"]["port"].(int)), nil))
}
