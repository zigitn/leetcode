package main

import "fmt"

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
	return false
}
