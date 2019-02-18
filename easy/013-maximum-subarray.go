package main

import (
	"fmt"
)

func main() {
	nums := []int{-3, -2, 1, 2, 2, 0, 1, 0}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, thisSum := nums[0], 0
	for _, v := range nums {
		if thisSum < 0 {
			thisSum = 0
		}
		thisSum += v
		if maxSum < thisSum {
			maxSum = thisSum
		}
	}
	return maxSum
}
