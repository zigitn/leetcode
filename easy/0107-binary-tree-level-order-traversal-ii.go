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
	node2 := TreeNode{5, nil, nil}
	node3 := TreeNode{4, nil, nil}
	node4 := TreeNode{3, nil, nil}

	// node6 := TreeNode{6, nil, nil}
	node7 := TreeNode{2, &node3, &node2}
	node8 := TreeNode{1, &node7, &node4}

	fmt.Println(levelOrderBottom(&node8))

}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res Arr
	levelOrderBottom1(root, &res, 0)
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
		res.List = append([][]int{{root.Val}}, res.List[:]...)
	} else {
		tmp := len(res.List) - 1 - d
		res.List[tmp] = append(res.List[tmp], root.Val)
	}
	levelOrderBottom1(root.Left, res, d+1)
	levelOrderBottom1(root.Right, res, d+1)
}
