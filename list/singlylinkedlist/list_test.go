package singlylinkedlist
//using container/list test cases

import (
	"fmt"
	"testing"

	//"github.com/seakingj/algorithms"
)

type TestStruct struct {
	Text string
}

func checkListLen(t *testing.T, l *List, len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func checkListPointers(t *testing.T, l *List, nodes []*Node){
	//root := l.root

	if !checkListLen(t, l, len(nodes)) {
		return
	}

	if len(nodes) == 0 {
		if l.root.next != nil || l.tail != l.root {
			t.Errorf("l.root.next, l.tail both should be nil")
		}
		return
	}

	if len(nodes) > 0 {
		var i int
		for node := l.Front(); node != nil; node = node.Next() {
			n := nodes[i]
			if i == 0 {
				if n != l.root.next {
					t.Errorf("node at position 0 should be the head of the list: %p <-> %p ", n, l.root.next)
					return
				}
			}

			//tails
			if i ==  len(nodes) - 1{
				if n != l.tail {
					t.Errorf("node %v and tail %v", *n, *l.tail)
					fmt.Println(n == l.root.next)
					t.Errorf("last node should be the tail of the list: %p <-> %p ", n, l.tail)
					return
				}
			}
			if node.Value != n.Value {
				t.Errorf("Node %v expected %v, got %v", i, node.Value, n.Value)
			}
			i++
		}
		//for i, n := range nodes {
		//	//head
		//	if i == 0 {
		//		if n != l.root.next {
		//			t.Errorf("node at position 0 should be the head of the list: %p <-> %p ", n, l.root.next)
		//			return
		//		}
		//	}
		//
		//	//tails
		//	if i ==  len(nodes) - 1{
		//		if n != l.tail {
		//			t.Errorf("node %v and tail %v", *n, *l.tail)
		//			fmt.Println(n == l.root.next)
		//			t.Errorf("last node should be the tail of the list: %p <-> %p ", n, l.tail)
		//			return
		//		}
		//	}
		//
		//
		//
		//}
	}
}

func TestNewList(t *testing.T) {

	list := (&List{}).Init()
	if list.Front() != nil {
		t.Error("(&List{}).Init().Front() should be nil")
	}
	if list.Back() != nil {
		t.Error("(&List{}).Init().Back() should be nil")
	}
	checkListPointers(t, list, []*Node{})

	//test New list without element

	list = New()
	if list.Front() != nil {
		t.Error("(New().Front() should be nil")
	}
	if list.Back() != nil {
		t.Error("(New().Back() should be nil")
	}
	checkListPointers(t, list, []*Node{})


	nodes := list.nodes()
	if len(nodes) != 0 {
		t.Errorf("list.nodes() len should be 0 on List without element")
		return
	}

	//test New list with single element
	list = New(1)
	if list.Front() != list.root.next {
		t.Error("New(1).Front() should equal to list.root.next")
	}
	if list.Back() != list.root.next {
		t.Error("(New(1).Back() should equal to list.root.next")
	}
	nodes = list.nodes()
	if len(nodes) != 1 {
		t.Errorf("expected 1, got %v", len(nodes))
		return
	}

	//t.Logf("Test New List without Element %v %v", list.Len(), list.nodes())
	checkListPointers(t, list, list.nodes())

	//test New list with multiple elements
	list = New(1, "a", true, 1.1, TestStruct{"test struct"})
	if list.Front() != list.root.next {
		t.Error("List with multiple elements Front() should equal to list.root.next")
	}
	if list.Back() != list.tail {
		t.Error("List with multiple elements Back() should equal to list.tail")
	}

	nodes = list.nodes()
	if len(nodes) != 5 {
		t.Errorf("expected 5, got %v", len(nodes))
		return
	}
	checkListPointers(t, list, list.nodes())
}

func TestList(t *testing.T) {
	//1. Test List Init

	//list := New()
	//checkListPointers(t, list, []*Node{})
	//
	////list := New(1, "a")
	//
	//if actualValue := list.Len(); actualValue != 2 {
	//	t.Errorf("Got %v expected %v", actualValue, 2)
	//}

	//test single node
	list := New()
	n := list.PushFront("a")
	checkListPointers(t, list, []*Node{n})

	list.MoveToFront(n)
	checkListPointers(t, list, []*Node{n})
	list.MoveToBack(n)
	checkListPointers(t, list, []*Node{n})

}
