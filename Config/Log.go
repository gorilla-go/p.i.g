package Config

import Log2 "php-in-go/Include/Foundation/Http/Log"

func Log() map[string]interface{} {
	return map[string]interface{}{
		"log":       true,
		"logPath":   "Log",
		"logDriver": &Log2.Log{},
	}
}
