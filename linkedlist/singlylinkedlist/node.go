package singlylinkedlist

// Node is a node of a linked list.
type Node struct {

	next *Node

	// The list to which this element belongs.
	list *List

	//The value stored with this node
	Value interface{}
}

func newNode(value interface{}, list *List) *Node{
	return &Node{Value:value, list:list}
}


//need implement
func (n *Node) Next() *Node {
	return nil
}

//need implement
func (n *Node) Prev() *Node {
	return nil
}