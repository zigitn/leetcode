package main

import "fmt"

func main() {
	a := Constructor()
	a.Push(2)
	a.Push(1)
	a.Push(3)
	a.Push(4)
	fmt.Println(a.GetMin())
	a.Pop()
	fmt.Println(a.GetMin())
	a.Pop()
	fmt.Println(a.GetMin())
	a.Pop()
	fmt.Println(a.GetMin())
	a.Pop()
	fmt.Println(a.GetMin())
}

type MinStack struct {
	min  int
	Val  int
	Next *MinStack
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	pm := *m
	m.Val = x
	m.Next = &pm
	m.min = x
	if m.Next.Next != nil && x > pm.min {
		m.min = pm.min
	}

}

func (m *MinStack) Pop() {
	if m.Next != nil {
		*m = *m.Next
	} else {
		*m = MinStack{}
	}
}

func (m *MinStack) Top() int {
	if m == nil {
		return 0
	}
	return m.Val
}

func (m *MinStack) GetMin() int {
	return m.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
