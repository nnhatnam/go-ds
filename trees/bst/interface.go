package bst

type Interface interface {
	Compare(other Interface) int //-1 0 1
}

type Int int

func (i Int) Compare(other Interface) int {
	o := other.(Int)
	if i > o {
		return 1
	} else if i == o {
		return 0
 	} else {
 		return -1
	}
}

type TraverseFunc func(value Interface)