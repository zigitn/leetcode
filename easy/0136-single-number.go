package main

import "fmt"

func main() {
	a := []int{1, 1, 3, 3}

	fmt.Println(singleNumber(a))
}

func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}
