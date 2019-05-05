package main

import "bytes"

//Runtime: 8 ms, faster than 94.49% of Go online submissions for ZigZag Conversion.
//Memory Usage: 4.2 MB, less than 85.45% of Go online submissions for ZigZag Conversion.
func convert(s string, numRows int) string {
	MOD := 2*numRows - 2
	if len(s) <= 1 || numRows == 1 {
		return s
	}

	res := bytes.NewBufferString("")

	for row := 0; row < numRows; row++ {
		start1 := row
		start2 := (MOD - row) % MOD
		for start1 < len(s) {
			res.WriteByte(s[start1])
			if start2 != start1 && start2 < len(s) {
				res.WriteByte(s[start2])
			}
			start1 += MOD
			start2 += MOD
		}
	}
	return res.String()
}
