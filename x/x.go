package x

import (
	"fmt"
	"reflect"
)

func StringerToString(slice interface{}) []string {
	sliceReflect := reflect.ValueOf(slice)
	if sliceReflect.Kind() != reflect.Slice {
		panic("StringerToSlice provided non-slice value")
	}

	if sliceReflect.IsNil() {
		return []string{}
	}

	stringSlice := make([]string, sliceReflect.Len())
	for i := 0; i < sliceReflect.Len(); i++ {
		if str, ok := sliceReflect.Index(i).Interface().(fmt.Stringer); ok {
			stringSlice[i] = str.String()
		} else {
			panic("StrignerToSlice provieded non-stringer element")
		}
	}
	return stringSlice
}
