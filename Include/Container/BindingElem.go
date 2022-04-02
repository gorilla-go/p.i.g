package Container

import "reflect"

type BindingImpl struct {
	concrete reflect.Value
	shared   bool
	alias    string
}

func NewBindingImpl(concrete interface{}) *BindingImpl {
	return &BindingImpl{
		concrete: reflect.ValueOf(concrete),
		shared:   false,
		alias:    "",
	}
}

func (e *BindingImpl) GetConcrete() reflect.Value {
	return e.concrete
}

func (e *BindingImpl) SetShared() *BindingImpl {
	e.shared = true
	return e
}

func (e *BindingImpl) GetShared() bool {
	return e.shared
}

func (e *BindingImpl) SetAlias(name string) *BindingImpl {
	e.alias = name
	return e
}

func (e *BindingImpl) GetAlias() string {
	return e.alias
}
