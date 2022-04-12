package Http

import (
	"php-in-go/Include/Contracts/Middleware"
	Middleware2 "php-in-go/Include/Middleware"
)

// Middlewares middleware config.
func Middlewares() []Middleware.IMiddleware {
	return []Middleware.IMiddleware{
		&Middleware2.SessionMiddleware{},
	}
}
