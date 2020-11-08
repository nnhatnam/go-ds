package sllist

import (
	"fmt"
	"testing"
)



//using container/list test cases

type TestStruct struct {
	Text string
}

func printValues(l *List) {
	values := make([]interface{}, l.length)

	l.Traverse(func(n *Element, i int){
		values[i] =  n.Value
	})

	fmt.Println(values)
}

func manualCount(l *List) int {
	count := 0
	l.Traverse(func(n *Element, i int){
		count += 1
	})
	return count
}

func getNodeArr(l *List) []*Element {
	nodes := make([]*Element, l.length)
	l.Traverse(func(n *Element, i int){
		nodes[i] = n
	})
	return nodes
}

func checkListLen(t *testing.T, l *List, len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func checkListPointers(t *testing.T, l *List, nodes []*Element){
	//root := l.root

	if !checkListLen(t, l, len(nodes)) {
		return
	}

	//zero length
	if len(nodes) == 0 {
		if l.head != nil || l.tail != l.head {
			t.Errorf("l.root.head, l.tail both should be nil")
		}
		return
	}

	//lenght > 0
	l.Traverse(func(node *Element, i int) {
		if i == 0 {
			if node != l.head {
				t.Errorf("node at position 0 should be the head of the list: %p <-> %p ", node, l.head)
				return
			}

			if i == l.Len() - 1 {
				if node != l.tail {
					t.Errorf("last node should be the tail of the list: %p <-> %p ", node, l.tail)
				}
			}
		}
		if node != nodes[i] {
			t.Errorf("node are not repectively equal: %p <-> %p ", node, nodes[i])
		}
	})

}

func TestNewList(t *testing.T) {

	list := &List{}
	if list.First() != nil {
		t.Error("(&List{}).Front() should be nil")
	}
	if list.Last() != nil {
		t.Error("(&List{}).Back() should be nil")
	}
	checkListPointers(t, list, []*Element{})

	//test New list without element

	list = New()
	if list.First() != nil {
		t.Error("(New().Front() should be nil")
	}
	if list.Last() != nil {
		t.Error("(New().Back() should be nil")
	}
	checkListPointers(t, list, []*Element{})


	//test New list with single element
	list = New(1)
	if list.First() != list.head {
		t.Error("New(1).Front() should equal to list.root.head")
	}
	if list.Last() != list.tail {
		t.Error("(New(1).Back() should equal to list.root.tail")
	}

	if list.length != 1 || manualCount(list) != 1 {
		t.Errorf("expected 1, got %v", list.length)
		return
	}

	//t.Logf("Test New List without Element %v %v", list.Len(), list.nodes())
	checkListPointers(t, list, getNodeArr(list))

	//test New list with multiple elements
	list = New(1, "a", true, 1.1, TestStruct{"test struct"})
	if list.First() != list.head.next {
		t.Error("List with multiple elements First() should equal to list.root.head")
	}
	if list.Last() != list.tail {
		t.Error("List with multiple elements Last() should equal to list.root.tail")
	}

	if list.length != 5 || manualCount(list) != 5 {
		t.Errorf("expected 5, got %v", list.length)
		return
	}
	checkListPointers(t, list, getNodeArr(list))
}

func TestList(t *testing.T) {
	l := New()
	checkListPointers(t, l, []*Element{})

	// Single element list
	e := l.PushFront("a")
	checkListPointers(t, l, []*Element{e})

	l.MoveToFront(e)
	checkListPointers(t, l, []*Element{e})
	l.MoveToBack(e)
	checkListPointers(t, l, []*Element{e})
	l.Remove(e)
	checkListPointers(t, l, []*Element{})

	// Bigger list
	e2 := l.PushFront(2)
	e1 := l.PushFront(1)
	e3 := l.Push(3)
	e4 := l.Push("banana")
	checkListPointers(t, l, []*Element{e1, e2, e3, e4})

	l.Remove(e2)
	checkListPointers(t, l, []*Element{e1, e3, e4})

	l.MoveToFront(e3) // move from middle
	checkListPointers(t, l, []*Element{e3, e1, e4})

	l.MoveToFront(e1)
	l.MoveToBack(e3) // move from middle
	checkListPointers(t, l, []*Element{e1, e4, e3})

	l.MoveToFront(e3) // move from back
	checkListPointers(t, l, []*Element{e3, e1, e4})
	l.MoveToFront(e3) // should be no-op
	checkListPointers(t, l, []*Element{e3, e1, e4})
	l.MoveToBack(e3) // move from front
	checkListPointers(t, l, []*Element{e1, e4, e3})
	l.MoveToBack(e3) // should be no-op
	checkListPointers(t, l, []*Element{e1, e4, e3})

	e2 = l.InsertBefore(2, e1) // insert before front
	checkListPointers(t, l, []*Element{e2, e1, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(2, e4) // insert before middle
	checkListPointers(t, l, []*Element{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(2, e3) // insert before back
	checkListPointers(t, l, []*Element{e1, e4, e2, e3})
	l.Remove(e2)

	e2 = l.InsertAfter(2, e1) // insert after front
	checkListPointers(t, l, []*Element{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(2, e4) // insert after middle
	checkListPointers(t, l, []*Element{e1, e4, e2, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(2, e3) // insert after back
	checkListPointers(t, l, []*Element{e1, e4, e3, e2})
	l.Remove(e2)

	// Check standard iteration.
	sum := 0
	for e := l.First(); e != nil; e = e.Next() {
		if i, ok := e.Value.(int); ok {
			sum += i
		}
	}
	if sum != 4 {
		t.Errorf("sum over l = %d, want 4", sum)
	}

	// Clear all elements by iterating
	var next *Element
	for e := l.First(); e != nil; e = next {
		next = e.Next()
		l.Remove(e)
	}
	checkListPointers(t, l, []*Element{})
}

func checkList(t *testing.T, l *List, es []interface{}) {
	if !checkListLen(t, l, len(es)) {
		return
	}

	i := 0

	for e := l.First(); e != nil; e = e.Next() {
		le := e.Value.(int)
		if le != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
}

func TestExtending(t *testing.T) {
	l1 := New()
	l2 := New()

	l1.Push(1)
	l1.Push(2)
	l1.Push(3)

	l2.Push(4)
	l2.Push(5)

	l3 := New()
	l3.PushFrontList(l1)

	checkList(t, l3, []interface{}{1, 2, 3})

	l3.PushFrontList(l2)

	checkList(t, l3, []interface{}{1, 2, 3, 4, 5})

	l3 = New()
	l3.PushFrontList(l2)

	checkList(t, l3, []interface{}{4, 5})
	l3.PushFrontList(l1)

	checkList(t, l3, []interface{}{1, 2, 3, 4, 5})

	checkList(t, l1, []interface{}{1, 2, 3})
	checkList(t, l2, []interface{}{4, 5})

	l3 = New()
	l3.PushFrontList(l1)

	checkList(t, l3, []interface{}{1, 2, 3})

	l3.PushFrontList(l3)

	checkList(t, l3, []interface{}{1, 2, 3, 1, 2, 3})

	l3 = New()
	l3.PushFrontList(l1)

	checkList(t, l3, []interface{}{1, 2, 3})
	l3.PushFrontList(l3)

	checkList(t, l3, []interface{}{1, 2, 3, 1, 2, 3})
	l3 = New()
	l1.PushFrontList(l3)
	checkList(t, l1, []interface{}{1, 2, 3})
	l1.PushFrontList(l3)
	checkList(t, l1, []interface{}{1, 2, 3})
}

func TestRemove(t *testing.T) {
	l := New()
	e1 := l.Push(1)
	e2 := l.Push(2)

	checkListPointers(t, l, []*Element{e1, e2})
	e := l.First()
	l.Remove(e)

	checkListPointers(t, l, []*Element{e2})
	l.Remove(e)
	checkListPointers(t, l, []*Element{e2})
}

func TestIssue4103(t *testing.T) {
	l1 := New()
	l1.Push(1)
	l1.Push(2)

	l2 := New()
	l2.Push(3)
	l2.Push(4)

	e := l1.First()
	l2.Remove(e) // l2 should not change because e is not an element of l2
	if n := l2.Len(); n != 2 {
		t.Errorf("l2.Len() = %d, want 2", n)
	}

	l1.InsertBefore(8, e)
	if n := l1.Len(); n != 3 {
		t.Errorf("l1.Len() = %d, want 3", n)
	}
}

func TestIssue6349(t *testing.T) {
	l := New()
	l.Push(1)
	l.Push(2)

	e := l.First()
	l.Remove(e)
	if e.Value != 1 {
		t.Errorf("e.value = %d, want 1", e.Value)
	}
	if e.Next() != nil {
		t.Errorf("e.Next() != nil")
	}
	//if e.Prev() != nil {
	//	t.Errorf("e.Prev() != nil")
	//}
}

func TestMove(t *testing.T) {
	l := New()
	e1 := l.Push(1)
	e2 := l.Push(2)
	e3 := l.Push(3)
	e4 := l.Push(4)

	l.MoveAfter(e3, e3)
	checkListPointers(t, l, []*Element{e1, e2, e3, e4})
	l.MoveBefore(e2, e2)
	checkListPointers(t, l, []*Element{e1, e2, e3, e4})
	l.MoveAfter(e3, e2)
	checkListPointers(t, l, []*Element{e1, e2, e3, e4})

	l.MoveBefore(e2, e3)

	checkListPointers(t, l, []*Element{e1, e2, e3, e4})
	l.MoveBefore(e2, e4)
	checkListPointers(t, l, []*Element{e1, e3, e2, e4})
	e2, e3 = e3, e2

	l.MoveBefore(e4, e1)

	checkListPointers(t, l, []*Element{e4, e1, e2, e3})
	e1, e2, e3, e4 = e4, e1, e2, e3

	l.MoveAfter(e4, e1)
	checkListPointers(t, l, []*Element{e1, e4, e2, e3})
	e2, e3, e4 = e4, e2, e3

	l.MoveAfter(e2, e3)
	checkListPointers(t, l, []*Element{e1, e3, e2, e4})
	e2, e3 = e3, e2
}

// Test PushFront, Push, PushFrontList, PushFrontList() with uninitialized List
func TestZeroList(t *testing.T) {
	var l1 = new(List)
	l1.PushFront(1)
	checkList(t, l1, []interface{}{1})

	var l2 = new(List)
	l2.Push(1)
	checkList(t, l2, []interface{}{1})

	var l3 = new(List)
	l3.PushFrontList(l1)
	checkList(t, l3, []interface{}{1})

	var l4 = new(List)
	l4.PushFrontList(l2)
	checkList(t, l4, []interface{}{1})
}

// Test that a list l is not modified when calling InsertBefore with a mark that is not an element of l.
func TestInsertBeforeUnknownMark(t *testing.T) {
	var l List
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.InsertBefore(1, new(Element))
	checkList(t, &l, []interface{}{1, 2, 3})
}

// Test that a list l is not modified when calling InsertAfter with a mark that is not an element of l.
func TestInsertAfterUnknownMark(t *testing.T) {
	var l List
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.InsertAfter(1, new(Element))
	checkList(t, &l, []interface{}{1, 2, 3})
}

// Test that a list l is not modified when calling MoveAfter or MoveBefore with a mark that is not an element of l.
func TestMoveUnknownMark(t *testing.T) {
	var l1 List
	e1 := l1.Push(1)

	var l2 List
	e2 := l2.Push(2)

	l1.MoveAfter(e1, e2)
	checkList(t, &l1, []interface{}{1})
	checkList(t, &l2, []interface{}{2})

	l1.MoveBefore(e1, e2)
	checkList(t, &l1, []interface{}{1})
	checkList(t, &l2, []interface{}{2})
}

