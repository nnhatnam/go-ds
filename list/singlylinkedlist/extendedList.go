package singlylinkedlist

type FlexList struct {
	*List
}

func NewFlexibleList(values ...interface{}) *FlexList {
	l := New(values)
	return &FlexList{List:l,}
}