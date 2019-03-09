package main

import (
	"fmt"
)

func main() {
	nums := []int{9, 9, 9}
	fmt.Println(plusOne(nums))
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}
