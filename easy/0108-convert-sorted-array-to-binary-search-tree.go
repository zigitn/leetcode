package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := sortedArrayToBST([]int{-11, -10, -3, 0, 5, 9, 10})
	fmt.Println(a)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return buildTree(nums)
}
func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  buildTree(nums[:len(nums)/2]),
		Right: buildTree(nums[len(nums)/2+1:]),
	}
}
