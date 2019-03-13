package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node2 := TreeNode{5, nil, nil}
	node3 := TreeNode{4, nil, nil}
	node4 := TreeNode{3, nil, nil}
	node7 := TreeNode{2, &node3, &node2}
	node8 := TreeNode{1, &node7, &node4}

	fmt.Println(isBalanced(&node8))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if intAbs(countFloor(root.Left)-countFloor(root.Right)) > 1 {
		return false
	} else if isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}
func countFloor(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(countFloor(root.Left), countFloor(root.Right))
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
	}
	return b
}
func intAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
