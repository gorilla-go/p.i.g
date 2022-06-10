package HttpServer

import (
	"fmt"
	"net/http"
	"php-in-go/Config"
	"php-in-go/Include/Contracts/Http/App"
	"php-in-go/Include/Http/Request"
	Http2 "php-in-go/Include/Http/Response"
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
			Request.BuildRequest(request),
			Http2.BuildResponse(writer),
		)
	})

	// error?
	panic(
		http.ListenAndServe(
			fmt.Sprintf(":%d", Config.Loader()["app.port"].(int)),
			nil,
		),
	)
}
