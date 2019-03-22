package main

func main() {

}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
