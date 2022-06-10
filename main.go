package main

import (
	_ "net/http/pprof"
	"php-in-go/Bootstrap"
	Http2 "php-in-go/Include/Foundation/Server/HttpServer"
)

func main() {
	// set global container.
	httpServer := &Http2.HttpServer{
		App: &Bootstrap.App{},
	}

	// start to listen.
	httpServer.Start()
}
