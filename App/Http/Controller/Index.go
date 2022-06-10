package Controller

import (
	"php-in-go/Include/Foundation/Http/Controller"
)

type Index struct {
	Controller.BaseController
}

func (t *Index) Index() {
	t.SetSession("app", "ok")
}

func (t *Index) Name() {
	t.GetSession("app")
}
