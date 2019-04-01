package main

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 2)
}

// Runtime: 52 ms, faster than 100.00% of Go online submissions for Rotate Array.
// Memory Usage: 7.6 MB, less than 52.94% of Go online submissions for Rotate Array.
func rotate(nums []int, k int) {
	if k != 0 {
		nLen := len(nums)
		if k > nLen {
			k %= nLen
		}
		for i := 0; i < nLen/2; i++ {
			nums[i], nums[nLen-i-1] = nums[nLen-i-1], nums[i]
		}
		for i := 0; i < k/2; i++ {
			nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
		}
		for i := k; i < (nLen+k)/2; i++ {
			nums[i], nums[nLen+k-1-i] = nums[nLen+k-1-i], nums[i]
		}
	}
}
