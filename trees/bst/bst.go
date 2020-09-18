package bst

import "github.com/nnhatnam/go-ds/trees"

type Node struct {
	Left *Node
	Right *Node
	Key interface{}
	Value interface{}
}

type TraverseFunc func(interface{}, interface{})

func NewNode(k, v interface{}) *Node {
	return &Node{Key: k, Value:v,}
}


//func (node *Node) insert(k, v interface{}, compare trees.CompareFunc) {
//	if node == nil || compare(k, node.Key) == 0 {
//		return
//	}
//
//	switch compare(k, node.Key){
//	case 1:
//		if node.Right == nil {
//			node.addRightChild(k, v)
//		} else {
//			insert(node.Right, k, v, compare)
//		}
//	case -1:
//		if node.Left == nil {
//			node.addLeftChild(k, v )
//		} else {
//			insert(node.Left,k, v, compare)
//		}
//	default:
//		return
//	}
//
//}

func (node *Node) addLeftChild(k, v interface{}) {
	if node == nil {
		return
	}

	node.Left = &Node{Key: k, Value: v,}
}

func (node *Node) addRightChild(k, v interface{}) {
	if node == nil {
		return
	}

	node.Right = &Node{Key: k, Value: v,}
}

func (node *Node) hasRightChild() bool {
	return node.Right != nil
}

func (node *Node) hasLeftChild() bool {
	return node.Left != nil
}



type BST struct {
	Root *Node
	Size int

	compareFunc trees.CompareFunc
}

/*
Compute the "maxDepth" of a tree -- the number of nodes along
the longest path from the root node down to the farthest leaf node.
*/
func maxDepth(node *Node) int {
	if node == nil {
		return 0
	}
	lDepth := maxDepth(node.Left)
	rDepth := maxDepth(node.Right)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func insert(node *Node, k, v interface{}, compare trees.CompareFunc) {
	if node == nil || compare(k, node.Key) == 0 {
		return
	}

	switch compare(k, node.Key){
	case 1:
		if node.Right == nil {
			node.addRightChild(k, v)
		} else {
			insert(node.Right, k, v, compare)
		}
	case -1:
		if node.Left == nil {
			node.addLeftChild(k, v )
		} else {
			insert(node.Left,k, v, compare)
		}
	default:
		return
	}

}

func NewBST() *BST {
	return new(BST)
}

//func (bst *BST) Root() * Node {
//	return bst.Root
//}

//func (bst *BST) size(n *Node ) int {
//	if n == nil {
//		return 0
//	}
//
//	return bst.size(n.Left) + 1 + bst.size(n.Right)
//}

//func (bst *BST) Size() int {
//	return bst.Size
//}

func (bst *BST) Insert(k , v interface{}) * BST {
	if bst.Root == nil {
		bst.Root = NewNode(k, v)

	} else {
		insert(bst.Root, k, v, bst.compareFunc)
	}
	bst.Size++
	return bst

}

func (bst *BST) MaxDepth() int {
	return maxDepth(bst.Root)
}

func (bst *BST) MinKey() interface{} {
	current := bst.Root
	if current == nil {
		return 0
	}
	for ; current.Left != nil;  {
		current = current.Left
	}
	return current.Key
}

func (bst *BST) MaxKey() interface{} {
	current := bst.Root
	if current == nil {
		return 0
	}
	for ; current.Right != nil;  {
		current = current.Right
	}
	return current.Key
}

func preOrderTraverse(n *Node, f TraverseFunc) {

	f(n.Key, n.Value)
	preOrderTraverse(n.Left, f)
	preOrderTraverse(n.Right, f)
}

func (bst *BST) PrintPreOder() {

}

func inOrderTraverse(n *Node, f TraverseFunc) {
	preOrderTraverse(n.Left, f)
	f(n.Key, n.Value)
	preOrderTraverse(n.Right, f)
}

func (bst *BST) PrintInOder() {

}

func postOrderTraverse(n *Node, f TraverseFunc) {
	
	preOrderTraverse(n.Left, f)
	f(n.Key, n.Value)
	preOrderTraverse(n.Right, f)
}

func (bst *BST) PrintPostOder() {

}