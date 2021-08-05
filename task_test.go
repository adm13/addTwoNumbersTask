package main

import (
	"fmt"
	"testing"
)

func TestAddLinkedLists(t *testing.T) {

	l1 := LinkedList{}
	l1.Add(Node{Value: 2})
	l1.Add(Node{Value: 4})
	l1.Add(Node{Value: 3})
	l2 := LinkedList{}
	l2.Add(Node{Value: 5})
	l2.Add(Node{Value: 6})
	l2.Add(Node{Value: 4})

	l3 := LinkedList{}
	l3.Add(Node{Value: 0})
	l4 := LinkedList{}
	l4.Add(Node{Value: 0})

	l5 := LinkedList{}
	l5.Add(Node{Value: 9})
	l5.Add(Node{Value: 9})
	l5.Add(Node{Value: 9})
	l5.Add(Node{Value: 9})
	l5.Add(Node{Value: 9})
	l5.Add(Node{Value: 9})
	l5.Add(Node{Value: 9})
	l6 := LinkedList{}
	l6.Add(Node{Value: 9})
	l6.Add(Node{Value: 9})
	l6.Add(Node{Value: 9})
	l6.Add(Node{Value: 9})

	var actual = AddLinkedLists(&l1, &l2)
	if fmt.Sprint(actual) != "Values of LinkedList: [7,0,8]" {
		t.Errorf("Expected value doesn't match to actual value! l1: [342], l2: [465]")
	}

	actual = AddLinkedLists(&l3, &l4)
	if fmt.Sprint(actual) != "Values of LinkedList: [0]" {
		t.Errorf("Expected value doesn't match to actual value! l3: [0], l4: [0]")
	}

	actual = AddLinkedLists(&l5, &l6)
	if fmt.Sprint(actual) != "Values of LinkedList: [8,9,9,9,0,0,0,1]" {
		t.Errorf("Expected value doesn't match to actual value! l5: [9,9,9,9,9,9,9], l6: [9,9,9,9]")
	}
}
