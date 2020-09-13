package stack

import (
	"github.com/nnhatnam/go-ds/list/dllist"
)

type Stack struct {
	list *dllist.List
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(value interface{}) {
	if s.list != nil {
		s.list = dllist.New()
	}

	s.list.Prepend(value)

}

func (s *Stack) Pop() interface{} {
	if s.list == nil {
		return nil
	}

	return s.list.Remove(s.list.First())
}

func (s *Stack) Size() int {
	return s.list.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}

