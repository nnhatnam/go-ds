package BSTree

import (
	"github.com/nnhatnam/go-ds/binarytree"
)

type BSTNode struct {
	Left *BSTNode
	Right *BSTNode
	Value binarytree.ValueInterface
}

func NewBSTNode(v binarytree.ValueInterface) *BSTNode {
	return &BSTNode{
		Value: v,
	}
}

func (node *BSTNode) hasRightChild() bool {
	return node.Right != nil
}

func (node *BSTNode) hasLeftChild() bool {
	return node.Left != nil
}

func (node *BSTNode) hasNoChild() bool {
	return !node.hasLeftChild() && !node.hasRightChild()
}


type BSTTree struct {
	Root *BSTNode
	Size int
}

/*
Compute the "maxDepth" of a tree -- the number of nodes along
the longest path from the root node down to the farthest leaf node.
*/
func maxDepth(node *BSTNode) int {
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

func insert(node *BSTNode, v binarytree.ValueInterface) {
	if node == nil || v.Compare(node.Value) == 0 {
		return
	}

	switch v.Compare(node.Value){
	case 1:
		if node.Right == nil {
			node.Right.Value = v
		} else {
			insert(node.Right, v)
		}
	case -1:
		if node.Left == nil {
			node.Left.Value = v
		} else {
			insert(node.Left, v)
		}
	default:
		return
	}

}

func NewBSTTree() *BSTTree {
	return new(BSTTree)
}

//func (bst *BSTTree) Root() * Node {
//	return bst.Root
//}

//func (bst *BSTTree) size(n *Node ) int {
//	if n == nil {
//		return 0
//	}
//
//	return bst.size(n.Left) + 1 + bst.size(n.Right)
//}

//func (bst *BSTTree) Size() int {
//	return bst.Size
//}

func (bst *BSTTree) Insert(v binarytree.ValueInterface) *BSTTree {
	if bst.Root == nil {
		bst.Root = NewBSTNode(v)

	} else {
		insert(bst.Root, v)
	}
	bst.Size++
	return bst

}

func (bst *BSTTree) MaxDepth() int {
	return maxDepth(bst.Root)
}

func findMin(n *BSTNode) *BSTNode {
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
func (bst *BSTTree) Min() *BSTNode {
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

func findMax(n *BSTNode) *BSTNode {
	current := n
	if current == nil {
		return nil
	}
	for ; current.Right != nil;  {
		current = current.Right
	}
	return  current
}

func (bst *BSTTree) Max() *BSTNode {
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

func preOrderTraverse(n *BSTNode, f binarytree.TraverseFunc) {
	if n == nil {
		return
	}
	f(n.Value)
	preOrderTraverse(n.Left, f)
	preOrderTraverse(n.Right, f)

}

func (bst *BSTTree) PreOderTraverse(f binarytree.TraverseFunc) {
	if bst != nil {
		preOrderTraverse(bst.Root, f)
	}

}

func inOrderTraverse(n *BSTNode, f binarytree.TraverseFunc) {
	if n == nil {
		return
	}
	preOrderTraverse(n.Left, f)
	f(n.Value)
	preOrderTraverse(n.Right, f)
}

func (bst *BSTTree) InOderTraverse(f binarytree.TraverseFunc) {
	if bst != nil {
		inOrderTraverse(bst.Root, f)
	}
}

func postOrderTraverse(n *BSTNode, f binarytree.TraverseFunc) {
	if n == nil {
		return
	}
	preOrderTraverse(n.Left, f)
	f(n.Value)
	preOrderTraverse(n.Right, f)

}

func (bst *BSTTree) PostOderTraverse(f binarytree.TraverseFunc) {
	if bst != nil {
		postOrderTraverse(bst.Root, f)
	}
}

func search(n *BSTNode, v binarytree.ValueInterface) *BSTNode {
	if n == nil || n.Value.Compare(v) == 0 {
		return n
	}

	if n.Value.Compare(v) == -1 {
		return search(n.Left, v)
	}
	return search(n.Right, v)
}

func (bst *BSTTree) Search(v binarytree.ValueInterface) *BSTNode {
	if bst == nil {
		return nil
	}

	return search(bst.Root, v)
}

//Hibbard deletion algorithm
//https://algs4.cs.princeton.edu/32bst/
//optimize later
func remove(root *BSTNode, v binarytree.ValueInterface) binarytree.ValueInterface {
	if root == nil {
		return nil
	}
	parent := root
	curr := root

	//compareResult := compareFunc(curr.Key, key)
	for ; curr != nil && curr.Value.Compare(v) != 0 ;  {
		parent = curr
		if curr.Value.Compare(v) == 1 {
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
		curr.Value = successor.Value
		return remove(successor, successor.Value)
	} else {											//Case 3: node to be deleted has only one children
		var child *BSTNode
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

func (bst *BSTTree) Remove(v binarytree.ValueInterface) binarytree.ValueInterface {
	if bst == nil {
		return nil
	}

	removedNode := remove(bst.Root, v)
	if removedNode != nil {
		bst.Size--
	}
	return removedNode
}

