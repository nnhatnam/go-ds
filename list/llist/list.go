package llist

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

type sentinel struct {
	head *Node
	tail *Node
}


type List struct {
	//head *Node // sentinel list element, only &root, root.prev, and root.next are used

	root sentinel
	length int //  current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List) init() *List {
	l.root.head = nil
	l.root.tail = nil
	l.length = 0
	return l
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.root.head == nil {
		l.init()
	}
}

func New(values ...interface{}) *List{
	list := (&List{})
	if len(values) > 0 {
		list.appendMany(values...)
	}
	return list
}

func (l *List) append(value interface{}) * Node {

	node := &Node{Value:value, list:l}
	if l.length == 0 {
		l.root.head , l.root.tail = node, node

	} else {
		l.root.tail.next = node
		l.root.tail = node
	}

	l.length += 1

	return node

}

func (l *List) appendMany(values ...interface{}) {
	for _, val := range values {
		l.append(val)
	}

	//fmt.Println("root after: ",l.root)
}

func (l *List) prepend(value interface{}) *Node{
	node := &Node{
		next:l.root.head,
		Value:value,
		list:l,
	}
	if l.length == 0 {
		l.root.tail = node
	}
	l.root.head = node
	l.length++
	return node
}

func (l *List) Size() int { return l.length }

func (l *List) First() *Node {
	return l.root.head
}

func (l *List) Last() *Node {
	return l.root.tail
}

//Append a node to the end of the list O(1)
func (l* List) Append(value interface{}) *Node {
	return l.append(value)

}

// Prepend a node to the beginning of the list
func (l *List) Prepend(value interface{}) *Node {
	return l.prepend(value)
}


type lookupFunc func(n *Node) bool //private find by node (index doesn't need, because we can access it in private level)
type IterateFunc func(n *Node, position int) bool //find by value, add index for more flexible

//the same function like findOne, but doesn't provide index
func (l *List) lookup(f lookupFunc) *Node {

	for node := l.First(); node != nil; node = node.Next() {
		if f(node) {
			return node
		}
	}
	return nil
}

func (l *List) Traverse(f IterateFunc){
	var i = 0
	for node := l.First(); node != nil; node = node.Next() {
		f(node,i)
		i++
	}

}


//insert node n after node at, increments l.len, and return n, update tail if needed
func (l *List) insert(n, after *Node) *Node {

	n.list = l
	n.next = after.next
	after.next = n

	if n.next == nil {
		l.root.tail = n.next
	}
	l.length++
	return n
}

func (l *List) insertValue(v interface{}, after *Node) *Node {
	return l.insert(&Node{Value:v,}, after)
}

//need implement
func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}


//Note later
func (l *List) InsertBefore(v interface{}, mark *Node) *Node {
	if mark.list != l {
		return nil
	}

	nodeBeforeMark := l.lookup(func(n *Node) bool {
		return n.next == mark
	})

	return l.insertValue(v, nodeBeforeMark)
}

//move node n after node at
func (l *List) move(n, after *Node) *Node {
	if n.list != l || after.list != l {
		return nil
	}
	if n == after || after.next == n {
		return n
	}
	nodeBeforeN := l.lookup(func(node *Node) bool {
		return node.next == n
	})
	nodeBeforeN.next = n.next
	n.next = after.next
	after.next = n
	if l.root.tail == after {
		l.root.tail = n
	}
	return n
}

//need to re-implement
func (l *List) MoveAfter(n, mark *Node) {
	if n.list != l || n == mark || mark.list != l {
		return
	}
	l.move(n, mark)
}

//need implement
func (l *List) MoveBefore(n, mark *Node){
	if n.list != l || n == mark || mark.list != l {
		return
	}
	temp := &Node{next:l.root.head}
	l.root.head = temp
	nodeBeforeMark := l.lookup(func(n *Node) bool {
		return n.next == mark
	})
	l.move(n, nodeBeforeMark)
	l.root.head = l.root.head
	temp.next = nil
}

//need implement
func (l *List) MoveToBack(n *Node) {
	if n.list != l || l.root.tail == n {
		return
	}
	l.move(n, l.root.tail)
}

//need implement
func (l *List) MoveToFront(n *Node){
	if n.list != l || l.root.tail == n {
		return
	}
}

//need implement
func (l *List) PushBack(v interface{}) *Node {
	return l.append(v)
}

//need implement
func (l *List) PushBackList(other *List){
	l.lazyInit()
	for node := other.First(); node != nil; node = node.Next() {
		l.insertValue(node.Value, l.root.tail )
		//nodes = append(nodes, node)
	}
}

//need implement
func (l *List) PushFront(v interface{}) *Node {
	return l.prepend(v)
}

//need implement
func (l *List) PushFrontList(other *List){
	l.lazyInit()
	for node := other.First(); node != nil; node = node.Next() {
		//l.insertValue(node.Value, l.root )
		l.append(node.Value)
		//nodes = append(nodes, node)
	}
}

func (l *List) remove(n *Node) *Node {
	node := l.lookup(func(n *Node) bool {
		return n.next == n
	})
	node.next = n.next
	n.next = nil
	n.list = nil
	return n
}

//need implement
func (l *List) Remove(n *Node) interface{}{
	if n.list == l {
		l.remove(n)
	}

	return n.Value
}

//need implement
func (l *List) Clear() {
	l.lazyInit()
}


