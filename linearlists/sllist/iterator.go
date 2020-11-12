package sllist

type Node struct {

	*element

	list *List

}

func (iter *Node) End() bool {
	if iter.element == nil {
		return true
	}
	return false
}


func (iter *Node) Next() bool {
	if iter.element == nil {
		return false
	}
	iter.element = iter.element.next
	return true
}

func (iter *Node) Value() interface{} {

	if iter.element == nil {
		return nil
	}
	return iter.element.value
}