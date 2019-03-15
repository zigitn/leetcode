package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	num := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 || j == i {
				num[j] = 1

			} else {
				num[j] = num[j] + num[j-1]
			}
		}
	}
	return num
}
