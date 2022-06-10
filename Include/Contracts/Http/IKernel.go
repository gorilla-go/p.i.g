package Http

import (
	"php-in-go/Include/Config"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Debug"
	"php-in-go/Include/Contracts/Http/Log"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Contracts/Routing"
	"php-in-go/Include/Http/Request"
	Http "php-in-go/Include/Http/Response"
)

type IKernel interface {
	Bootstrap(
		router Routing.IRouter,
		httpExceptionHandler Debug.IExceptionHandler,
		session Session.ISession,
		cache Cache.ICache,
		log Log.ILog,
		config Config.Loader,
	)
	Handle(request *Request.Request, response *Http.Response)
	Close()
}
