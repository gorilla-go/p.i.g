package Controller

import (
	"context"
	"fmt"
	"php-in-go/Include/Foundation/Http/Controller"
	"php-in-go/Include/Http"
)

type Index struct {
	*Controller.BaseController
}

func (t *Index) Index(request *Http.Request, response *Http.Response) *Http.Response {
	ctx := context.WithValue(request.Context(), "a", "bb")
	fmt.Println(ctx.Value("a"))
	return nil
}

func (t *Index) NoFound() *Http.Response {
	fmt.Println("index no found")
	return nil
}
