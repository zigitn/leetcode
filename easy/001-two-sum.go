package main

import "fmt"

func main() {
	nums := []int{-3, 4, 3, 90, 4}
	target := 0
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	numMap0 := make(map[int]int)
	numMap1 := make(map[int]int)
	for k, v := range nums {
			numMap0[k] = v
			numMap1[v] = k
	}
	for k0, v0 := range numMap0 {
		//fmt.Println(target-v0, numMap1[target-v0])
		if k1, ok := numMap1[target-v0]; ok && k0 != k1 {
			return []int{k0, k1}
		}
	}
	fmt.Println("err")
	return []int{404, 404}
}
