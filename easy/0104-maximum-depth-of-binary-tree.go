package main

import (
	"fmt"
	"sync"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func main() {
	node1 := TreeNode{1, nil, nil}
	node2 := TreeNode{1, &node1, nil}
	node3 := TreeNode{1, &node2, nil}
	node4 := TreeNode{1, nil, &node3}

	node6 := TreeNode{1, nil, nil}
	node7 := TreeNode{1, &node6, nil}
	node8 := TreeNode{1, &node7, &node4}

	fmt.Println("res=", maxDepth(&node8))
}

func maxDepth(root *TreeNode) int {
	var wg sync.WaitGroup
	return maxDepthGo(root, wg)
}

func maxDepthGo(root *TreeNode, wg sync.WaitGroup) int {
	if root == nil {
		return 0
	}
	countLeft := 1
	countRight := 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		if root.Left != nil {
			countLeft += maxDepth(root.Left)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if root.Right != nil {
			countRight += maxDepth(root.Right)
		}
	}()
	wg.Wait()
	if countLeft > countRight {
		return countLeft
	} else {
		return countRight
	}
}
