package module1

import "fmt"

type node struct {
	num int
	ptr *node
}
type Linkedlist struct {
	head *node
}

func Linkedlistprint() {
	ll := Linkedlist{}
	ll.Insert(5)
	ll.Insert(6)
	ll.Insert(4)

	ll.printll()

}
func (ll *Linkedlist) printll() {
	temp := ll.head
	for temp != nil {
		fmt.Print("->", temp.num)
		temp = temp.ptr
	}
}
func (ll *Linkedlist) Insert(data int) {
	newnode := &node{num: data}
	if ll.head == nil {
		ll.head = newnode
		return
	}
	temp := ll.head
	for temp.ptr != nil {
		temp = temp.ptr
	}
	temp.ptr = newnode
}
