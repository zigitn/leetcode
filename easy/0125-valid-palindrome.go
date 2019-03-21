package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	tmpLen := len(s)
	if tmpLen < 2 {
		return true
	}
	for i, j := 0, tmpLen-1; i < tmpLen/2; {
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			if !(unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i]))) {
				if !(unicode.IsLetter(rune(s[j])) || unicode.IsNumber(rune(s[j]))) {
					j--
				}
				i++
				continue
			} else {
				if !(unicode.IsLetter(rune(s[j])) || unicode.IsNumber(rune(s[j]))) {
					j--
					continue
				}
				return false
			}
		}
		i++
		j--
	}
	return true
}
