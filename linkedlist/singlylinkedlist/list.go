package singlylinkedlist

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

func (l *List) Back() *Node {
	return l.tail
}

func (l *List) Front() *Node {
	return l.head
}

type LookupFunc func(n *Node) bool

func (l *List) findOne(f LookupFunc) *Node{
	for node := l.Front(); node != nil; node = node.Next() {
		if f(node) {
			return node
		}
	}
	return nil
}

func (l *List) Find(f LookupFunc) *Node {
	return l.findOne(f)
}


//Append a node to the end of the list O(1)
func (l* List) Append(value interface{}){
	l.appendOne(value)
}

func (l *List) appendOne(value interface{}){
	node := newNode(value,l)
	if l.head == nil {
		l.head, l.tail = node, node
		l.len++
		return
	}

	l.tail.next = node
	l.tail = node
	return
}

func (l *List) appendMany(values ...interface{}) {
	for _, val := range values {
		l.appendOne(val)
		l.len++
	}
}

// Prepend a node to the beginning of the list
func (l *List) Prepend(value interface{}){
	l.prepend(value)
}


func (l *List) prepend(value interface{}) *Node{
	node := &Node{
		next:l.head,
		Value:value,
		list:l,
	}
	l.head = node
	if l.head.next == nil {
		l.tail = node
		return l.head
	}

	l.len++
	return node
}


//insert node n after node at, increments l.len, and return n, update tail if needed
func (l *List) insert(n, at *Node) *Node {
	n.list = l
	n.next = at.next
	at.next = n
	if n.next == nil {
		l.tail = n.next
	}
	l.len++
	return n
}

func (l *List) insertValue(v interface{}, at *Node) *Node {
	return l.insert(&Node{Value:v,}, at)
}

//need implement
func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	return l.insertValue(v, mark)
}


//Note later
func (l *List) InsertBefore(v interface{}, mark *Node) *Node {
	if l.head == mark {
		return l.prepend(mark)
	}

	nodeBeforeMark := l.findOne(func(n *Node) bool {
		return n.next == mark
	})

	return l.insertValue(v, nodeBeforeMark)
}

//need to re-implement
func (l *List) MoveAfter(n, mark *Node) {
	if n.list != l || n == mark || mark.list != l {
		return
	}
	l.findOne(func(n *Node) bool {
		return n.next == n || n.next == mark
	})
}

//need implement
func (l *List) MoveBefore(n, mark *Node){

}

//need implement
func (l *List) MoveToBack(n *Node) {

}

//need implement
func (l *List) MoveToFront(n *Node){

}

//need implement
func (l *List) PushBack(v interface{}) *Node {
	return nil
}

//need implement
func (l *List) PushBackList(other *List){

}

//need implement
func (l *List) PushFront(v interface{}) *Node {
	return nil
}

//need implement
func (l *List) PushFrontList(other *List){

}

//need implement
func (l *List) Remove(n *Node) interface{}{
	return nil
}

//need implement
func (l *List) Clear() {

}


//need implement with cache
func (l *List) NodeAt(index int) *Node {
	return nil
}

