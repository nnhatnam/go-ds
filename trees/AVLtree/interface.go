package AVLtree

//test a difference approach
type Interface interface {
	Less(other Interface) bool
	Equal(other Interface) bool
	Greater(other Interface) bool
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

func (i Int) Less(other Interface) bool {
	return i < other.(Int)
}

func (i Int) Equal(other Interface) bool {
	return i == other.(Int)
}

func (i Int) Greater(other Interface) bool {
	return i > other.(Int)
}

type TraverseFunc func(value Interface)
