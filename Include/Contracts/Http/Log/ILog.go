package Log

import Http2 "php-in-go/Include/Http"

type ILog interface {
	Log(request *Http2.Request, response *Http2.Response)
}
