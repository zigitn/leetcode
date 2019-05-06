package main

import "bytes"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	reps := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := bytes.NewBufferString("")
	for i := 0; i < 13; i++ {
		for num >= values[i] {
			num -= values[i]
			res.WriteString(reps[i])
		}
	}
	return res.String()
}
