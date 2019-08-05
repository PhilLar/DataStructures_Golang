package list

import (
	"fmt"
)

type Node struct {
	Data interface {}
	next *Node
	prev *Node
}

type List struct {
	head *Node
	length int
}

func (l *List) Show() {
	if l.length != 0 {
		tmp := l.head
		// List.Append(n) work as stack append
		// so first find the last node
		for i:=1; i<l.length; i++ {
			tmp = l.head.next
		}
		// second print nodes from the end
		for i:=0; i<l.length; i++ {
			fmt.Println(tmp.Data)
			tmp = tmp.prev
		}
	}
}

func (l *List) Append(data ... interface{}) {
	for _, d:= range data {
	node := &Node{Data:d}
	node.next = l.head
	if l.head != nil {
		l.head.prev = node
	}
	l.head = node
	l.length++
}
}

// func (l *List) Remove(node *Node) {

// }