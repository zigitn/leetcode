package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit2([]int{1, 2, 1, 4, 5}))
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
			continue
		}

	}
	return res
}
