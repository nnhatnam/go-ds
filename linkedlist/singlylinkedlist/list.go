package singlylinkedlist

import "fmt"

type List struct {
	//head *Node // sentinel list element, only &root, root.prev, and root.next are used
	root *Node //root.next is the head
	tail *Node
	len int //  current list length excluding (this) sentinel element
}

func debug(format string, a ...interface{}){
	fmt.Printf(format,a...)
	fmt.Println("")
}

func (l *List) Init() *List {
	//l.root = newNode(nil,&l)
	node := &Node{list:l}

	l.root = node
	l.tail = l.root
	l.len = 0
	return l
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func New(values ...interface{}) *List{
	list := (&List{}).Init()
	if len(values) > 0 {
		list.appendMany(values...)
	}
	return list
}

//for testing purpose only
func (l *List) nodes() (nodes []*Node) {
	if l.root != nil {
		for node := l.Front(); node != nil; node = node.Next() {
			nodes = append(nodes, node)
		}
	}
	return
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int { return l.len }

func (l *List) root_() *Node {
	return l.root
}

func (l *List) Back() *Node {
	if l.len == 0 {
		return nil
	}
	return l.tail
}

func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}


func (l *List) appendOne(value interface{}) * Node {

	node := &Node{Value:value, list:l}
	if l.root == l.tail {
		l.root.next = node
	}

	l.tail.next = node
	l.tail = l.tail.next
	l.len += 1
	return node

}

func (l *List) appendMany(values ...interface{}) {
	for _, val := range values {
		l.appendOne(val)
	}

	//fmt.Println("root after: ",l.root)
}

// Prepend a node to the beginning of the list
//func (l *List) Prepend(value interface{}){
//	l.prepend(value)
//}


func (l *List) prepend(value interface{}) *Node{
	node := &Node{
		next:l.root.next,
		Value:value,
		list:l,
	}
	l.root.next = node
	if l.root == l.tail {
		l.tail = node
	}
	//if l.head.next == nil {
	//	l.tail = node
	//}

	l.len++
	return node
}


type lookupFunc func(n *Node) bool //private find by node (index doesn't need, because we can access it in private level)
type LookupFunc func(n *Node, index int) bool //find by value, add index for more flexible
type mapNodeFunc func(n *Node) interface{}
type MapFunc func(n *Node, index int) interface{} //by value


//find a Node depend on value
func (l *List) findOne(f LookupFunc) *Node {
	var i = 0
	for node := l.Front(); node != nil; node = node.Next() {
		if f(node,i) {
			return node
		}
		i++
	}
	return nil
}

//the same function like findOne, but doesn't provide index
func (l *List) lookup(f lookupFunc) *Node {

	for node := l.Front(); node != nil; node = node.Next() {
		if f(node) {
			return node
		}
	}
	return nil
}

//
func (l *List) lookupFromRoot(f lookupFunc) *Node {

	for node := l.root_(); node != nil; node = node.Next() {
		if f(node) {
			return node
		}
	}
	return nil
}

//return list of Node by LookupFunc
func (l *List) findAllValues(f LookupFunc) (nodes []interface{}){
	var i = 0
	for node := l.Front(); node != nil; node = node.Next() {
		if f(node,i) {
			nodes = append(nodes,node.Value)
		}
		i++
	}
	return
}

func (l *List) findAllNodes(f LookupFunc) (nodes []*Node){
	var i = 0
	for node := l.Front(); node != nil; node = node.Next() {
		if f(node,i) {
			nodes = append(nodes,node)
		}
		i++
	}
	return
}

func (l *List) Find(f LookupFunc) *Node {
	return l.findOne(f)
}

//return value
func (l *List) FindValue(f LookupFunc) interface{} {
	result := l.findOne(f)
	if result != nil {
		return result.Value
	}
	return nil
}


func (l *List) Filter(f LookupFunc) []*Node {
	return l.findAllNodes(f)
}


func (l *List) FilterValues(f LookupFunc) []interface{} {
	return l.findAllValues(f)
}

//func (l *List) FilterNodes(f LookupFunc) []*Node {
//	result := l.findAllOnValue(f, true)
//	if len(result) > 0 {
//		return result.([]*Node)
//
//	}
//	//return l.findAllByNode()
//}

func (l *List) map_ (f MapFunc, allowNull bool) (results []interface{}){
	var i = 0
	for node := l.Front(); node != nil; node = node.Next() {
		fResult := f(node, i)
		if fResult != nil || allowNull {
			results = append(results, fResult)
		}
		i++
	}
	return
}
//The Map() method creates a new array with the results of calling a provided function on every element in the calling array.
func (l *List) Map(f MapFunc, allowNull ...bool) []interface{}{
	var acceptNull = true
	if len(allowNull) > 0 {
		acceptNull = allowNull[0]
	}
	return l.map_(f,acceptNull)

}

func (l *List) Concat(values ...interface{}){
	l.appendMany(values)
}

//Append a node to the end of the list O(1)
func (l* List) Append(value interface{}){
	l.appendOne(value)
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
	if mark.list != l {
		return nil
	}
	//if l.head == mark {
	//	return l.prepend(mark)
	//}

	nodeBeforeMark := l.lookup(func(n *Node) bool {
		return n.next == mark
	})

	return l.insertValue(v, nodeBeforeMark)
}

//move node n after node at
func (l *List) move(n, at *Node) *Node {
	if n == at || at.next == n {
		return n
	}
	nodeBeforeN := l.lookup(func(node *Node) bool {
		return node.next == n
	})
	nodeBeforeN.next = n.next
	n.next = at.next
	at.next = n
	if l.tail == at {
		l.tail = n
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
	nodeBeforeMark := l.lookupFromRoot(func(n *Node) bool {
		return n.next == mark
	})
	l.move(n, nodeBeforeMark)
}

//need implement
func (l *List) MoveToBack(n *Node) {
	if n.list != l || l.tail == n {
		return
	}
	l.move(n, l.tail)
}

//need implement
func (l *List) MoveToFront(n *Node){
	if n.list != l || l.tail == n {
		return
	}
}

//need implement
func (l *List) PushBack(v interface{}) *Node {
	return l.appendOne(v)
}

//need implement
func (l *List) PushBackList(other *List){

}

//need implement
func (l *List) PushFront(v interface{}) *Node {
	return l.prepend(v)
}

//need implement
func (l *List) PushFrontList(other *List){

}

func (l *List) remove(n *Node) *Node {
	return nil
}

//need implement
func (l *List) Remove(n *Node) interface{}{
	//node := l.lookup(func(n *Node) bool {
	//	return n.next == n
	//})

	return nil
}

//need implement
func (l *List) Clear() {
	l.lazyInit()
}


//need implement with cache
func (l *List) NodeAt(index int) *Node {
	return nil
}

