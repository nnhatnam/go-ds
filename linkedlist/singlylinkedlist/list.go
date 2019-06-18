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

//need implement
func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	return nil
}

//need implement
func (l *List) InsertBefore(v interface{}, mark *Node) *Node {
	return nil
}

//need implement
func (l *List) MoveAfter(n, mark *Node) {

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

//Append a node to the end of the list O(1)
func (l* List) Append(value interface{}){
	l.appendOne(value)
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
		l.len++
	}
}

// Prepend a node to the beginning of the list
func (l *List) Prepend(value interface{}){
	l.prepend(value)
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

//need implement with cache
func (l *List) NodeAt(index int) *Node {
	return nil
}

