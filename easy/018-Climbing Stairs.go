package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	one, two := 3, 2
	for i := 4; i <= n; i++ {
		two, one = one, two+one
	}
	return one
}
