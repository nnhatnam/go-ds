package bst

type Node struct {
	Left  *Node
	Right *Node
	Value Interface
}

func NewNode(v Interface) *Node {
	return &Node{
		Value: v,
	}
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


type Tree struct {
	//root node is dummy node, the actual root is root.Left
	root *Node
	//caching size of tree
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


func NewTree() *Tree {
	return new(Tree).Init()
}

func (bst *Tree) Init() *Tree{
	bst.root = NewNode(nil)
	bst.size = 0
	return bst
}

func (bst *Tree) lazyInit() {
	if bst.root == nil {
		bst.Init()
	}
}

func (bst *Tree) Root() *Node {
	return bst.root.Left
}

//func (bst *Tree) size(n *Node ) int {
//	if n == nil {
//		return 0
//	}
//
//	return bst.size(n.Left) + 1 + bst.size(n.Right)
//}

//func (bst *Tree) Size() int {
//	return bst.Size
//}

//Insert value v into the tree, v must be not nil
func (bst *Tree) Insert(v Interface) *Tree {
	bst.lazyInit()

	if v == nil {
		return nil
	}

	bst.size++
	if bst.root.Left == nil {
		bst.root.Left = NewNode(v)
		return bst
	}

	insert(bst.root.Left, v)


	return bst

}

//insert a node node into a tree. if there is node that has value equal to v. v will be added to the right child
//v must be not nil
func insert(node *Node, v Interface) {
	//if node is nil, nothing happens
	if node == nil {
		return
	}

	switch v.Compare(node.Value){
	//if v >= node.Value then go right
	case 0, 1:
		if node.Right == nil {
			node.Right = NewNode(v)
		} else {
			insert(node.Right, v)
		}
	//	if v < node.Value, then go left
	case -1:
		if node.Left == nil {
			node.Left = NewNode(v)
		} else {
			insert(node.Left, v)
		}
	default:
		return
	}
}


func (bst *Tree) MaxDepth() int {
	return maxDepth(bst.Root())
}

//find min from a node, return the node and its parent. If n is the min node, parent will be nil
//n must be not nil
func findMin(n *Node) (parent *Node, node *Node) {
	parent = nil
	node = n
	for ; node.Left != nil; node = node.Left {
		parent = node
	}
	return
}

//return the value with min key stored in the tree
func (bst *Tree) Min() *Node {
	if bst == nil {
		return nil
	}
	_, n := findMin(bst.Root())
	return n
}

//find max from a node, return the node and its parent. If n is the max node, parent will be nil
//n must be not nil
func findMax(n *Node)  (parent *Node, node *Node)  {
	parent = nil
	node = n
	for ; node.Right != nil; node = node.Right {
		parent = node
	}
	return
}

func (bst *Tree) Max() *Node {
	if bst == nil {
		return nil
	}
	_, n := findMax(bst.Root())
	return n
}

func preOrderTraverse(n *Node, f TraverseFunc) {
	if n == nil {
		return
	}
	f(n.Value)
	preOrderTraverse(n.Left, f)
	preOrderTraverse(n.Right, f)

}

func (bst *Tree) PreOderTraverse(f TraverseFunc) {
	if bst != nil {
		preOrderTraverse(bst.Root(), f)
	}

}

func inOrderTraverse(n *Node, f TraverseFunc) {
	if n == nil {
		return
	}
	inOrderTraverse(n.Left, f)
	f(n.Value)
	inOrderTraverse(n.Right, f)
}

func (bst *Tree) InOderTraverse(f TraverseFunc) {
	if bst != nil {
		inOrderTraverse(bst.Root(), f)
	}
}

func postOrderTraverse(n *Node, f TraverseFunc) {
	if n == nil {
		return
	}
	preOrderTraverse(n.Left, f)
	f(n.Value)
	preOrderTraverse(n.Right, f)

}

func (bst *Tree) PostOderTraverse(f TraverseFunc) {
	if bst != nil {
		postOrderTraverse(bst.Root(), f)
	}
}

//v, n must not be nil
func search(n *Node, v Interface) (parent, node *Node) {
	for ; n != nil && v.Compare(n.Value) != 0 ; {
		parent = n
		if v.Compare(n.Value) == -1 {
			n = n.Left
		} else {
			n = n.Right
		}
	}

	return parent, n
}

func (bst *Tree) Search(v Interface) *Node {
	if bst == nil {
		return nil
	}
	_ , n := search(bst.root, v)
	return n
}

func (bst *Tree) Remove(v Interface) *Node {
	if bst == nil || v == nil {
		return nil
	}
	if v.Compare(bst.Root().Value) == 0 {
		bst.size--
		return bst.removeHibbard(bst.root, bst.Root())
	}
	bst.size--
	return bst.removeHibbard(search(bst.Root(), v))
	//return removedNode
}



//Hibbard deletion algorithm
//https://algs4.cs.princeton.edu/32bst/
//remove deletedNode from its parent using Hibbard delete algorithm. if deletedNode is the root. parent would be the dummy node
//remove can't remove parent node due to Go doesn't support pass by reference
//return the new node that take place of deleted node
//remove don't handle size of the tree
func (bst *Tree) removeHibbard(parent *Node, deletedNode *Node) *Node {
	if parent == nil || deletedNode == nil {
		return nil
	}
	//value := deletedNode.Value
	if deletedNode.hasNoChild() {
		//Case 1: Node to be deleted has zero children
		//What we need to do is set node to nil
		//Because we use dummy root node mechanism, current != dummy root
		if parent.Left == deletedNode {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return deletedNode
	} else if deletedNode.Left != nil && deletedNode.Right != nil {
		//Case 2: Node to be deleted has two children
		//Find minimum element in the right subtree of the node to be removed
		s_parent , successor := findMin(deletedNode.Right)
		deletedNode.Value = successor.Value
		bst.removeHibbard(s_parent, successor)
		return deletedNode
	} else {
		//Case 3: node to be deleted has only one children
		var child *Node
		if deletedNode.Left != nil {
			child = deletedNode.Left
			deletedNode.Left = nil
		} else {
			child = deletedNode.Right
			deletedNode.Right = nil
		}

		if parent.Left == deletedNode {
			parent.Left = child
		} else {
			parent.Right = child
		}

		return child
	}
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

