package main

import (
	"fmt"
)

func main() {
	s0 := "IXC"
	s1 := "D"
	//s2 := "IVXLCDM"
	fmt.Println(romanToInt(s0))
	fmt.Println(romanToInt(s1))
	//fmt.Println(romanToInt(s2))
}

func romanToInt(s string) int {
	map0 := map[uint8]int{73: 1, 86: 5, 88: 10, 76: 50, 67: 100, 68: 500, 77: 1000}
	sum := 0
	for i:=0; i<len(s) -1; i++{
		if map0[s[i]] < map0[s[i+1]]{
			sum -= map0[s[i]]
		}else{
			sum += map0[s[i]]
		}
	}
	return sum + map0[s[len(s)-1]]
}
