package main

import (
	"fmt"
)

func main() {
	fmt.Println(trailingZeroes(1000))
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Factorial Trailing Zeroes.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Factorial Trailing Zeroes.
func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		res += n / 5
		n = n / 5
	}
	return res
}

// Runtime: 4 ms, faster than 30.21% of Go online submissions for Factorial Trailing Zeroes.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Factorial Trailing Zeroes.
// func trailingZeroes(n int) int {
// 	res := 0
// 	for i := 5; n/i != 0; i *= 5 {
// 		res += n / i
// 	}
// 	return res
// }
