package main

import (
	"fmt"
	"math"
)

func main() {
	x := -123
	fmt.Println(reverse(x))
}
func reverse(x int) (num int) {
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return
}
