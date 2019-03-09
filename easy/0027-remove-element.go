package main

import "fmt"

func main() {
	nums := []int{}
	val := 2
	fmt.Println(removeElement(nums, val))
}
func removeElement(nums []int, val int) int {
	sum := 0
	index := 0
	for _, v := range nums {
		if v != val {
			nums[index] = v
			sum++
			index++
		}
	}
	return sum
}
