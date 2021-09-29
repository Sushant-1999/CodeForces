package main

import "fmt"

func KBeautiful(num int) bool {

}

func main() {
	var T, n, k, x, flag int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&n)
		fmt.Scan(&k)
		flag = 1
		for x = n; flag != 0; x++ {
			res := KBeautiful(x)
			if res {
				flag = 0
			} else {
				flag = 1
			}
		}

		fmt.Print(x)

	}

}
