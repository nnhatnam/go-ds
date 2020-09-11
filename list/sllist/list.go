package sllist

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

func (l *List) Len() int { return l.length }

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
type IterateFunc func(n *Node, position int) //find by value, add index for more flexible

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


//insert node n after node "at", increments l.len, and return n, update tail if needed. Node "at" is nil, do nothing
func (l *List) insert(n, at *Node) *Node {
	if at == nil {
		return nil
	}

	if l.root.tail == at {
		l.root.tail = n
	}
	afterAt := at.next
	at.next = n
	n.next = afterAt
	n.list = l
	//n.next = at.next
	//after.next = n


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

	if l.root.head == mark {
		return l.prepend(v)
	}

	nodeBeforeMark := l.lookup(func(n *Node) bool {
		return n.next == mark
	})

	return l.insertValue(v, nodeBeforeMark)
}

//move node n after node at
func (l *List) move(n, at *Node) *Node {
	if n.list != l || at.list != l {
		return nil
	}
	if n == at || at.next == n {
		return n
	}

	if l.root.head == n {
		l.root.head = n.next

	} else {
		nodeBeforeN := l.lookup(func(node *Node) bool {
			return node.next == n
		})
		nodeBeforeN.next = n.next

	}

	afterAt := at.next
	n.next = afterAt
	at.next = n

	if l.root.tail == at {
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

// MoveBefore moves node n to its new position before mark.
// If n or mark is not an node of l, or n == mark, the list is not modified.
// The node and mark must not be nil.
func (l *List) MoveBefore(n, mark *Node){
	if n.list != l || n == mark || mark.list != l || n.next == mark {
		return
	}
	temp := &Node{next:l.root.head, list:l,}
	l.root.head = temp
	nodeBeforeMark := l.lookup(func(n *Node) bool {
		return n.next == mark
	})
	l.move(n, nodeBeforeMark)
	l.root.head = l.root.head.next
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
	if n.list != l || l.root.head == n {
		return
	}
	prev := l.lookup(func(node *Node) bool {
		return node.next == n
	})
	prev.next = n.next
	n.next = l.root.head
	if n == l.root.tail {
		l.root.tail = prev
	}
	l.root.head = n
}

//need implement
func (l *List) PushBack(v interface{}) *Node {
	return l.append(v)
}

//need implement
func (l *List) PushBackList(other *List){
	l.lazyInit()
	//can't use traverse due to infinite loop if the list push back itself
	for i, node := other.Len(), other.First(); i > 0; i, node = i - 1,  node.Next() {
		if l.root.head == nil {
			l.append(node.Value)
		} else {
			l.insertValue(node.Value, l.root.tail )
		}
		//nodes = append(nodes, node)
	}
	//fmt.Println("out ", l.length)
}

//need implement
func (l *List) PushFront(v interface{}) *Node {
	return l.prepend(v)
}

//need implement
func (l *List) PushFrontList(other *List){
	l.lazyInit()
	tail := l.root.tail
	for i, node := other.Len(), other.First(); i > 0; i, node = i - 1,  node.Next() {
		l.append(node.Value)
		//nodes = append(nodes, node)
	}

	if tail != nil && l.root.tail != tail {
		l.root.tail.next = l.root.head
		l.root.head = tail.next
		tail.next = nil
	}
}

func (l *List) remove(n *Node) *Node {
	if l.length == 0 || n.list != l {
		return nil
	}

	if l.root.head == n {
		l.root.head = n.next
		if l.root.tail == n {
			l.root.tail = nil
		}
	} else {

		node := l.lookup(func(e *Node) bool {
			return e.next == n
		})

		node.next = n.next
		if l.root.tail == n {
			l.root.tail = node
		}

	}

	n.next = nil
	n.list = nil
	l.length -= 1
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


