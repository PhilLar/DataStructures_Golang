package list

import (
	"fmt"
)

type Node struct {
	Data interface {}
	Next *Node
	Prev *Node
}

type List struct {
	head *Node
	length int
}

func(l *List) Head() *Node {
	return l.head
}

func(l *List) SetHead(node *Node) {
	l.head = node
}

func(l *List) Length() int {
	return l.length
}

func(l *List) SetLength(length int) {
	l.length = length
}

func (l *List) Show() {
	if l.length != 0 {
		tmp := l.head
		// List.Append(n) work as stack append
		// so first find the last node
		for i:=1; i<l.length; i++ {
			tmp = tmp.Next
		}
		// second print nodes from the end
		for i:=0; i<l.length; i++ {
			fmt.Println(tmp.Data)
			tmp = tmp.Prev
		}
	}
}

func (l *List) Append(data ... interface{}) {
	for _, d:= range data {
	node := &Node{Data:d}
	node.Next = l.head
	if l.head != nil {
		l.head.Prev = node
	}
	l.head = node
	l.length++
}
}

func (l *List) Remove(index int) bool{
	if l.length != 0 {
		tmp := l.head
		for i:=0; i<l.length; i++ {
			if index==i {
				tmp.Prev.Next = tmp.Next
				l.length--
				return true
			}
			tmp = tmp.Next
		}
	}
	return false
}