package main

import "fmt"

func main() {
	fmt.Println(generate(1))
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	res := make([][]int, numRows)

	res[0] = []int{1}
	res[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = row
	}
	return res
}

/*
1
1 1
1 2 1
1 3 3 1
1 4 6 4 1

*/
