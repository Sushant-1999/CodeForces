package main

import "fmt"

func main() {
	var T, K, temp int
	temp = 0
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&K)
		for K != 0 {
			temp = temp + 1
			if temp%3 != 0 && temp%10 != 3 {
				// Not Discarded
				K = K - 1
			}
		}

		fmt.Println(temp)
		temp = 0
	}

}
