package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "0"
	b := "0"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	aa := strings.Split(a, "")
	bb := strings.Split(b, "")
	la := len(a)
	lb := len(b)
	sl := 0
	var extra []string
	diff := 0
	if la > lb {
		diff = la - lb
		sl = lb
		extra = aa[:diff]
		aa = aa[diff:]
	} else {
		diff = lb - la
		sl = la
		extra = bb[:diff]
		bb = bb[diff:]
	}

	carry := false
	for i := sl - 1; i >= 0; i-- {
		if aa[i] == bb[i] {
			if carry {
				aa[i] = "1"
				carry = bb[i] == "1"
			} else {
				aa[i] = "0"
				carry = bb[i] == "1"
			}
		} else {
			if carry {
				aa[i] = "0"
				carry = true
			} else {
				aa[i] = "1"
				carry = false
			}
		}
	}
	for i := diff - 1; i >= 0; i-- {
		if carry {
			if extra[i] == "1" {
				extra[i] = "0"
			} else {
				extra[i] = "1"
				carry = false
				break
			}
		}
	}
	if carry {
		extra = append([]string{"1"}, extra...)
	}

	aa = append(extra, aa...)

	return strings.Trim(strings.Replace(fmt.Sprint(aa), " ", "", -1), "[]")
}
