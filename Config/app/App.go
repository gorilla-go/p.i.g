package app

// App get app config.
func AppConfig() map[string]interface{} {
	return map[string]interface{}{
		// server.
		"port":  8084,
		"debug": true,

		// session.
		"sessionExpire": 60 * 60 * 24,
		"sessionKey":    "PHP_SSID",
	}
}