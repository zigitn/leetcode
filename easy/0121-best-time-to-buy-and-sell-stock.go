package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	minP, maxP := math.MaxInt32, 0
	for i := 0; i < len(prices); i++ {
		if minP > prices[i] {
			minP = prices[i]
		}
		if tmp := prices[i] - minP; maxP < tmp {
			maxP = tmp
		}
	}
	return maxP
}
