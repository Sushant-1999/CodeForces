package main

import "fmt"

func main() {
	var n int64
	var i int64
	var j int64
	var k int64

	fmt.Scan(&n)
	fmt.Scan(&k)
	arr := make([]int64, n*2)

	// First saving odd integers
	val := n % 2
	j = 0
	if val == 0 {
		for i = 1; i < n; i = i + 2 {
			arr[j] = i
			j++
		}

		for i = 2; i <= n; i = i + 2 {
			arr[j] = i
			j++
		}

	} else if val == 1 {
		// Odd n
		for i = 1; i <= n; i = i + 2 {
			arr[j] = i
			j++
		}

		for i = 2; i < n; i = i + 2 {
			arr[j] = i
			j++
		}

	}

	fmt.Print(arr[k-1])

}
