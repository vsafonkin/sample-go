package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	v int
	n *Node
}

func main() {
	var ll LinkedList
	ll.Add(5)
	ll.Add(2)
	ll.Add(17)
	ll.Add(13)
	ll.Add(1)
	ll.String()
	fmt.Println("---")
	ll.Reverse()
	ll.String()
}

func (ll *Node) add(n int) {
	if ll == nil {
		return
	}
	if ll.n == nil {
		ll.n = &Node{
			v: n,
		}
		return
	}
	ll.n.add(n)
}

func (node *Node) Sum() int {
	if node == nil {
		return 0
	}
	return node.v + node.n.Sum()
}

func (ll *LinkedList) Add(n int) {
	if ll == nil {
		return
	}
	if ll.head == nil {
		ll.head = &Node{
			v: n,
		}
		return
	}
	ll.head.add(n)
}

func (ll *LinkedList) String() {
	if ll == nil || ll.head == nil {
		return
	}
	node := ll.head
	for node != nil {
		fmt.Printf("%p %d %p\n", node, node.v, node.n)
		node = node.n
	}
}

func (ll *LinkedList) Reverse() {
	if ll == nil || ll.head == nil {
		return
	}

	var prev *Node
	current := ll.head
	next := ll.head.n
	for next != nil {
		current.n = prev
		prev = current
		current = next
		next = current.n
	}
	current.n = prev
	ll.head = current
}
