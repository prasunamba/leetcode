package module1

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
type Listhead struct {
	head *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func (l1 *Listhead) Printlink() {
	temp := l1.head
	for temp != nil {
		fmt.Print("->", temp.Val)
		temp = temp.Next
	}
}
func (l1 *Listhead) Push(num int) {
	node := &ListNode{Val: num}
	if l1.head == nil {
		l1.head = node
		return // dont know
	}
	temp := l1.head
	if temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = node
}

/* func main() {
	l1 := Listhead{}
	l1.Push(5)
	l1.Push(6)
	l1.Push(4)
	l1.printlink()
	l2 := Listhead{}
	l2.Push(7)
	l2.Push(0)
	l2.Push(8)
	l2.printlink()
	addTwoNumbers(l1, l2)
} */
