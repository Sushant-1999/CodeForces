package main

import (
	"fmt"
	"reflect"
)

type x struct {
	Foo string
	Bar int
}

func main() {

	a := x{"Krishna", 100}

	v := reflect.ValueOf(a)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	fmt.Printf("%T", values)
	fmt.Println(values)
	//var x interface{} = "abc"
	str := fmt.Sprintf("%v", values)
	fmt.Printf("%T", str)
	fmt.Println(str)

}
