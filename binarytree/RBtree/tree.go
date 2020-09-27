package RBtree

type Node struct {
	Key int
	Red bool
	Left *Node
	Right *Node
	Parent *Node
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

func (tree *RBTree) adjustTree(node *Node) {
	for ;node.Parent.Red == true && node != tree.Root; {

	}
}

func (tree *RBTree) Insert(key int) {

	new_node := &Node{
		Key:    key,
		Red:    false,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
	// Nothing in tree
	if tree.Root == nil {
		tree.Root = new_node
		return
	}

	//Search to find the node's correct place
	current := tree.Root
	var potentialParent *Node
	for ;current != nil; {
		potentialParent = current
		if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	//Assigns parents and siblings to the new node
	new_node.Parent = potentialParent

}