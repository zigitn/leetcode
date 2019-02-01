package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}


func main() {
	l3 := ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := ListNode{
		Val:  2,
		Next: &l3,
	}
	l1 := ListNode{
		Val:  1,
		Next: &l2,
	}

	l6 := ListNode{
		Val:  3,
		Next: nil,
	}
	l5 := ListNode{
		Val:  2,
		Next: &l6,
	}
	l4 := ListNode{
		Val:  1,
		Next: &l5,
	}
	printListNode(mergeTwoLists(&l1,&l4))
}




func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val{
		res = l2
		res.Next = mergeTwoLists(l1,l2.Next)
	}else{
		res = l1
		res.Next = mergeTwoLists(l1.Next,l2)
	}
	return res
}

func printListNode(l *ListNode) {
	if l.Next != nil {
		fmt.Println(l.Val)
		printListNode(l.Next)
	}
}
