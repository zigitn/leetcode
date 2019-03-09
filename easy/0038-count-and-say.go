package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=", countAndSay(16), "=")
}

func countAndSay(n int) string {
	s := "1"
	for i := 0; i < n-1; i++ {
		s = say(s)
	}
	return s
}

func say(s string) string {
	temp := s[0]
	thisCount := 0
	var buf bytes.Buffer
	for k := range s {
		if temp == s[k] {
			thisCount++
			continue
		}
		buf.WriteString(strconv.Itoa(thisCount))
		buf.WriteString(string(s[k-1]))
		temp = s[k]
		thisCount = 1
	}
	buf.WriteString(strconv.Itoa(thisCount))
	buf.WriteString(string(temp))
	return buf.String()
}
