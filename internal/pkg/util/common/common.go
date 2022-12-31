package common

import "reflect"

func Arguments(args ...interface{}) []interface{} {
	var result []interface{}
	for _, arg := range args {
		result = append(result, arg)
	}
	return result
}

func Value(model interface{}) interface{} {
	val := reflect.ValueOf(model).Elem()
	result := reflect.New(val.Type()).Elem()
	result.Set(val)
	return result.Interface()
}
