package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := 1, 1
	countLength(headA, &a)
	countLength(headB, &b)
	if a != b {
		if a > b {
			for i := 0; i < a-b; i++ {
				headA = headA.Next
			}
		} else {
			for i := 0; i < b-a; i++ {
				headB = headB.Next
			}
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func countLength(root *ListNode, c *int) {
	tmpRoot := root
	for tmpRoot.Next != nil {
		tmpRoot = tmpRoot.Next
		*c++
	}
}
