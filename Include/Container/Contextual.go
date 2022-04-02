package Container

import (
	"reflect"
)

type ContextualElem struct {
	scene  interface{}
	target interface{}
	origin interface{}
}

func (c *Container) AddContextual(contextualElems ...*ContextualElem) {
	for _, elem := range contextualElems {
		// map scene
		scenePkgName := GetPackageClassName(elem.scene)

		// map target
		targetPkgName := GetPackageClassName(elem.target)

		if reflect.ValueOf(c.contextual[scenePkgName]).IsNil() {
			c.contextual[scenePkgName] = make(map[string]interface{})
		}
		c.contextual[scenePkgName][targetPkgName] = elem.origin
	}
}

func NewContextualElem() *ContextualElem {
	return &ContextualElem{}
}

func (c *ContextualElem) When(abstract interface{}) *ContextualElem {
	c.scene = abstract
	return c
}

func (c *ContextualElem) Need(abstract interface{}) *ContextualElem {
	c.target = abstract
	return c
}

func (c *ContextualElem) Give(abstract interface{}) *ContextualElem {
	c.origin = abstract
	return c
}
