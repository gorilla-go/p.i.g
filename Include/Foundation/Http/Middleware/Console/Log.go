package Console

import (
	"fmt"
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Http/Response"
	"time"
)

type LogMiddleware struct{}

func (l *LogMiddleware) Handle(request *Request.Request, response *Response.Response, next func(request2 *Request.Request, response2 *Response.Response)) {
	next(request, response)

	// debug info
	fmt.Printf(
		"%s [%d] %s %s %dms  %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		response.Code,
		request.Method,
		request.RequestURI,
		time.Now().Sub(request.StartTime).Microseconds(),
		response.ErrorMessage,
	)
}
