package rope

type FlatRope struct {
	buf []byte
}

func NewFlatRopeFromString(s string) *FlatRope {
	return &FlatRope{buf: []byte(s)}
}

func (r *FlatRope) Len() int {
	return len(r.buf)
}

func (r *FlatRope) Depth() byte {
	return 0
}

func (r *FlatRope) ToString() string {
	return ""
}

func (r *FlatRope) CharAt(index int) string {
	return string(r.buf[index])
}