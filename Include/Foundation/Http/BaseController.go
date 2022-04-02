package Http

import (
	"context"
	"php-in-go/Include/Http"
)

type BaseController struct {
	Ctx      context.Context
	Request  *Http.Request
	Response *Http.Response
}
