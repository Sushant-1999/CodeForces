package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	a := []int{2, 5}
	var b reflect.Value = reflect.ValueOf(&a)
	m := map[int]reflect.Value{
		100: b,
	}
	/*
		// Convert each key/value pair in m to a string
		res := fmt.Sprint(m)
		// Do whatever you want to do with the string;
		// in this example I just print out each of them.
		fmt.Println(res)
	*/

	// Marshal the map into a JSON string.
	mJson, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonStr := string(mJson)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)

	table := "<table border=1>" + jsonStr + "</table>"
	fmt.Println(table)

}
