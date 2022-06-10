package Config

import Config "php-in-go/Include/Config"

func Loader() Config.Loader {
	return Config.Loader{
		"app.port":  8084,
		"app.debug": true,

		// session.
		"app.sessionExpire": 60 * 60 * 24,
		"app.sessionKey":    "PHP_SSID",

		// cache
		"cache.defaultExpired": 60 * 60 * 24,

		// log
		"log.log":     true,
		"log.logPath": "Log/",

		// route
		// uri suffix. if empty no used.
		"route.urlHtmlSuffix": "html",

		// pjax request params tag.
		"route.pjaxName": "_pjax",
	}
}
