package Container

import (
	"errors"
	"reflect"
)

type Container struct {
	singleton              map[string]interface{}
	bindings               map[string]*BindingImpl
	contextual             map[string]map[string]interface{}
	singletonAlias         map[string]string
	singletonAliasAbstract map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		singleton:              make(map[string]interface{}),
		bindings:               make(map[string]*BindingImpl),
		contextual:             make(map[string]map[string]interface{}),
		singletonAlias:         make(map[string]string),
		singletonAliasAbstract: make(map[string]interface{}),
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

// Singleton Register singleton.
func (c *Container) Singleton(instance interface{}, alias string) {
	// un-support non-pointer type.
	typeInstance := reflect.TypeOf(instance).Kind()
	if typeInstance != reflect.Ptr {
		panic(errors.New("Non-pointer type: " + typeInstance.String()))
	}
	mapName := GetPackageClassName(instance)
	c.singleton[mapName] = instance
	if alias != "" {
		c.singletonAlias[alias] = mapName
	}
}

// GetSingleton Get singleton.
func (c *Container) GetSingleton(instance interface{}) interface{} {
	mapName := GetPackageClassName(instance)
	if v, exist := c.singleton[mapName]; exist {
		return v
	}
	return nil
}

func (c *Container) GetSingletonByAlias(name string) interface{} {
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

func (c *Container) GetSingletonByAbstract(abstract interface{}) interface{} {
	concrete := c.getConcrete(GetPackageClassName(abstract))
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
				inject := reflect.ValueOf(c.GetSingletonByAlias(abstractElem.GetAlias()))
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

func (c *Container) Resolve(abstract interface{}, params map[string]interface{}, new bool) (r interface{}) {
	c.fireBeforeResolvingCallbacks(&abstract, &params)

	// get reflection value.
	refType := reflect.TypeOf(abstract)
	originRefType := refType

	// is ptr
	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
	}

	switch refType.Kind() {

	// interface ?
	case reflect.Interface:
		r = c.resolveAbstract(
			refType,
			params,
			new,
		)
		break
	// struct
	case reflect.Struct:
		r = c.build(abstract, params, new)
		break
	default:
		r = reflect.New(originRefType).Elem()
		break
	}

	// after resolving callback
	c.fireAfterResolvingCallbacks(&r, abstract)

	return
}

func (c *Container) resolveAbstract(abstract reflect.Type, params map[string]interface{}, new bool) interface{} {
	// get strand package name.
	packagePath := GetPackageClassNameByRef(abstract)

	// has cache ?
	if v, exist := c.singletonAliasAbstract[packagePath]; exist && new == false {
		return v
	}

	// get interface to concrete.
	concrete := c.getConcrete(packagePath)

	// invalid interface ?
	if concrete == nil {
		panic(errors.New("Unregistered interface mapping: " + abstract.String()))
	}

	// build object
	object := c.build((*concrete).Interface(), params, new)

	// cache build result for next.
	c.singletonAliasAbstract[packagePath] = object

	// shared ?
	if c.bindings[packagePath].shared == true {
		c.Singleton(object, c.bindings[packagePath].alias)
	}

	return object
}

// build  resolve params inject.
func (c *Container) build(object interface{}, params map[string]interface{}, new bool) interface{} {
	// search from cache
	cacheValue := c.GetSingleton(object)
	if cacheValue != nil || new == false {
		return cacheValue
	}

	// get scene class name.
	packageName := GetPackageClassName(object)

	// get reflection type for next.
	reflectionType := reflect.TypeOf(object)
	// get reflect type and value
	var refValue reflect.Value

	isPtr := false
	if reflectionType.Kind() == reflect.Ptr {
		isPtr = true
		reflectionType = reflectionType.Elem()
	}

	// re-construct struct for memory
	refValue = reflect.New(reflectionType).Elem()

	for i := 0; i < refValue.NumField(); i++ {
		// get current field struct map.
		fieldStruct := refValue.Type().Field(i)

		// is exported?
		if fieldStruct.IsExported() == false {
			continue
		}
		fieldName := GetPackageClassNameByRef(fieldStruct.Type)

		// check field, ptr? to elem.
		var fieldValue reflect.Value
		if refValue.Kind() == reflect.Ptr {
			fieldValue = refValue.Elem().Field(i)
		} else {
			fieldValue = refValue.Field(i)
		}

		// solve from params
		if len(params) > 0 {
			if inject, ok := params[fieldStruct.Name]; ok {
				fieldValue.Set(reflect.ValueOf(inject))
				continue
			}
		}

		// contextual inject.
		if cacheInjectObj, exist := c.contextual[packageName][fieldName]; exist {
			fieldValue.Set(reflect.ValueOf(cacheInjectObj))
			continue
		}

		// get type from current fieldStruct
		fieldType := fieldStruct.Type

		// is ptr? to elem.
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		// solve from container
		switch fieldType.Kind() {
		case reflect.Interface:
			if fieldValue.CanSet() == true {
				// search interface binding from storage.
				cacheImplValue := c.getConcrete(GetPackageClassNameByRef(fieldType))
				if cacheImplValue != nil {
					fieldValue.Set(*cacheImplValue)
					continue
				}

				// set field value
				fieldValue.Set(reflect.ValueOf(
					c.resolveAbstract(
						fieldType,
						nil,
						new,
					),
				))
			}
			break
		case reflect.Struct:
			if fieldValue.CanSet() == true {
				// build in-time
				fieldValue.Set(
					reflect.ValueOf(
						c.build(
							fieldValue.Interface(),
							nil,
							new,
						),
					),
				)
			}
		}
	}

	if isPtr {
		return refValue.Addr().Interface()
	}

	// cache
	c.Singleton(refValue.Interface(), "")
	return refValue.Interface()
}
