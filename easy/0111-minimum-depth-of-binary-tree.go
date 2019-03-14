package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

//4ms 5.4 MB
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	i := 0
	mind(root, &i)
	return i
}

func mind(root *TreeNode, i *int) {
	*i++
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left != nil && root.Right != nil {
		j := *i
		mind(root.Left, i)
		mind(root.Right, &j)
		if *i > j {
			*i = j
		}
		return
	}
	if root.Right != nil {
		mind(root.Right, i)
		return
	}
	if root.Left != nil {
		mind(root.Left, i)
		return
	}
}
