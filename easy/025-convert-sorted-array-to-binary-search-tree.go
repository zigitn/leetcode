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
	testData := []int{-10, -3, 0, 5, 9}

	sortedArrayToBST(testData)
	fmt.Println("1")
}

func sortedArrayToBST(nums []int) *TreeNode {
	tmpLen := len(nums) / 2
	tmp := TreeNode{nums[tmpLen], nil, nil}

	// wg:=sync.WaitGroup{}
	//
	// wg.Add(1)
	creatLeftTree(&tmp, nums[:tmpLen], tmpLen-1)
	creatRightTree(&tmp, nums[tmpLen+1:])

	return &tmp
}

func creatLeftTree(node *TreeNode, nums []int, l int) {
	node.Left = &TreeNode{nums[l], nil, nil}
	if l == 0 {
		return
	} else {
		l--
		creatLeftTree(node.Left, nums[:l+1], l)
	}
}
func creatRightTree(node *TreeNode, nums []int) {
	node.Right = &TreeNode{nums[0], nil, nil}

	if tmpLen := len(nums[1:]); tmpLen == 0 {
		return
	} else {
		creatRightTree(node.Right, nums[1:])
	}
}
