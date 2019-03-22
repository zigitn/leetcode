package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		tmp := numbers[i] + numbers[j]
		if tmp == target {
			return []int{i + 1, j + 1}
		}
		if tmp > target {
			j--
			continue
		} else {
			i++
			continue
		}
	}
	return []int{0, 0}
}
