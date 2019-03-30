package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("AA"))
}

func titleToNumber(s string) int {
	var res int
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		res += int(math.Pow(26, float64(j))) * int(int(s[i])-64)
	}
	return res
}
