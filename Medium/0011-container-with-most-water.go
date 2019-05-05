package main

func maxArea(height []int) int {
	tmp, res := 0, 0
	for l, r := 0, len(height)-1; l < r; {
		if height[l] < height[r] {
			tmp = height[l] * (r - l)
		} else {
			tmp = height[r] * (r - l)
		}
		if tmp > res {
			res = tmp
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
