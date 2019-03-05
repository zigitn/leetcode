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
	// node1 := TreeNode{1, nil, nil}
	node2 := TreeNode{7, nil, nil}
	node3 := TreeNode{15, nil, nil}
	node4 := TreeNode{20, &node2, &node3}

	// node6 := TreeNode{6, nil, nil}
	node7 := TreeNode{9, nil, nil}
	node8 := TreeNode{3, &node7, &node4}

	fmt.Println(levelOrderBottom(&node8))

}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res Arr
	depth := 0

	levelOrderBottom1(root, &res, depth+1)

	return res.List
}

type Arr struct {
	List [][]int
}

func levelOrderBottom1(root *TreeNode, res *Arr, d int) {
	if root == nil {
		return
	}

	if (len(res.List) - 1) < d {
		res.List = append(res.List, []int{root.Val})
	} else {
		res.List[d] = append(res.List[d], root.Val)
	}

	levelOrderBottom1(root.Left, res, d+1)
	levelOrderBottom1(root.Right, res, d+1)

	fmt.Println(d, len(res.List))

}
