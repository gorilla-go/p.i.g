package Controller

import "php-in-go/Include/Http"

type IController interface {
	NoFound() *Http.Response
}
