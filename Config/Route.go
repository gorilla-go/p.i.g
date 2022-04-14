package Config

import Routing2 "php-in-go/Include/Routing"

// Route get route config.
func Route() map[string]interface{} {
	return map[string]interface{}{
		// route driver. impl IRoute.
		"routeDriver": &Routing2.Router{},

		// uri suffix. if empty no used.
		"urlHtmlSuffix": "html",

		// pjax request params tag.
		"pjaxName": "_pjax",
	}
}
