package main

import (
	"os"
	Config2 "php-in-go/Config"
	"php-in-go/Include/Container"
	"php-in-go/Include/Contracts/Server"
	"php-in-go/Include/Foundation/Config"
	Http2 "php-in-go/Include/Foundation/Server"
)

func main() {
	basePath, pathError := os.Getwd()
	if pathError != nil {
		panic("unknown error.")
	}
	container := Container.NewContainer()

	// set global container.
	container.AddBinding((*Server.IServer)(nil), Container.NewBindingImpl(&Http2.HttpServer{}))
	httpServer := container.GetInstanceByAbstract((*Server.IServer)(nil)).(Server.IServer)

	// config.
	appConfig := Config2.App()

	// init http server.
	httpServer.Initializer(Config.HttpServer{
		BasePath:  basePath,
		Port:      appConfig["port"].(int),
		Container: container,
	})

	// start to listen.
	httpServer.Start()
}
