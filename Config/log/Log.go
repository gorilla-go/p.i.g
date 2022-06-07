package log

func LogConfig() map[string]interface{} {
	return map[string]interface{}{
		"log":     true,
		"logPath": "Log/",
	}
}
