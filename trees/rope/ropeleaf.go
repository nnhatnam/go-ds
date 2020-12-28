package rope

type ropeLeaf struct {
	data string
}

func newRopeLeaf(s string) *ropeLeaf {
	return &ropeLeaf{data: s}
}

func (leaf *ropeLeaf) concat(b baserope) {

}

func (leaf *ropeLeaf) len() int {
	return len(leaf.data)
}

func (leaf *ropeLeaf) depth() byte {
	return 0
}

func (leaf *ropeLeaf) toString() string {
	return leaf.data
}

func (leaf *ropeLeaf) index(i int) byte {
	return leaf.data[i]
}

