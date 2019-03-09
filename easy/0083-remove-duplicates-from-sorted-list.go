package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func main() {
	l5 := ListNode{
		Val:  1,
		Next: nil,
	}
	l4 := ListNode{
		Val:  1,
		Next: &l5,
	}
	l3 := ListNode{
		Val:  1,
		Next: &l4,
	}
	l2 := ListNode{
		Val:  1,
		Next: &l3,
	}
	l1 := ListNode{
		Val:  1,
		Next: &l2,
	}

	// l6 := ListNode{
	// 	Val:  2,
	// 	Next: nil,
	// }

	deleteDuplicates(&l1)
	deleteDuplicates(&l4)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for head.Next != nil && head.Val == head.Next.Val {
		head.Next = head.Next.Next
	}
	deleteDuplicates(head.Next)
	return head
}
