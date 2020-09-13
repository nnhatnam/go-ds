package stack

//import (
//	"github.com/nnhatnam/go-ds/linearlists/dllist"
//)

// When we use stack, we care more about value instead of node, so node should be private
type node struct {
	next *node

	value interface{}
}

type list struct {
	head *node
	len int
}

func newList() *list { return new(list).Init()}

//Init initializes or clears list l.
func (l *list) Init() *list {
	l.head = nil
	l.len = 0
	return l
}

func (l *list) Front() interface{} {
	if l.head != nil {
		return l.head.value
	}
	return nil
}

func (l *list) PushFront(v interface{}) {
	newNode := &node{value : v,}
	if l.head == nil {
		l.head = newNode

	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.len++
}

func (l *list) Len() int {
	return l.len
}

func (l *list) RemoveFront() interface{} {
	if l.head != nil {
		node := l.head
		l.head = l.head.next
		l.len--
		node.next = nil
		return node.value
	}
	return nil
}

type Stack struct {
	list *list
}

func NewStack() *Stack {
	s := NewStack()
	s.list = newList()
	return s
}

func (s *Stack) Push(v interface{}) {
	s.list.PushFront(v)

}

func (s *Stack) Pop() interface{} {
	return s.list.RemoveFront()
}

func (s *Stack) Top() interface{} {
	return s.list.Front()
}

func (s *Stack) Size() int {
	return s.list.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}

