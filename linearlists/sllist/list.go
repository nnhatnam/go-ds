package sllist



// Element is a node of a linked list.
type Element struct {

	next *Element

	// The list to which this Element belongs.
	list *List

	//The value stored with this node
	value interface{}
}

func (n *Element) Next() *Element {
	return n.next
}

type List struct {
	head *Element
	tail *Element
	//root sentinel
	length int //  current list length excluding (this) sentinel Element
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	//this linked list use dummy node technical
	node := &Element{list: l, value: nil}
	l.head = node
	l.tail = node
	l.length = 0
	return l
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.head == nil {
		l.Init()
	}
}

func New() *List{
	return new(List).Init()
}

//return the real head of list l
func (l *List) listHead() *Element {
	return l.head.next
}

//return the tail of list l
func (l *List) listTail() *Element {
	return l.tail
}

func (l *List) append(value interface{})  {
	if l.head == nil {
		l.lazyInit()
	}
	l.tail.next = &Element{list: l, value:value}
	l.tail = l.tail.next
	l.length += 1

}

func (l *List) appendMany(values ...interface{}) *Element {
	if l.head == nil {
		l.lazyInit()
	}

	var elem *Element
	for _, v := range values {
		l.tail.next = &Element{list: l, value:v}
		l.tail = l.tail.next
		l.length += 1
		if elem == nil {
			elem = l.tail
		}
	}
	return elem
}



func (l *List) prepend(value interface{}) *Element {

	if l.head == nil {
		l.lazyInit()
	}
	node := &Element{
		list: l,
		next:l.head.next,
		value:value,
	}
	l.head.next = node
	l.length++
	return node
}

func (l *List) prependMany(values ...interface{}) *Element {
	if l.head == nil {
		l.lazyInit()
	}

	for _ , v := range values {
		l.head.next = &Element{
			list: l,
			next:l.head.next,
			value:v,
		}
		l.length++
	}
	return l.head.next
}

//Len returns the number of Elements of list l. The complexity is O(1).
func (l *List) Len() int { return l.length }

//First returns the first value of list l or nil if l is empty. The complexity is O(1)
func (l *List) First() interface{} {
	if l.length > 0 {
		return l.head.next.value
	}
	return nil
}

//FirstElement returns the first element of list l or nil if l is empty. The complexity is O(1)
func (l *List) FirstElement() *Element {
	return l.head.next
}

//Last returns the last value of list l or nil if l is empty. The complexity is O(1)
func (l *List) Last() interface{} {
	if l.length > 0 {
		return l.tail.value
	}
	return nil
}

func (l *List) LastElement() *Element {
	if l.length > 0 {
		return l.tail
	}
	return nil
}

func (l *List) Append(values ...interface{}) *Element {
	return l.appendMany(values)
}

func (l *List) Prepend(values ...interface{}) *Element {
	return l.prependMany(values)
}

//PushFront inserts a new Element with value v at the beginning of list l. The complexity is O(1)
func (l *List) PushFront(v interface{}) {
	l.prepend(v)
}

//Push inserts a new Element with value v at the end of list l. The complexity is O(1)
func (l *List) PushBack(v interface{})  {
	l.append(v)
}

//Pop removes the last Element from list l and returns that Element. The complexity is O(n)
func (l *List) PopBack() interface{} {
	if l.length > 0 {
		if e := l.remove(l.tail); e != nil {
			return e.value
		}
	}

	return nil
}

//Pop removes the first Element from list l and returns that Element. The complexity is O(n)
func (l *List) PopFront() interface{} {
	if l.length > 0 {
		if e := l.remove(l.head.next); e != nil {
			return e.value
		}
	}

	return nil
}

// remove removes the Element e from list l (if e belong to the list), decrements l.len, and returns e. The complexity is O(n)
func (l *List) remove(e *Element) *Element {

	for el := l.head; el != nil; el = el.next {
		if el.next == e {
			//el is previous Element of e
			el.next = e.next
			if l.tail == e {
				l.tail = el
			}
			e.next = nil
			e.list = nil
			l.length--
			return e
		}
	}

	return nil
}



