package main

import "fmt"

func checkMagic(num int64) {
	flag := 0
	for num > 0 {
		if num%1000 == 144 {
			num = num / 1000
		} else if num%100 == 14 {
			num = num / 100
		} else if num%10 == 1 {
			num = num / 10
		} else {
			fmt.Println("NO")
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("YES")
	}

}

func main() {
	var num int64
	fmt.Scan(&num)

	checkMagic(num)

}
