package bst

type Node struct {
	Left *Node
	Right *Node
	Data int
}

func NewNode(data int) *Node {
	return &Node{Data:data,}
}

func (n *Node) Insert(data int) {
	if n == nil {
		return
	}

	if data <= n.Data {
		if n.Left == nil {
			n.addLeftChild(data)
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.addRightChild(data)
		} else {
			n.Right.Insert(data)
		}
	}
}


func (n *Node) addLeftChild(key int) {
	if n == nil {
		return
	}

	n.Left = &Node{Data: key,}
}

func (n *Node) addRightChild(key int) {
	if n == nil {
		return
	}

	n.Right = &Node{Data: key,}
}

func (n *Node) hasRightChild() bool {
	return n.Right != nil
}

func (n *Node) hasLeftChild() bool {
	return n.Left != nil
}



type BST struct {
	Root *Node
	size int
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

func (bst *BST) Size() int {
	return bst.size
}

func (bst *BST) Insert(data int) * BST {
	if bst.Root == nil {
		bst.Root = NewNode(data)

	} else {
		bst.Root.Insert(data)
	}
	bst.size++
	return bst

}

func (bst *BST) MaxDepth() int {
	return maxDepth(bst.Root)
}

func (bst *BST) MinValue() int {
	current := bst.Root
	if current == nil {
		return 0
	}
	for ; current.Left != nil;  {
		current = current.Left
	}
	return current.Data
}

func (bst *BST) MaxValue() int {
	current := bst.Root
	if current == nil {
		return 0
	}
	for ; current.Right != nil;  {
		current = current.Right
	}
	return current.Data
}

func (bst *BST) PrintPreOder() {
	
}

func (bst *BST) PrintInOder() {

}

func (bst *BST) PrintPostOder() {

}