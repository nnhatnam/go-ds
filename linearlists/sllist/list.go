package sllist

type Iterator struct {

	*element

	list *List

}

func (iter *Iterator) Next() bool {
	if iter.element == nil {
		return false
	}
	iter.element = iter.element.next
	return true
}

func (iter *Iterator) Value() interface{} {

	if iter.element == nil {
		return nil
	}
	return iter.element.value
}

// element is a node of a linked list.
type element struct {

	next *element

	// The list to which this element belongs.
	//list *List

	//The value stored with this node
	value interface{}
}

//func (n *element) Next() *element {
//	return n.next
//}

type List struct {
	head *element
	tail *element
	//root sentinel
	length int //  current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	//this linked list use dummy node technical
	node := &element{value: nil}
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

func (l *List) append(value interface{})  {
	if l.head == nil {
		l.lazyInit()
	}
	l.tail.next = &element{value:value}
	l.length += 1

}


func (l *List) prepend(value interface{}) *element {
	node := &element{
		next:l.head.next,
		value:value,
	}
	l.head.next = node
	l.length++
	return node
}

//Len returns the number of elements of list l. The complexity is O(1).
func (l *List) Len() int { return l.length }

//First returns the first value of list l or nil if l is empty. The complexity is O(1)
func (l *List) First() interface{} {
	return l.head.next.value
}

//Last returns the last value of list l or nil if l is empty. The complexity is O(1)
func (l *List) Last() interface{} {
	if l.length > 0 {
		return l.tail.value
	}
	return nil
}

//PushFront inserts a new element with value v at the beginning of list l. The complexity is O(1)
func (l *List) PushFront(v interface{}) {
	l.prepend(v)
}

//Push inserts a new element with value v at the end of list l. The complexity is O(1)
func (l *List) PushBack(v interface{})  {
	l.append(v)
}

//Pop removes the last element from list l and returns that element. The complexity is O(n)
func (l *List) PopBack() interface{} {
	if l.length > 0 {
		if e := l.remove(l.tail); e != nil {
			return e.value
		}
	}

	return nil
}

//Pop removes the first element from list l and returns that element. The complexity is O(n)
func (l *List) PopFront() interface{} {
	if l.length > 0 {
		if e := l.remove(l.head.next); e != nil {
			return e.value
		}
	}

	return nil
}

// remove removes the element e from list l (if e belong to the list), decrements l.len, and returns e. The complexity is O(n)
func (l *List) remove(e *element) *element {

	for el := l.head; el != nil; el = el.next {
		if el.next == e {
			//el is previous element of e
			el.next = e.next
			if l.tail == e {
				l.tail = el
			}
			e.next = nil
			l.length--
			return e
		}
	}

	return nil
}

//Iterator return an iterator pointing to the first element
func (l *List) Iterator() Iterator {
	return Iterator{
		element: l.head.next,
		list:    l,
	}
}



type IterateFunc func(e *element, position int)

//
//func (l *List) Traverse(f IterateFunc){
//	var i = 0
//	for node := l.First(); node != nil; node = node.Next() {
//		f(node,i)
//		i++
//	}
//}


type lookupFunc func(e *element) bool //private find by node (index doesn't need, because we can access it in private level)

func (l *List) lookup(f lookupFunc) *element {

	//loop from dummy node
	for e := l.head; e != nil; e = e.next {
		if f(e) {
			return e
		}
	}
	return nil
}






//insert an element e after the element at, increments l.len, and return n, update tail if needed. element at must not be nil
func (l *List) insert(e, at *element) *element {

	if l.tail == at {
		l.tail = e
	}
	e.next = at.next
	at.next = e

	l.length++
	return e
}

// insertValue is a convenience wrapper for insert(&element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *element) *element {
	return l.insert(&element{value:v,}, at)
}

// InsertAfter inserts a new element e with value v immediately after an element that mark iterator pointing in and returns the value.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(iter *Iterator, values ...interface{}) {

	if iter.list != l {
		return
	}

	elem := iter.element

	for _, v := range values {
		elem = l.insert(&element{value:v,}, elem)
	}
}


//Note later
func (l *List) InsertBefore(iter *Iterator , values ...interface{}) {
	if iter.list != l {
		return
	}

	// Find node before iter.element
	elem := l.lookup(func(n *element) bool {
		return n.next == iter.element
	})

	for _, v := range values {
		elem = l.insert(&element{value:v,}, elem)
	}

}

//move node n after node at
func (l *List) move(n, at *element) *element {

	if n == at || at.next == n {
		return n
	}

	if l.head.next == n {
		l.head.next = n.next

	} else {
		nodeBeforeN := l.lookup(func(node *element) bool {
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
func (l *List) MoveAfter(iter, mark *Iterator) {
	if iter.list != l || iter.element == mark.element || mark.list != l {
		return
	}
	l.move(iter.element, mark.element)
}

// MoveBefore moves node n to its new position before mark.
// If n or mark is not an node of l, or n == mark, the list is not modified.
// The node and mark must not be nil.
func (l *List) MoveBefore(iter, mark *Iterator){
	if iter.list != l || iter.element == mark.element || mark.list != l || iter.element.next == mark.element {
		return
	}
	nodeBeforeMark := l.lookup(func(n *element) bool {
		return n.next == mark.element
	})

	l.move(iter.element, nodeBeforeMark)

}

//need implement
func (l *List) MoveToBack(iter *Iterator) {
	if iter.list != l || l.tail == iter.element || iter.element == nil  {
		return
	}
	l.move(iter.element, l.tail)
}

//need implement
func (l *List) MoveToFront(iter *Iterator){

	elem := iter.element
	if iter.list != l || l.head.next == elem || elem == nil {
		return
	}

	prev := l.lookup(func(node *element) bool {
		return node.next == elem
	})

	prev.next = elem.next
	elem.next = l.head.next

	if elem == l.tail {
		l.tail = prev
	}
	l.head = elem
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
func (l *List) Remove(n *Iterator) interface{}{
	if n.list == l {
		l.remove(n.element)
	}

	return n.Value
}

//Removes all the elements from this list.
func (l *List) Clear() {
	l.Init()
}


