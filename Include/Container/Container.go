package Container

import (
	"fmt"
	"reflect"
)

type Container struct {
	singleton      map[string]interface{}
	bindings       map[string]*BindingImpl
	contextual     map[string]map[string]interface{}
	methodBindings map[string]func(params ...interface{})
	singletonAlias map[string]string
}

func NewContainer() *Container {
	return &Container{
		singleton:      make(map[string]interface{}),
		bindings:       make(map[string]*BindingImpl),
		contextual:     make(map[string]map[string]interface{}),
		methodBindings: make(map[string]func(params ...interface{})),
		singletonAlias: make(map[string]string),
	}
}

func (c *Container) AddBinding(abstract interface{}, elem *BindingImpl) {
	abstractName := GetPackageClassName(abstract)
	if _, exist := c.bindings[abstractName]; exist {
		return
	}
	c.bindings[abstractName] = elem
	if elem.GetShared() == true {
		var mapName string
		if elem.GetAlias() != "" {
			mapName = elem.GetAlias()
		}
		c.Singleton(elem.GetConcrete().Interface(), mapName)
	}
}

func (c *Container) Singleton(instance interface{}, alias string) {
	mapName := GetPackageClassName(instance)
	c.singleton[mapName] = instance
	if alias != "" {
		c.singletonAlias[alias] = mapName
	}
}

func (c *Container) GetInstanceByAlias(name string) interface{} {
	aliasName, ok := c.singletonAlias[name]
	if !ok {
		panic("none value")
	}

	inject, exist := c.singleton[aliasName]
	if !exist {
		panic("none value")
	}
	return inject
}

func (c *Container) GetInstanceByAbstract(abstract interface{}) interface{} {
	abstractName := GetPackageClassName(abstract)
	concrete := c.getConcrete(abstractName)
	if concrete == nil {
		return nil
	}
	return (*concrete).Interface()
}

func (c *Container) getConcrete(abstract string) *reflect.Value {
	if abstractElem, ok := c.bindings[abstract]; ok {
		if abstractElem.GetShared() == true {

			// search with alias
			if abstractElem.GetAlias() != "" {
				inject := reflect.ValueOf(c.GetInstanceByAlias(abstractElem.GetAlias()))
				return &inject
			}

			// search from instance
			mapName := GetPackageClassNameByRef(abstractElem.GetConcrete().Type())
			obj, exist := c.singleton[mapName]
			if !exist {
				panic("unknown error")
			}
			inject := reflect.ValueOf(obj)
			return &inject
		}
		con := abstractElem.GetConcrete()
		return &con
	}
	return nil
}

func (c *Container) Resolve(abstract interface{}, params map[string]interface{}, raiseEvents bool) interface{} {
	return c.resolveAbstract(
		reflect.TypeOf(abstract).Elem(),
		params,
		raiseEvents,
	)
}

func (c *Container) resolveAbstract(abstract reflect.Type, params map[string]interface{}, raiseEvents bool) interface{} {
	if raiseEvents == true {
		c.fireBeforeResolvingCallbacks(&abstract, &params)
	}

	// get strand package name.
	packagePath := GetPackageClassNameByRef(abstract)

	// has cache ?
	if object, exist := c.singleton[packagePath]; exist == true && len(params) == 0 {
		return object
	}

	// get interface to concrete.
	concrete := c.getConcrete(packagePath)

	// invalid interface ?
	if concrete == nil || reflect.TypeOf(*concrete).Kind() != reflect.Struct {
		panic("invalid abstract")
	}

	// ptr? to elem.
	if abstract.Kind() == reflect.Ptr {
		abstract = abstract.Elem()
	}

	// build it
	object := c.Build((*concrete).Interface(), params)

	// cache it
	if abstractElem, ok := c.bindings[packagePath]; ok && abstractElem.GetShared() == true && len(params) == 0 {
		c.singleton[packagePath] = object
	}

	if raiseEvents == true {
		c.fireAfterResolvingCallbacks(&abstract, &object)
	}
	return object
}

// Build class resolve params inject.
func (c *Container) Build(object interface{}, params map[string]interface{}) interface{} {
	// get scene class name.
	packageName := GetPackageClassName(object)

	// get reflect type and value
	refValue := reflect.ValueOf(object)
	refType := reflect.TypeOf(object)

	// check reflect type, is ptr? to elem.
	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()

		//// re-build memory
		//refValue = reflect.New(refType)
	}
	fmt.Println(refType.NumField())
	for i := 0; i < refType.NumField(); i++ {
		// get current field struct map.
		FieldStruct := refType.Field(i)
		pkgClassName := GetPackageClassNameByRef(FieldStruct.Type)

		// check field, ptr? to elem.
		var fieldValue reflect.Value
		if refValue.Kind() == reflect.Ptr {
			fieldValue = refValue.Elem().Field(i)
		} else {
			fieldValue = refValue.Field(i)
		}

		// solve from params
		if len(params) > 0 {
			if inject, ok := params[FieldStruct.Name]; ok {
				fieldValue.Set(reflect.ValueOf(inject))
				continue
			}
		}

		// contextual inject.
		if cacheInjectObj, exist := c.contextual[packageName][pkgClassName]; exist {
			fieldValue.Set(reflect.ValueOf(cacheInjectObj))
			continue
		}

		// get type from current fieldStruct
		fieldType := FieldStruct.Type

		// is ptr? to elem.
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		// solve from container
		switch fieldType.Kind() {
		case reflect.Interface:
			// search interface binding from storage. private?
			if bindingItem, exist := c.bindings[pkgClassName]; !exist || bindingItem.GetShared() != true {
				continue
			}

			// resolve struct
			resolveRes := reflect.ValueOf(
				c.resolveAbstract(
					fieldType,
					nil,
					false,
				),
			)

			if resolveRes.Type().Kind() == reflect.Ptr {
				resolveRes = reflect.Indirect(resolveRes)
			}
			fmt.Println("----")
			fmt.Println(resolveRes)
			fmt.Println(fieldValue.CanSet())
			// set field value
			fieldValue.Set(resolveRes)
			break
		case reflect.Struct:
			// search from cache
			if cacheType, ok := c.singleton[pkgClassName]; ok {
				cacheRef := reflect.ValueOf(cacheType)
				if cacheRef.Kind() == fieldValue.Kind() {
					fieldValue.Set(cacheRef)
				}
				break
			}

			// build in-time
			fieldValue.Set(
				reflect.ValueOf(
					c.Build(
						fieldValue.Interface(),
						nil,
					),
				),
			)
		}
	}

	return refValue.Interface()
}
