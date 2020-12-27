package rope

import "math"



type ConcatenationRope struct {
	left Rope
	right Rope
	depth byte
	len int
}

func NewConcatenationRope(left, right Rope) *ConcatenationRope {
	return &ConcatenationRope{
		left:  left,
		right: right,
		depth: byte(math.Max(float64(left.Depth()), float64(right.Depth()))) + 1,
		len:   left.Len() + right.Len(),
	}
}


func (r *ConcatenationRope) Len() int {
	return r.len
}

func (r *ConcatenationRope) Depth() byte {
	return r.depth
}

func (r *ConcatenationRope) ToString() string {
	return ""
}

func (r *ConcatenationRope) CharAt(index int) string {
	if index > r.len {
		panic(ERR_INDEX_OUT_OF_BOUND)
	}

	if index < r.len {
		return r.left.CharAt(index)
	}
	return r.right.CharAt(index)
}

func (r *ConcatenationRope) Left() Rope {
	return r.left
}

func (r *ConcatenationRope) Right() Rope {
	return r.right
}

