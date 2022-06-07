package Bootstrap

import (
	"php-in-go/App/Exception"
	"php-in-go/Config/app"
	"php-in-go/Config/cache"
	"php-in-go/Config/database"
	"php-in-go/Config/log"
	"php-in-go/Config/route"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Debug"
	Http2 "php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Contracts/Http/Log"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Contracts/Routing"
	Cache2 "php-in-go/Include/Foundation/Cache"
	Http3 "php-in-go/Include/Foundation/Http"
	Log2 "php-in-go/Include/Foundation/Http/Log"
	Session2 "php-in-go/Include/Foundation/Http/Session"
	"php-in-go/Include/Http"
	Routing2 "php-in-go/Include/Routing"
	"php-in-go/Routes"
)

type App struct {
	router           Routing.IRouter
	exceptionHandler Debug.IExceptionHandler
	kernel           Http2.IKernel
	session          Session.ISession
	cache            Cache.ICache
	log              Log.ILog
	config           map[string]map[string]interface{}
}

func (a *App) Initializer() {
	// set app config.
	a.config = map[string]map[string]interface{}{
		"app":      app.AppConfig(),
		"cache":    cache.CacheConfig(),
		"database": database.DatabaseConfig(),
		"log":      log.LogConfig(),
		"route":    route.RouteConfig(),
	}

	// set http kernel route.
	a.router = &Routing2.Router{}
	a.router.Initializer(Routes.Route(), route.RouteConfig())

	// set http exception handler.
	a.exceptionHandler = &Exception.Handler{}

	// cache.
	a.cache = &Cache2.MemoryCache{}
	a.cache.StartCacheManager()

	// set session driver.
	a.session = &Session2.Session{}
	a.session.StartSessionManager(a.cache, Session2.Config{
		Expire: a.config["app"]["sessionExpire"].(int),
		Name:   a.config["app"]["sessionKey"].(string),
	})

	// set log server.
	logConfig := log.LogConfig()
	a.log = &Log2.Log{
		LogPath: logConfig["logPath"].(string),
	}
	a.log.StartLogManager()

	// set http kernel.
	a.kernel = &Http3.Kernel{}
	a.kernel.Bootstrap(a)
}

func (a *App) Handle(request *Http.Request, response *Http.Response) {
	a.kernel.Handle(request, response)
}

func (a *App) Close() {
	a.session.CloseSessionManager()
	a.cache.CloseCacheManager()
	a.log.CloseLogManager()
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

func (a *App) GetConfigs() map[string]map[string]interface{} {
	return a.config
}
