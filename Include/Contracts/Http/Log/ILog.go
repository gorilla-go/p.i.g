package Log

import Http2 "php-in-go/Include/Http"

type ILog interface {
	StartLogManager()
	Log(request *Http2.Request, response *Http2.Response)
	CloseLogManager()
}
