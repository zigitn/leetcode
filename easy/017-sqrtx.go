package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(2147395599))
}

func mySqrt(x int) int {
	n := float64(x)
	n2 := float64(x)
	for (n*n - float64(x)) > 0.0001 {
		n = (n + n2/n) / 2
	}
	fmt.Println(n)
	return int(n)
}
