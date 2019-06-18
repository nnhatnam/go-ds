package singlylinkedlist

import (
	//"fmt"
	"testing"

	//"github.com/seakingj/algorithms"
)

func TestNewList(t *testing.T) {
	list := New(1, "b")

	if actualValue := list.Len(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}
