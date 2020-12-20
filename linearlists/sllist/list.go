package sllist

// Element is a node of a linked list.
type Element struct {

	next *Element

	// The list to which this Element belongs.
	list *List

	//The value stored with this element
	Value interface{}
}

//Next returns the next list element or nil
func (n *Element) Next() *Element {
	return n.next
}

//List represents a singly linked list.
type List struct {

	//this linked list use dummy element technical

	head *Element //head points to the dummy element
	tail *Element //tail points to the last list element or dummy element if the list is empty

	len int //  current list len excluding (this) dummy Element
}

//Len returns the number of Elements of list l. The complexity is O(1).
func (l *List) Len() int { return l.len }

//Front returns the first element of list l or nil if l is empty. The time complexity is O(1)
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

//Back returns the last element of list l or nil if l is empty. The time complexity is O(1)
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}

	return l.tail
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.head == nil {
		l.Init()
	}
}

// Init initializes or clears list l.
func (l *List) Init() *List {

	node := &Element{list: l, Value: nil}
	l.head, l.tail, l.len = node , node, 0
	return l
}

// New returns an initialized list.
func New() *List{
	return new(List).Init()
}


//insert an Element e after the Element at, increments l.len, and return n, update tail if needed.
//Element at must not be nil. Element at must not belong to list l, so any call to l.insert must check if at NOT belong
//to list l
func (l *List) insert(e, at *Element) *Element {
	e.next = at.next
	at.next = e
	l.len++

	if l.tail == at {
		l.tail = e
	}
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value:v, list: l}, at)
}

// remove removes the Element e from list l (if e belong to the list), decrements l.len, and returns e. The time complexity is O(n)
func (l *List) remove(e *Element) *Element {

	//iterate over the list l, starting from dummy element, to find the element el before element e then remove e and update
	//el reference
	for el := l.head; el != nil; el = el.next {
		if el.next == e {
			//el is previous Element of e
			el.next = e.next

			if l.tail == e {
				l.tail = el
			}
			e.next = nil
			e.list = nil
			l.len--
			return e
		}
	}

	return nil
}

type lookupFunc func(e *Element) bool

//lookup loops through list l from dummy element and call function f for each visited element.
//If f() return true then stop the loop and return the element
func (l *List) lookup(f lookupFunc) *Element {

	//loop from dummy node
	for e := l.head; e != nil; e = e.next {
		if f(e) {
			return e
		}
	}
	return nil
}


//move moves e next to at and return e. e and at must belong to the list l
func (l *List) move(e, at *Element) *Element {

	if e == at || at.next == e {
		return e
	}

	if l.head.next == e {
		l.head.next = e.next

	} else {
		nodeBeforeE := l.lookup(func(node *Element) bool {
			return node.next == e
		})
		nodeBeforeE.next = e.next

	}

	afterAt := at.next
	e.next = afterAt
	at.next = e

	if l.tail == at {
		l.tail = e
	}
	return e
}

//append appends value into list l. List l must be initialized
func (l *List) append(value interface{}) *Element {

	e := &Element{list: l, Value:value}
	l.tail.next = e
	l.tail = l.tail.next
	l.len += 1
	return e
}

func (l *List) appendMany(values ...interface{}) *Element {

	var elem *Element
	for _, v := range values {
		l.tail.next = &Element{list: l, Value:v}
		l.tail = l.tail.next
		l.len += 1
		if elem == nil {
			elem = l.tail
		}
	}
	return elem
}

//prepend insert value into the front of list l, after dummy element. List l must be initialized
func (l *List) prepend(value interface{}) *Element {

	e := &Element{
		list: l,
		next:l.head.next,
		Value:value,
	}

	l.head.next = e
	if l.len == 0 {
		l.tail = e
	}
	l.len++
	return e
}

func (l *List) prependMany(values ...interface{}) *Element {

	for _ , v := range values {
		l.head.next = &Element{
			list: l,
			next:l.head.next,
			Value:v,
		}
		l.len++
	}
	return l.head.next
}


//First returns the first value of list l or nil if l is empty. The complexity is O(1)
func (l *List) First() interface{} {
	if l.len > 0 {
		return l.head.next.Value
	}
	return nil
}


//Last returns the last value of list l or nil if l is empty. The complexity is O(1)
func (l *List) Last() interface{} {
	if l.len > 0 {
		return l.tail.Value
	}
	return nil
}



func (l *List) Append(values ...interface{}) *Element {
	l.lazyInit()
	return l.appendMany(values)
}

func (l *List) Prepend(values ...interface{}) *Element {
	l.lazyInit()
	return l.prependMany(values)
}

//PushFront inserts a new Element with value v at the beginning of list l. The complexity is O(1)
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.prepend(v)
}

//Push inserts a new Element with value v at the end of list l. The complexity is O(1)
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.append(v)
}

//Pop removes the last Element from list l and returns that Element. The complexity is O(n)
func (l *List) PopBack() interface{} {
	if l.len > 0 {
		if e := l.remove(l.tail); e != nil {
			return e.Value
		}
	}

	return nil
}

//Pop removes the first Element from list l and returns that Element. The complexity is O(n)
func (l *List) PopFront() interface{} {
	if l.len > 0 {
		if e := l.remove(l.head.next); e != nil {
			return e.Value
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

// InsertAfter inserts a new Element e with value v immediately after an Element that mark iterator pointing in and returns the value.
// If mark is not an Element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {

	if mark.list != l {
		return nil
	}

	return l.insertValue(v, mark)
}


// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
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



// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
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
	l.head.next = e
}



//need implement
func (l *List) PushBackList(other *List){
	l.lazyInit()
	//can't use traverse due to infinite loop if the list push back itself
	for i, elem := other.Len(), other.head.next; i > 0; i, elem = i - 1,  elem.next {
		if l.head.next == nil {
			l.append(elem.Value)
		} else {
			l.insertValue(elem.Value, l.tail )
		}
		//nodes = append(nodes, node)
	}
	//fmt.Println("out ", l.len)
}


// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List){
	l.lazyInit()
	//tail := l.tail
	pointer := l.head
	for i, elem := other.Len(), other.head.next; i > 0; i, elem = i - 1,  elem.next {

		l.insertValue(elem.Value, pointer)
		pointer = pointer.next
		//nodes = append(nodes, node)
	}

	//if tail != nil && l.tail != tail {
	//	l.tail.next = l.head.next
	//	l.head.next = tail.next
	//	tail.next = nil
	//}
}



// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List) Remove(e *Element) interface{}{
	if e.list == l {
		l.remove(e)
	}

	return e.Value
}

//Removes all the Elements from this list.
func (l *List) Clear() {
	l.Init()
}


