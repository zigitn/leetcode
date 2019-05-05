package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("dw"))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	begin, maxLen := 0, 1

	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
		}
		i = e + 1
		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
		}
		newLen := e + 1 - b
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}
	return s[begin : begin+maxLen]
}

//Runtime: 4 ms, faster than 92.35% of Go online submissions for Longest Palindromic Substring.
//Memory Usage: 2.2 MB, less than 98.26% of Go online submissions for Longest Palindromic Substring.

//func longestPalindrome(s string) string {
//	if len(s) < 2 {
//		return s
//	}
//	start := 0
//	end := 0
//	for i := 0; i < len(s); i++ {
//		len1 := expandAroundCenter(s, i, i)
//		len2 := expandAroundCenter(s, i, i+1)
//		if len2 > len1 {
//			len1 = len2
//		}
//
//		if len1 > end-start {
//			start = i - (len1-1)/2
//			end = i + len1/2
//		}
//	}
//	return s[start : end+1]
//}
//
//func expandAroundCenter(s string, left int, right int) int {
//	for left >= 0 && right < len(s) && s[left] == s[right] {
//		left--
//		right++
//	}
//	return right - left - 1
//}
