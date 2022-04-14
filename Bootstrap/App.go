package Bootstrap

import (
	"php-in-go/Config"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Debug"
	Http2 "php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Contracts/Http/Log"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Contracts/Routing"
	"php-in-go/Include/Http"
	"php-in-go/Routes"
)

type App struct {
	router           Routing.IRouter
	exceptionHandler Debug.IExceptionHandler
	kernel           Http2.IKernel
	session          Session.ISession
	cache            Cache.ICache
	log              Log.ILog
	config           map[string]interface{}
}

func (a *App) Initializer() {
	// set app config.
	a.config = Config.App()

	// set http kernel route.
	a.router = a.config["routeDriver"].(Routing.IRouter)
	a.router.Initializer(Routes.Route(), a.config)

	// set http exception handler.
	a.exceptionHandler = a.config["exceptionHandleDriver"].(Debug.IExceptionHandler)

	// cache.
	a.cache = a.config["cacheDriver"].(Cache.ICache)
	a.cache.StartCacheManager()

	// set session driver.
	a.session = a.config["sessionDriver"].(Session.ISession)
	a.session.StartSessionManager(a.cache)

	// set log server.
	a.log = a.config["logDriver"].(Log.ILog)
	a.log.StartLogManager()

	// set http kernel.
	a.kernel = a.config["kernelDriver"].(Http2.IKernel)
	a.kernel.Bootstrap(a)
}

func (a *App) Handle(request *Http.Request, response *Http.Response) {
	a.kernel.Handle(request, response)
}

func (a *App) GetRouter() Routing.IRouter {
	return a.router
}

func (a *App) GetLogger() Log.ILog {
	return a.log
}

func (a *App) GetSession() Session.ISession {
	return a.session
}

func (a *App) GetCache() Cache.ICache {
	return a.cache
}

func (a *App) GetExceptionHandler() Debug.IExceptionHandler {
	return a.exceptionHandler
}

func (a *App) GetConfigs() map[string]interface{} {
	return a.config
}
