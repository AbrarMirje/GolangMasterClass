// package main provides math solution
package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(6, 8))
	fmt.Println(sum(3, 8, 9, 2))
}

// sum add unlimited number of values of type int
func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
