package main

import "fmt"

func main() {
	var strs = []string{
		"flower", "flow", "flight",
	}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	minLen := len(strs[0])
	index := 0
	for i := 0; i < l; i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
			index = i
		}
	}
	for i := 0; i < minLen; i++ {
		for k := 0; k < len(strs); k++ {
			if strs[index][i] != strs[k][i] {
				return strs[index][:i]
			}
		}
	}
	return strs[index]
}
