package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 1234321
	y := -1234321
	fmt.Println(isPalindrome1(x))
	fmt.Println(isPalindrome2(y))
}

func isPalindrome1(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	} else {
		num := 0
		for x > num {
			num = num*10 + x%10
			x = x / 10
		}
		return x == num || num/10 == x

	}
}

func isPalindrome2(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
