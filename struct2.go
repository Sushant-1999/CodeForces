package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func createKeyValuePairs(m map[int]reflect.Value) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s %s", key, value)
		
	}
	return b.String()
}
func main() {
	a := []int{2, 5}

	var b reflect.Value = reflect.ValueOf(&a)
	m := map[int]reflect.Value{
		100: b,
	}
	res := createKeyValuePairs(m)

	fmt.Println(res)

}
