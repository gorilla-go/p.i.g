package Middleware

import "php-in-go/Include/Http"

type Session struct {
}

func (s *Session) Handle(request *Http.Request, response *Http.Response, next func(request2 *Http.Request, response2 *Http.Response)) {
	//TODO implement me
	panic("implement me")
}
