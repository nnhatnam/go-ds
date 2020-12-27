package rope

//http://ahmadsoft.org/source/xref/ropes-1.2.5/src/org/ahmadsoft/ropes/impl/

//const small = 128
//
//type Rope struct {
//	root tree
//}
//
//
//type node struct {
//	left, right tree
//	nbytes      int
//}
//
//func (n *node) concat(right tree) tree {
//	var left tree
//	s1, ok1 := n.right.(leaf)
//	s2, ok2 := n.left.(leaf)
//	if ok1 && ok2 && len(string(s1)) + len(string(s2)) <= small {
//		right = leaf(string(s1) + string(s2))
//		left = n.left
//	} else {
//		left = n
//	}
//
//	return &node{left, right, left.length() + right.length()}
//}
//
//func (n *node) index(i int) byte {
//	if i > n.nbytes {
//		panic(fmt.Sprintf("%d out of bounds for length-%d rope", i, n.length))
//	}
//	leftlength := n.left.length()
//	if i < leftlength {
//		return n.left.index(i)
//	}
//	return n.right.index(i - leftlength)
//}
//
//func (n *node) length() int {
//	return n.nbytes
//}
//
//func (n *node) slice(i, j int) tree {
//	leftlen := n.left.length()
//	switch {
//	case j <= leftlen:
//		return n.left.slice(i, j)
//	case i >= leftlen:
//		return n.right.slice(i-leftlen, j-leftlen)
//	}
//
//	var left, right tree
//	if i == 0 {
//		left = n.left
//	} else {
//		left = n.left.slice(i, leftlen)
//	}
//	if j == leftlen+n.right.length() {
//		right = n.right
//	} else {
//		right = n.right.slice(0, j-leftlen)
//	}
//	return left.concat(right)
//}
//
//type leaf string
//
//func (s leaf) concat(t tree) tree {
//	if s2, ok := t.(leaf); ok && len(string(s)) + len(string(s2)) <= small {
//		return leaf(string(s) + string(s2))
//	}
//
//	return &node{s, t, len(s) + t.length()}
//}
//
//func (s leaf) index(i int) byte {
//	return string(s)[i]
//}
//
//func (s leaf) length() int {
//	return len(string(s))
//}
//
//func (s leaf) slice(i, j int) tree {
//	return leaf(string(s)[i:j])
//}
//
//type tree interface {
//	concat(tree) tree
//	index(int) byte
//	length() int
//	slice(int, int) tree
//
//}