package Container

import "reflect"

func GetPackageClassName(abstract interface{}) string {
	return GetPackageClassNameByRef(reflect.TypeOf(abstract))
}

func GetPackageClassNameByRef(refType reflect.Type) string {
	isPtr := false
	if refType.Kind() == reflect.Ptr {
		isPtr = true
		refType = refType.Elem()
	}

	className := refType.Name()
	if isPtr && refType.Kind() != reflect.Interface {
		className = "*" + refType.Name()
	}
	return refType.PkgPath() + "@" + className
}
