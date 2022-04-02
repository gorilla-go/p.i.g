package Server

import (
	"php-in-go/Include/Foundation/Config"
)

type IServer interface {
	Initializer(config Config.HttpServer)
	Start()
}
