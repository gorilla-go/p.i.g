package Config

import (
	"php-in-go/Include/Contracts/Container"
)

type HttpServer struct {
	BasePath  string
	Port      int
	Container Container.IContainer
}
