package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for _, v := range nums {
		if v != nums[i] {
			i++
			nums[i] = v
		}

	}
	return i + 1
}
