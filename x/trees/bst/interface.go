package bst

import "github.com/nnhatnam/go-ds/x/trees"

//import "github.com/nnhatnam/go-ds/trees"

type INode interface {
	Value() interface{}
	SetValue(v interface{})
	Left() INode
	SetLeft(v interface{}) INode
	Right() INode
	SetRight(v interface{}) INode
}

var _ INode = (*Node)(nil)

type Node struct {
	left *Node
	right *Node
	value interface{}
}

func NewNode(v interface{}) *Node {
	return &Node{value: v,}
}

func (n *Node) Value() interface{} {
	return nil
}

func (n *Node) SetValue(v interface{}) {
	if n != nil {
		n.value = v
	}
}

func (n *Node) Left() INode {
	return n.left
}

func (n *Node) SetLeft(v interface{}) INode {
	if n != nil {
		new_node := NewNode(v)
		n.left = new_node
		return new_node
	}
	return nil
}

func (n *Node) Right() INode {
	return n.right
}

func (n *Node) SetRight(v interface{}) INode{
	if n != nil {
		new_node := NewNode(v)
		n.right = NewNode(v)
		return new_node
	}
	return nil
}

//calculate max depth from a node
func MaxDepth(node INode) int {
	if node == nil {
		return 0
	}
	lDepth := MaxDepth(node.Left())
	rDepth := MaxDepth(node.Right())

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func FindMin(n INode) INode {
	current := n
	if current == nil {
		return nil
	}
	for ; current.Left() != nil;  {
		current = current.Left()
	}
	return  current
}

func FindMax(n INode) INode {
	current := n
	if current == nil {
		return nil
	}
	for ; current.Right() != nil;  {
		current = current.Right()
	}
	return  current
}

func PreOrderTraverse(n INode, f trees.TraverseFunc) {
	if n == nil {
		return
	}

	f(n.Value())
	PreOrderTraverse(n.Left(), f)
	PreOrderTraverse(n.Right(), f)

}

func InOrderTraverse(n INode, f trees.TraverseFunc) {
	if n == nil {
		return
	}
	InOrderTraverse(n.Left(), f)
	f(n.Value())
	InOrderTraverse(n.Right(), f)
}

func PostOrderTraverse(n INode, f trees.TraverseFunc) {
	if n == nil {
		return
	}
	PostOrderTraverse(n.Left(), f)
	f(n.Value())
	PostOrderTraverse(n.Right(), f)

}

func Insert(node INode, v interface{}, compare trees.Comparator) INode {

	if node == nil || compare(v, node.Value()) == 0 {
		return node
	}

	//if node.Value > v then add v to left child
	if compare(node.Value() , v ) == 1 {
		if node.Left() == nil {
			return node.SetLeft(v)
		} else {
			return Insert(node.Left(), v, compare)
		}
	} else { 		//if node.Value < v then add v to right child
		if node.Right() == nil {
			return node.SetRight(v)
		} else {
			return Insert(node.Right(), v, compare)
		}
	}
}