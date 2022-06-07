package route

// RouteConfig get route config.
func RouteConfig() map[string]interface{} {
	return map[string]interface{}{
		// uri suffix. if empty no used.
		"urlHtmlSuffix": "html",

		// pjax request params tag.
		"pjaxName": "_pjax",
	}
}
