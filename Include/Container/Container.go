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
		c.Singleton(elem.GetConcrete(), mapName)
	}
}

// Singleton Register singleton.
func (c *Container) Singleton(instance interface{}, alias string) {
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
	return concrete
}

func (c *Container) getConcrete(abstract string) interface{} {
	if abstractElem, ok := c.bindings[abstract]; ok {
		if abstractElem.GetShared() == true {

			// search with alias
			if abstractElem.GetAlias() != "" {
				return c.GetSingletonByAlias(abstractElem.GetAlias())
			}

			// search from instance
			return c.GetSingleton(abstractElem.GetConcrete())
		}
		return abstractElem.GetConcrete()
	}
	return nil
}

func (c *Container) Resolve(abstract interface{}, params map[string]interface{}, new bool) interface{} {
	// get reflection value.
	refType := reflect.TypeOf(abstract)

	switch refType.Kind() {

	// interface ?
	case reflect.Interface:
		return c.resolveAbstract(
			refType,
			params,
			new,
		)
	// struct
	case reflect.Struct:
		// search from cache
		cacheValue := c.GetSingleton(abstract)
		if cacheValue != nil && new == false && (params == nil || len(params) == 0) {
			return cacheValue
		}
		r := c.build(abstract, params)
		// cache
		if new == false && (params == nil || len(params) == 0) {
			c.Singleton(r, "")
		}
		return r
	case reflect.Func:
		return c.callFunc(abstract, params)
	case reflect.Ptr:
		// search from cache
		cacheValue := c.GetSingleton(abstract)
		if cacheValue != nil && new == false && (params == nil || len(params) == 0) {
			return cacheValue
		}
		// real-time resolve
		d := reflect.ValueOf(abstract)

		if new == true {
			d = reflect.New(d.Elem().Type())
		}
		if d.Elem().IsValid() == false || d.Elem().CanSet() == false {
			return abstract
		}
		o := c.Resolve(
			d.Elem().Interface(),
			params,
			new,
		)
		d.Elem().Set(reflect.ValueOf(o))
		r := d.Interface()
		// cache
		if new == false && (params == nil || len(params) == 0) {
			c.Singleton(r, "")
		}
		return r
	default:
		return abstract
	}
}

// callFunc resolve function and call with params.
func (c *Container) callFunc(abstract interface{}, params map[string]interface{}) []reflect.Value {
	// resolve target method
	targetMethod := reflect.ValueOf(abstract)

	// invalid method ?
	if targetMethod.IsValid() == false {
		panic("Invalid Method.")
	}

	// format params.
	var paramValues []reflect.Value
	for _, i := range params {
		paramValues = append(paramValues, reflect.ValueOf(i))
	}

	// resolve func params.
	var paramsArr []reflect.Value
	paramsNum := targetMethod.Type().NumIn()

	for i := 0; i < paramsNum; i++ {
		if len(paramValues) > i {
			paramsArr = append(paramsArr, paramValues[i])
			continue
		}

		// get param item type.
		paramItemType := targetMethod.Type().In(i)

		// is interface?
		if paramItemType.Kind() == reflect.Interface {
			paramsArr = append(
				paramsArr,
				reflect.ValueOf(
					c.resolveAbstract(paramItemType, nil, false),
				),
			)
			continue
		}

		// append to params arr.
		paramsArr = append(
			paramsArr,
			reflect.ValueOf(
				c.Resolve(reflect.New(paramItemType).Elem().Interface(), nil, false),
			),
		)
	}

	// try to call controller.
	return targetMethod.Call(paramsArr)
}

func (c *Container) resolveAbstract(abstract reflect.Type, params map[string]interface{}, new bool) interface{} {
	// get strand package name.
	packagePath := GetPackageClassNameByRef(abstract)

	// has cache ?
	v, exist := c.singletonAliasAbstract[packagePath]
	if exist && new == false && (params == nil || len(params) == 0) {
		return v
	}

	// get interface to concrete.
	concrete := c.getConcrete(packagePath)

	// invalid interface ?
	if concrete == nil {
		panic(errors.New("Unregistered interface mapping: " + abstract.String()))
	}

	// build object
	object := c.Resolve(concrete, params, new)

	// cache build result for next.
	c.singletonAliasAbstract[packagePath] = object

	// shared ?
	if c.bindings[packagePath].shared == true && new == false && (params == nil || len(params) == 0) {
		c.Singleton(object, c.bindings[packagePath].alias)
	}

	return object
}

// build  resolve struct.
func (c *Container) build(object interface{}, params map[string]interface{}) interface{} {
	refType := reflect.TypeOf(object)

	// get scene class name.
	packageName := GetPackageClassNameByRef(refType)

	// re-construct struct for memory
	refValue := reflect.New(refType).Elem()

	for i := 0; i < refValue.NumField(); i++ {
		// get current field struct map.
		fieldStruct := refValue.Type().Field(i)

		fieldValue := refValue.Field(i)

		// is exported?
		if fieldStruct.IsExported() == false || fieldValue.CanSet() == false {
			continue
		}

		// inject tag
		if tag, ok := fieldStruct.Tag.Lookup("inject"); ok && tag == "false" {
			continue
		}

		// solve from params
		if len(params) > 0 {
			if inject, ok := params[fieldStruct.Name]; ok {
				fieldValue.Set(reflect.ValueOf(inject))
				continue
			}
		}

		// contextual inject.
		fieldName := GetPackageClassNameByRef(fieldStruct.Type)
		if cacheInjectObj, exist := c.contextual[packageName][fieldName]; exist {
			fieldValue.Set(reflect.ValueOf(cacheInjectObj))
			continue
		}

		if fieldValue.Type().Kind() == reflect.Interface {
			fieldValue.Set(
				reflect.ValueOf(
					c.resolveAbstract(
						fieldValue.Type(),
						nil,
						false,
					),
				),
			)
			continue
		}

		// set attribute
		if tag, ok := fieldStruct.Tag.Lookup("def"); ok {
			refDefaultValue := reflect.ValueOf(tag)
			if refDefaultValue.CanConvert(fieldStruct.Type) {
				fieldValue.Set(refDefaultValue.Convert(fieldStruct.Type))
				continue
			}
		}

		// build in-time
		fieldValue.Set(
			reflect.ValueOf(
				c.Resolve(
					fieldValue.Interface(),
					nil,
					false,
				),
			),
		)
	}

	return refValue.Interface()
}
