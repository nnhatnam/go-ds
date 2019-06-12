package singlylinkedlist

// Node is a node of a linked list.
type Node struct {

	next *Node

	// The list to which this element belongs.
	//list *List

	//The value stored with this node
	Value interface{}
}

func newNode(value interface{}) *Node{
	return &Node{Value:value}
}

type List struct {
	head *Node // sentinel list element, only &root, root.prev, and root.next are used
	tail *Node
	len int //  current list length excluding (this) sentinel element
}

func New(values ...interface{}) *List{
	list := new(List)
	if len(values) > 0 {
		list.appendMany(values)
	}
	return list
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int { return l.len }

func (l *List) Front() *Node {
	return l.head
}


func (l *List) Back() *Node {
	return l.tail
}

func (l *List) appendOne(value interface{}){
	if l.head == nil {
		l.head = newNode(value)
		l.tail = l.head
		l.len++
		return
	}

	node := l.head
	for node.next != nil {
		node = node.next
	}

	node.next = newNode(value)
	l.tail = node.next
	l.len++
}

func (l *List) appendMany(values ...interface{}) {
	for _, val := range values {
		l.appendOne(val)
	}
}

//Append a node to the end of the list O(1)
func (l* List) Append(value interface{}){
	l.appendOne(value)
}

func (l *List) prepend(value interface{}){
	if l.head == nil {
		l.head = newNode(value)
		l.tail = l.head
		l.len++
		return
	}

	newHead := newNode(value)
	newHead.next = l.head
	l.head = newHead
	l.len++
}

// Prepend a node to the beginning of the list
func (l *List) Prepend(value interface{}){
	l.prepend(value)
}

