package Controller

import (
	"fmt"
	"php-in-go/Include/Foundation/Http"
)

type Index struct {
	Http.BaseController
}

func (t *Index) Index() {
	fmt.Println("index controller")
}
