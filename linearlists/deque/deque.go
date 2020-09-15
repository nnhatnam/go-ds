package deque

type (

	node struct {
		next, prev *node

		value interface{}
	}

 	list struct {
		head, tail *node
		len int
	}

)


func newList() *list { return new(list).Init()}

//Init initializes or clears list l.
func (l *list) Init() *list {
	l.head = nil
	l.tail = nil
	l.len = 0
	return l
}

func (l *list) Front() interface{} {
	if l.head != nil {
		return l.head.value
	}
	return nil
}

func (l *list) Back() interface{} {
	if l.head != nil {
		return l.head.value
	}
	return nil
}

func (l *list) PushFront(v interface{}) {
	newNode := &node{value : v,}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.len++
}

func (l *list) PushBack(v interface{}) {
	newNode := &node{value : v,}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
	l.len++
}

func (l *list) Len() int {
	return l.len
}

func (l *list) RemoveFront() interface{} {
	if l.head != nil && l.head == l.tail {
		v := l.head.value
		l.Init()
		return v
	}

	if l.head != nil {
		node := l.head
		l.head = l.head.next
		l.len--
		node.next = nil
		return node.value
	}
	return nil
}

func (l *list) RemoveBack() interface{} {
	if l.head != nil && l.head == l.tail {
		v := l.head.value
		l.Init()
		return v
	}
	if l.tail != nil {
		node := l.tail
		l.tail = l.tail.prev
		l.len--
		node.prev = nil
		return node.value
	}
	return nil
}


type Deque struct {
	list *list  //underline data struct
}

func NewDeque() *Deque {
	d := new(Deque)
	d.list.Init()
	return d
}

func (d *Deque) InsertFront(v interface{}) {
	d.list.PushFront(v)
}

func (d *Deque) InsertRear(v interface{}) {
	d.list.PushBack(v)
}

func (d *Deque) DeleteFront() interface{}{
	return d.list.RemoveFront()
}

func (d *Deque) DeleteRear() interface{}{
	return d.list.RemoveBack()
}

func (d *Deque) Front() interface{} {
	return d.list.head.value
}

func (d *Deque) Rear() interface{} {
	return d.list.tail.value
}

func (d *Deque) Clear() {
	d.list.Init()
}