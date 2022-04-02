package Container

import (
	Container2 "php-in-go/Include/Container"
)

type IContainer interface {
	AddBinding(abstract interface{}, elem *Container2.BindingImpl)
	Singleton(instance interface{}, alias string)
	Resolve(abstract interface{}, params map[string]interface{}, raiseEvents bool) interface{}
	Build(object interface{}, params map[string]interface{}) interface{}
	AddContextual(contextualElems ...*Container2.ContextualElem)
	GetInstanceByAlias(name string) interface{}
	GetInstanceByAbstract(abstract interface{}) interface{}
}
