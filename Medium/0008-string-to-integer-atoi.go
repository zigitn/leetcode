package main

import (
	"math"
)

//Runtime: 4 ms, faster than 100.00% of Go online submissions for String to Integer (atoi).
//Memory Usage: 2.3 MB, less than 85.71% of Go online submissions for String to Integer (atoi).
func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	res := 0
	begin := 0
	end := len(str) - 1

	//从前往后判断
	signSet := false
	sign := 1
	for i := 0; i < len(str) && (str[i] < '0' || str[i] > '9'); i++ {
		if signSet || (str[i] < '0' && !(str[i] == '+' || str[i] == ' ' || str[i] == '-')) || str[i] > 57 {
			return 0
		} else {
			if str[i] == ' ' {
				begin = i + 1
				continue
			}
			if str[i] == '-' || str[i] == '+' {
				begin = i + 1
				sign = 44 - int(str[i])
				signSet = true
				continue
			}
		}
	}

	//从后往前判断
	for i := len(str) - 1; i >= begin; i-- {
		if str[i] < '0' || str[i] > '9' {
			end = i - 1
		}
	}

	for i := begin; i <= end; i++ {
		res = res*10 + int(str[i]-'0')
		if res > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
	}

	return res * sign
}
