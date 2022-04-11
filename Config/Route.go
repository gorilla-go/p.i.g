package Config

// Route get route config.
func Route() map[string]interface{} {
	return map[string]interface{}{
		// uri suffix. if empty no used.
		"urlHtmlSuffix": "html",

		"pjaxName": "_pjax",
	}
}
