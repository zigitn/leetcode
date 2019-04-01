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
	for k, v := range s {
		res += int(math.Pow(26, float64(len(s)-1-k))) * int(v-64)
	}
	return res
}
