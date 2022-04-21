package Config

import (
	"php-in-go/App/Exception"
	"php-in-go/Include/Foundation/Http"
	"php-in-go/Include/Foundation/Http/Session"
	"php-in-go/Include/Util/Map"
)

// App get app config.
func App() map[string]interface{} {
	return Map.Merge(
		map[string]interface{}{
			// server.
			"port":  8084,
			"debug": true,

			// session.
			"sessionExpire": 60 * 60 * 24,
			"sessionKey":    "PHP_SSID",
			"sessionDriver": &Session.Session{},

			// kernel driver.
			"kernelDriver": &Http.Kernel{},

			// exception handler.
			"exceptionHandleDriver": &Exception.Handler{},
		},
		Route(),
		Cache(),
		Log(),
		Database(),
	)
}
