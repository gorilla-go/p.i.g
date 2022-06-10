package Http

import (
	"php-in-go/Include/Contracts/Middleware"
	"php-in-go/Include/Foundation/Http/Middleware/Console"
)

// Middlewares middleware config.
func Middlewares() []Middleware.IMiddleware {
	return []Middleware.IMiddleware{
		// log.
		&Console.LogMiddleware{},
	}
}
