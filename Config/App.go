package Config

// App get app config.
func App() map[string]interface{} {
	return map[string]interface{}{
		// server.
		"port":    8084,
		"debug":   true,
		"logPath": "Log/",

		// session.
		"sessionExpire": 60 * 60 * 24,
		"sessionKey":    "PHP_SSID",
	}
}
