package main
import "fmt"
var (
	factVal uint64 = 1
	i       int    = 1
)
func Factorial(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
	}
	return factVal /* return from function*/
}
func main() {
	var num int
	fmt.Scan(&num)
	res := Factorial(num)
	fmt.Printf("Factorial of %v is : %v", num, res)
}
