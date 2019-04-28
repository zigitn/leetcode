package main

import (
	"fmt"
)

func main() {
	a := lengthOfLongestSubstring("asdw")
	fmt.Println(a)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	max, cur := 1, 0
	for i := 1; i < len(s); i++ {
		j := cur
		for j < i {
			if s[j] == s[i] {
				break
			}
			j++
		}
		if i == j {
			if max < j-cur+1 {
				max = j - cur + 1
			}
		} else {
			cur = j + 1
		}
	}
	return max
}
