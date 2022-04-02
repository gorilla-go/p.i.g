package Config

// App get app config.
func App() map[string]interface{} {
	return map[string]interface{}{
		// server start port
		"port": 8084,
	}
}
