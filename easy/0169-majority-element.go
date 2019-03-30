package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

// Runtime: 16 ms, faster than 100.00% of Go online submissions for Majority Element.
// Memory Usage: 5.9 MB, less than 56.82% of Go online submissions for Majority Element.
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}
	var res int
	var tmp int
	for k, v := range numMap {
		if v > tmp {
			tmp = v
			res = k
		}
	}
	return res
}

// Runtime: 388 ms, faster than 5.07% of Go online submissions for Majority Element.
// Memory Usage: 5.9 MB, less than 22.73% of Go online submissions for Majority Element.
// func majorityElement(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	for i, j := 0, 1; i < len(nums) && j < len(nums); {
// 		if nums[i] == nums[j] {
// 			j++
// 			continue
// 		}
// 		nums = append(nums[1:j], nums[j+1:]...)
// 		j = 1
// 	}
// 	return nums[0]
// }
