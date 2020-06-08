package main

import (
	"fmt"
	"reflect"
)
//
//func walk(x interface{}, fn func(input string))  {
//	val := getValue(x)
//	if val.Kind() == reflect.Slice {
//		for i:=0; i < val.Len(); i++ {
//			walk(val.Index(i).Interface(), fn)
//		}
//		return
//	}
//	for i:= 0; i < val.NumField(); i++ {
//		field := val.Field(i)
//		switch field.Kind() {
//		case reflect.String:
//			fn(field.String())
//		case reflect.Struct:
//			walk(field.Interface(),fn)
//		}
//	}
//}

func walk(x interface{}, fn func(input string))  {
	val := getValue(x)
	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind(){
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _,key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(),fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v,ok =val.Recv(){
			walk(v.Interface(),fn)
		}
	case reflect.Func:
	valFnResult := val.Call(nil)
	for _, res := range valFnResult {
	walk(res.Interface(), fn)
		}
	}

	for i:=0; i< numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value  {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr{
		val = val.Elem()
	}
	return val
}

func main()  {
	expected := "Chris"
	var got []string
	x := struct {
		name string
	}{expected}
	walk(x, func(input string){
		got = append(got, input)
	})
	fmt.Printf("\ngot %q, want %q", got[0], expected)
}