package Config

import (
	"php-in-go/Include/Contracts/Http/App"
)

type HttpServer struct {
	BasePath string
	Port     int
	App      App.IApp
}
