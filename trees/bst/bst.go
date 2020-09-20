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

func (node *Node) hasNoChild() bool {
	return !node.hasLeftChild() && !node.hasRightChild()
}



type BST struct {
	Root *Node
	Size int

	CompareFunc trees.CompareFunc
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
		insert(bst.Root, k, v, bst.CompareFunc)
	}
	bst.Size++
	return bst

}

func (bst *BST) MaxDepth() int {
	return maxDepth(bst.Root)
}

func findMin(n *Node) *Node {
	current := n
	if current == nil {
		return nil
	}
	for ; current.Left != nil;  {
		current = current.Left
	}
	return  current
}

//return the value with min key stored in the tree
func (bst *BST) Min() *Node {
	if bst == nil {
		return nil
	}
	return findMin(bst.Root)

	//current := bst.Root
	//if current == nil {
	//	return 0
	//}
	//for ; current.Left != nil;  {
	//	current = current.Left
	//}

}

func findMax(n *Node) *Node {
	current := n
	if current == nil {
		return nil
	}
	for ; current.Right != nil;  {
		current = current.Right
	}
	return  current
}

func (bst *BST) Max() *Node {
	if bst == nil {
		return nil
	}
	return findMax(bst.Root)
	//current := bst.Root
	//if current == nil {
	//	return 0
	//}
	//for ; current.Right != nil;  {
	//	current = current.Right
	//}
	//return current.Key
}

func preOrderTraverse(n *Node, f TraverseFunc) {
	if n == nil {
		return
	}
	f(n.Key, n.Value)
	preOrderTraverse(n.Left, f)
	preOrderTraverse(n.Right, f)

}

func (bst *BST) PreOderTraverse(f TraverseFunc) {
	if bst != nil {
		preOrderTraverse(bst.Root, f)
	}

}

func inOrderTraverse(n *Node, f TraverseFunc) {
	if n == nil {
		return
	}
	preOrderTraverse(n.Left, f)
	f(n.Key, n.Value)
	preOrderTraverse(n.Right, f)
}

func (bst *BST) InOderTraverse(f TraverseFunc) {
	if bst != nil {
		inOrderTraverse(bst.Root, f)
	}
}

func postOrderTraverse(n *Node, f TraverseFunc) {
	if n == nil {
		return
	}
	preOrderTraverse(n.Left, f)
	f(n.Key, n.Value)
	preOrderTraverse(n.Right, f)

}

func (bst *BST) PostOderTraverse(f TraverseFunc) {
	if bst != nil {
		postOrderTraverse(bst.Root, f)
	}
}

func search(n *Node, key interface{}, compareFunc trees.CompareFunc) *Node{
	if n == nil || compareFunc(n.Key, key) == 0 {
		return n
	}

	if compareFunc(n.Key, key) == -1 {
		return search(n.Left, key, compareFunc)
	}
	return search(n.Right, key, compareFunc)
}

func (bst *BST) Search(key interface{}) *Node {
	if bst == nil {
		return nil
	}

	return search(bst.Root, key, bst.CompareFunc)
}

//Hibbard deletion algorithm
//https://algs4.cs.princeton.edu/32bst/
//optimize later
func remove(root *Node, key interface{} , compareFunc trees.CompareFunc) interface{} {
	if root == nil {
		return nil
	}
	parent := root
	curr := root

	//compareResult := compareFunc(curr.Key, key)
	for ; curr != nil && compareFunc(curr.Key, key) != 0 ;  {
		parent = curr
		if compareFunc(curr.Key, key) == 1 {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	if curr == nil {
		return nil
	}

	value := curr.Value
	if curr.hasNoChild() { 				//Case 1: Node to be deleted has no subtree

		if curr != root {
			if parent.Left == curr {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		} else {
			root = nil
		}

	} else if curr.Left != nil && curr.Right != nil {	//Case 2: Node to be deleted has two children
		successor := findMin(curr.Right)
		curr.Key, curr.Value = successor.Key, successor.Value
		return remove(successor, successor.Key, compareFunc)
	} else {											//Case 3: node to be deleted has only one children
		var child *Node
		if curr.Left != nil {
			child = curr.Left
		} else {
			child = curr.Right
		}

		if curr != root {
			if parent.Left == curr {
				parent.Left = child
			} else {
				parent.Right = child
			}
		} else {
			root = child
		}
	}
	return value
	//Case 2: Node to be deleted only has one children


	//node need to be deleted
	//deleteNode := search(root, key, compareFunc)
	//if deleteNode == nil {
	//	return nil
	//}

	//Case 1: No has no subtree
	//if deleteNode.hasNoChild() {
	//	deleteNode = nil
	//	return nil
	//}

}

func (bst *BST) Remove(key interface{}) interface{}{
	if bst == nil {
		return nil
	}
	return remove(bst.Root, key, bst.CompareFunc)
}