type IterateFunc func(e *Element, position int)


func (l *List) Traverse(f IterateFunc){

	var i = 0
	for elem := l.head.next; elem != nil; elem = elem.next {
		f(elem, i)
		i++
	}

}


type lookupFunc func(e *Element) bool //private find by node (index doesn't need, because we can access it in private level)

func (l *List) lookup(f lookupFunc) *Element {

	//loop from dummy node
	for e := l.head; e != nil; e = e.next {
		if f(e) {
			return e
		}
	}
	return nil
}






//insert an Element e after the Element at, increments l.len, and return n, update tail if needed. Element at must not be nil
func (l *List) insert(e, at *Element) *Element {
	if e.list != at.list {
		return nil
	}
	if l.tail == at {
		l.tail = e
	}
	e.next = at.next
	at.next = e

	l.length++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{value:v,}, at)
}

// InsertAfter inserts a new Element e with value v immediately after an Element that mark iterator pointing in and returns the value.
// If mark is not an Element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {

	if mark.list != l {
		return nil
	}

	return l.insertValue(v, mark)
}


//Note later
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {

	if mark.list != l {
		return nil
	}

	// Find node before iter.Element
	elem := l.lookup(func(n *Element) bool {
		return n.next == mark
	})

	return l.insertValue(v, elem)
}

//move node n after node at
func (l *List) move(n, at *Element) *Element {

	if n == at || at.next == n {
		return n
	}

	if l.head.next == n {
		l.head.next = n.next

	} else {
		nodeBeforeN := l.lookup(func(node *Element) bool {
			return node.next == n
		})
		nodeBeforeN.next = n.next

	}

	afterAt := at.next
	n.next = afterAt
	at.next = n

	if l.tail == at {
		l.tail = n
	}
	return n
}

//need to re-implement
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// MoveBefore moves node n to its new position before mark.
// If n or mark is not an node of l, or n == mark, the list is not modified.
// The node and mark must not be nil.
func (l *List) MoveBefore(e, mark *Element){
	if e.list != l || e == mark || mark.list != l {
		return
	}

	nodeBeforeMark := l.lookup(func(n *Element) bool {
		return n.next == mark
	})

	l.move(e, nodeBeforeMark)
}

//need implement
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.tail == e {
		return
	}
	l.move(e, l.tail)
}

//need implement
func (l *List) MoveToFront(e *Element){


	if e.list != l || l.head.next == e {
		return
	}

	prev := l.lookup(func(node *Element) bool {
		return node.next == e
	})

	prev.next = e.next
	e.next = l.head.next

	if e == l.tail {
		l.tail = prev
	}
	l.head = e
}



//need implement
func (l *List) PushBackList(other *List){
	l.lazyInit()
	//can't use traverse due to infinite loop if the list push back itself
	for i, elem := other.Len(), other.head.next; i > 0; i, elem = i - 1,  elem.next {
		if l.head.next == nil {
			l.append(elem.value)
		} else {
			l.insertValue(elem.value, l.tail )
		}
		//nodes = append(nodes, node)
	}
	//fmt.Println("out ", l.length)
}


//need implement
func (l *List) PushFrontList(other *List){
	l.lazyInit()
	tail := l.tail
	for i, elem := other.Len(), other.head.next; i > 0; i, elem = i - 1,  elem.next {
		l.append(elem.value)
		//nodes = append(nodes, node)
	}

	if tail != nil && l.tail != tail {
		l.tail.next = l.head.next
		l.head.next = tail.next
		tail.next = nil
	}
}



//need implement
func (l *List) Remove(e *Element) interface{}{
	if e.list == l {
		l.remove(e)
	}

	return e.value
}

//Removes all the Elements from this list.
func (l *List) Clear() {
	l.Init()
}


