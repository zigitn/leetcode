package main

import (
	"fmt"
)

func main() {
	nums := []int{1,8}
	target := 7
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	index := 0
	for k, v := range nums {
		if target == v {
			return k
		} else if target > v {
			index = k+1
			continue
		}
		return index
	}
	return index
}
