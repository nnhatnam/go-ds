package singlylinkedlist

// Node is a node of a linked list.
type Node struct {

	next *Node

	// The list to which this element belongs.
	list *List

	//The value stored with this node
	Value interface{}
}



func (n *Node) Next() *Node {
	return n.next
}

//
//func (n *Node) Prev() *Node {
//	return nil
//}