package main

import (
	_ "net/http/pprof"
	"os"
	"php-in-go/Bootstrap"
	Config2 "php-in-go/Config"
	"php-in-go/Include/Foundation/Config"
	Http2 "php-in-go/Include/Foundation/Server"
)

func main() {
	basePath, pathError := os.Getwd()
	if pathError != nil {
		panic("unknown error.")
	}

	// config.
	appConfig := Config2.App()

	// set global container.
	httpServer := &Http2.HttpServer{}

	// init http server.
	httpServer.Initializer(Config.HttpServer{
		BasePath: basePath,
		Port:     appConfig["port"].(int),
		App:      &Bootstrap.App{},
	})

	// start to listen.
	httpServer.Start()
}
