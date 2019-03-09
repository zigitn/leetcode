package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	a := strings.Fields(s)
	if len(a) == 0 {
		return 0
	}
	return len(a[len(a)-1])
}
