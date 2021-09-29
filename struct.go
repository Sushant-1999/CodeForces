package main

import (
	"fmt"
	"reflect"
)

type Gopher struct {
	Name  string
	Color string
	Year  int
}

func main() {
	g := Gopher{Name: "AAA", Color: "BBBB", Year: 2021}

	gtype := reflect.TypeOf(g)

	numFields := gtype.NumField()

	rg := reflect.ValueOf(&g)

	for i := 0; i < numFields; i++ {
		//fmt.Println(rg.Elem().Field(i))
		ptr := rg.Elem().Field(i)
		fmt.Printf("%T", ptr)
		y := ptr.Interface().(string) // y will have type float64.
		fmt.Printf("%T", y)

	}

}
