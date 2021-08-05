package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Value uint8
	Prev  *Node
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Add(n Node) {
	if l.Head == nil {
		l.Head, l.Tail = &n, &n
	} else {
		l.Tail.Next = &n
		n.Prev = l.Tail
		l.Tail = &n
	}
}

// Create printable format for values of LinkedList
func (l LinkedList) String() string {
	Node := l.Head
	printable := "Values of LinkedList: ["
	for Node.Next != nil {
		printable += fmt.Sprintf("%+v,", Node.Value)
		Node = Node.Next
	}
	// The following line helps remove last comma ([1,2,3,])
	printable += fmt.Sprintf("%+v", Node.Value) + "]"
	return printable
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	l1 := LinkedList{}
	for i := 0; i < 7; i++ {
		l1.Add(Node{Value: uint8(rand.Intn(10))})
	}

	l2 := LinkedList{}
	for i := 0; i < 5; i++ {
		l2.Add(Node{Value: uint8(rand.Intn(10))})
	}

	var newLinkedList = AddLinkedLists(&l1, &l2)

	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println("New LinkedList: ")
	fmt.Println(newLinkedList)
}

func addNodes(Node1 *Node, Node2 *Node, newNode *Node, extra uint8) {
	if Node1 == nil {
		Node1 = &Node{}
	}
	if Node2 == nil {
		Node2 = &Node{}
	}

	val := Node1.Value + Node2.Value + extra
	newNode.Value = val % 10
	extra = val / 10

	if Node1.Next != nil || Node2.Next != nil || extra > 0 {
		newNode.Next = &Node{}
		addNodes(Node1.Next, Node2.Next, newNode.Next, extra)
	}
}

func AddLinkedLists(ll1 *LinkedList, ll2 *LinkedList) (sum *LinkedList) {
	sum = &LinkedList{}
	sum.Head = &Node{}
	addNodes(ll1.Head, ll2.Head, sum.Head, 0)
	return sum
}
