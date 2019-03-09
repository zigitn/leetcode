package main

import (
	"fmt"
)

func main() {
	haystack := ""
	needle := ""
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	for i := 0; i <= l1-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}
