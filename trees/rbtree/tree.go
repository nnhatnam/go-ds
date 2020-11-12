package rbtree

type Node struct {
	Key int
	Black bool
	Left *Node
	Right *Node
	Parent *Node
}

//Create new RED node
func NewNode(key int) *Node {
	return &Node{
		Key: key,
	}
}

func (n *Node) isBlack() bool {
	if n == nil || n.Black {
		return true
	}
	return false 
}

func (n *Node) isRed() bool {
	if n == nil || !n.Black {
		return true
	}
	return false
}

func (n *Node) flipColor() {
	if n != nil {
		n.Black = !n.Black
	}
}

func (n *Node) setBlackColor() {
	if n != nil {
		n.Black = true
	}
}

func (n *Node) setRedColor() {
	if n != nil {
		n.Black = false
	}
}

type RBTree struct {
	Root *Node
	Size int
}

func NewRBTree() *RBTree {
	return new(RBTree)
}

func search(node *Node, key int) *Node {
	current := node
	for ;current != nil ; {
		if current.Key == key {
			return current
		}
		if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return current
}

func (tree *RBTree) Search(key int) *Node {
	return search(tree.Root, key)
}

func (tree *RBTree) leftRotate(node *Node) {

	//if node == tree.Root {
	//	rChild := node.Right
	//	if rChild != nil {
	//		tree.Root = rChild
	//		rChild.Parent = nil
	//
	//	}
	//}
	rChild := node.Right


	//make sure rChild.Left == nil
	if rChild.Left != nil {
		node.Right = rChild.Left
		rChild.Left = nil
	}

	if node == tree.Root {    //parent == nil
		tree.Root = rChild
		tree.Root.Parent = nil
		tree.Root.Left = node
		node.Parent = tree.Root
	} else {
		parent := node.Parent
		parent.Right = rChild
		rChild.Parent = parent
		parent.Left = node
		node.Parent = parent
	}
}

func (tree *RBTree) rightRotate(node *Node) {
	lChild := node.Left

	//make sure lChild.Right == nil
	if lChild.Right != nil {
		node.Left = lChild.Right
		lChild.Right = nil
	}

	if node == tree.Root {    //parent == nil
		tree.Root = lChild
		tree.Root.Parent = nil
		tree.Root.Right = node
		node.Parent = tree.Root
	} else {
		parent := node.Parent
		parent.Left = lChild
		lChild.Parent = parent
		parent.Left = node
		node.Parent = parent
	}
}



func (tree *RBTree) Insert(key int) * Node{
	//Naming conversion:
	//	1. PNode = Parent Node
	//	2. UNode = Uncle Node / Sibling of Parent Node
	//	3. GNode = Grandparent Node
	//	4. NNode = New Node (Insertion Node)

	//red node
	NNode := NewNode(key)
	// If Tree is empty, we make NNode the root of the tree and color it black.
	if tree.Root == nil {
		tree.Root = NNode
		tree.Root.setBlackColor()
		return tree.Root
	}

	//Search to find the node's correct place
	current := tree.Root
	var PNode = tree.Root // PNode = tree.Root more semantic compare to PNode = current
	for ;current != nil; {
		PNode = current
		if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	//Assigns parents and siblings to the new node
	NNode.Parent = PNode

	if NNode.Key < PNode.Key {
		PNode.Left = NNode
	} else {
		PNode.Right = NNode
	}

	return tree.adjustTree(NNode)
}

func (tree *RBTree) adjustTree(NNode *Node) *Node{
	PNode := NNode.Parent
	GNode := PNode.Parent
	UNode := GNode.Left
	if UNode == PNode {
		UNode = GNode.Right
	}
	switch {
	//Case 1: P is black
	//if P is black, it can not violate any of the properties. Therefore, in this case, we do not need to do anything.
	case PNode.isBlack():

	//	Case 2: P is red
	//If the parent node PNode is red, this violates the property 4. P and N are now both red.
	//The grandparent node GNode must be black node because the tree before insertion must be a valid red-black tree.
	//To resolve this case, we need to check whether UNode is red or black.
	case PNode.isRed():
		//Case 2.1: Uncle Node is red
		//In this case, we flip the color of nodes P, U, and G.
		//That means, P becomes black, U becomes black and, G becomes red.
		//If G is the root. We don't flip color
		GNode := PNode.Parent
		if UNode.isRed() {

			if GNode != tree.Root {
				GNode.flipColor()
			}
			UNode.flipColor()
			PNode.flipColor()
			NNode.flipColor()
		} else {	//UNode.isBlack()
			//	Case 2.2: P is red and U is black (or NULL)
			//	This is more complicated than case 2.1.
			//	If the uncle node U is black, we need single or double tree rotations depending upon whether N is a left or right child of P.

			//	Case 2.2.1: P is right child of G and N is right child of P.
			//	We first perform the left-rotation at G that makes G the new sibling S of N. Next, we change the color of S to red and P to black.
			if GNode.Right == PNode && PNode.Right == NNode {
				tree.leftRotate(PNode)
				GNode.flipColor()
				PNode.flipColor()

			//	Case 2.2.2: P is right child of G and K is left child of P.
			//	In this case, we first do the right-rotation at P.
			//	This reduces it to the case 2.2.1. We next use the rules given in case 2.2.1 to fix the tree.
			} else if GNode.Right == PNode && PNode.Left == NNode {
				tree.rightRotate(PNode)
				tree.leftRotate(NNode)
				GNode.flipColor()
				NNode.flipColor()

			//	Case 2.2.3: P is left child of G and K is left child of P.
			//	This is the mirror of case 2.2.1
			} else if GNode.Left == PNode && PNode.Right == NNode {
				tree.rightRotate(PNode)
				GNode.flipColor()
				PNode.flipColor()

			//	GNode.Left == PNode && PNode.Left == NNode
			//	Case 2.2.4: P is left child of G and K is left child of P.
			//	This is the mirror of case 2.2.2
			} else {
				tree.leftRotate(PNode)
				tree.rightRotate(NNode)
				GNode.flipColor()
				NNode.flipColor()
			}
		}
	}
	return NNode
}

//func (tree *RBTree) Delete(key int) *Node {
//	deletedNode := search(tree.Root, key)
//
//}