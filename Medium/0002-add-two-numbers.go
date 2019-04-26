package medium

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	p := l1
	q := l2
	carry := 0
	for {
		p.Val = p.Val + q.Val + carry
		if p.Val > 9 {
			p.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
		if p.Next == nil {
			if q.Next == nil {
				if carry == 1 {
					p.Next = &ListNode{1, nil}
					carry = 0
				}
				break
			}
			p.Next = &ListNode{0, nil}
		}
		if q.Next == nil {
			q.Next = &ListNode{0, nil}

		}
		p = p.Next
		q = q.Next
	}
	return l1
}
