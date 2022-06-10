package Bootstrap

import (
	"php-in-go/App/Exception"
	Config2 "php-in-go/Config"
	Http2 "php-in-go/Include/Contracts/Http"
	Cache2 "php-in-go/Include/Foundation/Cache"
	Http3 "php-in-go/Include/Foundation/Http"
	Log2 "php-in-go/Include/Foundation/Http/Log"
	Session2 "php-in-go/Include/Foundation/Http/Session"
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Http/Response"
	Routing2 "php-in-go/Include/Routing"
)

type App struct {
	kernel Http2.IKernel
}

func (a *App) Initializer() {
	a.kernel = &Http3.Kernel{}

	// init http kernel.
	a.kernel.Bootstrap(
		&Routing2.Router{},
		&Exception.Handler{},
		&Session2.Session{},
		&Cache2.MemoryCache{},
		&Log2.Log{
			LogPath: Config2.Loader()["log.logPath"].(string),
		},
		Config2.Loader(),
	)
}

func (a *App) Handle(request *Request.Request, response *Response.Response) {
	a.kernel.Handle(request, response)
}

func (a *App) Close() {
	a.kernel.Close()
}
