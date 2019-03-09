package main

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for k, v := range nums2 {
			nums1[k] = v
		}
		return
	}
	for i, j, k := m-1, n-1, m+n-1; i >= 0 && j >= 0 && k >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			if i == 0 {
				for l := j; l >= 0; l-- {
					nums1[l] = nums2[l]
				}
				return
			}
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
