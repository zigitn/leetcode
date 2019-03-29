package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(28))
}

func convertToTitle(n int) string {
	if n <= 0 {
		return ""
	}
	s := ""
	for n > 0 {
		n--
		s = string(rune(n%26+65)) + s
		n = n / 26
	}
	return s
}
