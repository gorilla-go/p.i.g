package Container

import (
	Container2 "php-in-go/Include/Container"
)

type IContainer interface {
	AddBinding(abstract interface{}, elem *Container2.BindingImpl)
	Singleton(instance interface{}, alias string)
	Resolve(abstract interface{}, params map[string]interface{}, new bool) interface{}
	AddContextual(contextualElems ...*Container2.ContextualElem)
	GetSingleton(instance interface{}) interface{}
	GetSingletonByAlias(name string) interface{}
	GetSingletonByAbstract(abstract interface{}) interface{}
}
