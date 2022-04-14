package Debug

import (
	"fmt"
	"php-in-go/Include/Container"
	"reflect"
	"strconv"
	"strings"
)

// Dump format interface to text
func Dump(i interface{}, level int) string {
	inputType := typeOf(i)
	switch inputType.Kind() {
	case reflect.Interface:
		return inputType.Kind().String()
	case reflect.Ptr:
		return "*" + Dump(reflect.ValueOf(i).Elem().Interface(), level)
	case reflect.Invalid:
		return "nil (Invalid)"
	case reflect.String:
		return fmt.Sprintf("%#v", i) + " (String)"
	case reflect.Int:
		return strconv.Itoa(i.(int)) + " (Int)"
	case reflect.Struct:
		refValue := reflect.ValueOf(i)
		refType := reflect.TypeOf(i)
		fieldNum := refValue.NumField()
		dumpStr := Container.GetPackageClassName(i) + " (Struct)"
		for ii := 0; ii < fieldNum; ii++ {
			fieldType := refType.Field(ii)
			if fieldType.IsExported() == false {
				continue
			}
			tmpStr := ""
			if refValue.Field(ii).IsValid() && refValue.Field(ii).IsZero() == false {
				tmpStr = Dump(refValue.Field(ii).Interface(), level+1)
			} else {
				tmpStr = "nil (" + fieldType.Type.Kind().String() + ")"
			}

			dumpStr = fmt.Sprintf(
				"%s<br>%s%s%s",
				dumpStr,
				strings.Repeat("&nbsp&nbsp&nbsp&nbsp", level),
				fieldType.Name+": ",
				tmpStr,
			)
		}
		return dumpStr
	case reflect.Func:
		return inputType.String() + " (Func)"
	case reflect.Array:
		refValue := reflect.ValueOf(i)
		dumpStr := refValue.Type().String()
		for i := 0; i < refValue.Len(); i++ {
			dumpStr = fmt.Sprintf(
				"%s<br>%s%s",
				dumpStr,
				strings.Repeat("&nbsp&nbsp&nbsp&nbsp", level),
				Dump(refValue.Index(i).Interface(), level+1),
			)
		}
		return dumpStr
	case reflect.Bool:
		str := "true"
		if i.(bool) == false {
			str = "false"
		}
		return str + " (Bool)"
	case reflect.Chan:
		return inputType.String() + " (Chan)"
	case reflect.Complex64:
		return fmt.Sprintf("%s", i) + " (Complex64)"
	case reflect.Complex128:
		return fmt.Sprintf("%s", i) + " (Complex128)"
	case reflect.Float32:
		return strconv.FormatFloat(i.(float64), 'E', -1, 32) + " (Float32)"
	case reflect.Float64:
		return strconv.FormatFloat(i.(float64), 'E', -1, 64) + " (Float64)"
	case reflect.Int8:
		return fmt.Sprintf("%d", i) + " (Int8)"
	case reflect.Int16:
		return fmt.Sprintf("%d", i) + " (Int16)"
	case reflect.Int32:
		return fmt.Sprintf("%d", i) + " (Int32)"
	case reflect.Int64:
		return fmt.Sprintf("%d", i) + " (Int64)"
	case reflect.Map:
		refValue := reflect.ValueOf(i)
		dumpStr := refValue.Type().String() + " (Map)"
		for _, key := range refValue.MapKeys() {
			dumpStr = fmt.Sprintf(
				"%s<br>%s%s%s",
				dumpStr,
				strings.Repeat("&nbsp&nbsp&nbsp&nbsp", level),
				fmt.Sprintf("%#v", key.Interface())+": ",
				Dump(refValue.MapIndex(key).Interface(), level+1),
			)
		}
		return dumpStr
	case reflect.Slice:
		dumpStr := "Slice:"
		refValue := reflect.ValueOf(i)
		for i := 0; i < refValue.Len(); i++ {
			dumpStr = fmt.Sprintf(
				"%s%s<br>%s%s",
				strings.Repeat("&nbsp&nbsp&nbsp&nbsp", level),
				dumpStr,
				strings.Repeat("&nbsp&nbsp&nbsp&nbsp", level+1),
				Dump(refValue.Index(i).Interface(), level+1),
			)
		}
		return dumpStr
	case reflect.Uint:
		return fmt.Sprintf("%d", i) + " (Uint)"
	case reflect.Uint8:
		return fmt.Sprintf("%d", i) + " (Uint8)"
	case reflect.Uint16:
		return fmt.Sprintf("%d", i) + " (Uint16)"
	case reflect.Uint32:
		return fmt.Sprintf("%d", i) + " (Uint32)"
	case reflect.Uint64:
		return fmt.Sprintf("%d", i) + " (Uint64)"
	case reflect.Uintptr:
		return fmt.Sprintf("%d", i) + " (Uintptr)"
	case reflect.UnsafePointer:
		return fmt.Sprintf("%p", i) + " (UnsafePointer)"
	default:
		return fmt.Sprintf("%s", i) + " (" + inputType.Kind().String() + ")"
	}
}

func typeOf(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}
