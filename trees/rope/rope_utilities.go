package rope

type RopeUtilities struct {

}

func (r *RopeUtilities) concatenate(left Rope, right Rope) Rope {
	if left.Len() == 0 {
		return right
	}

	if right.Len() == 0 {
		return left
	}
	//TODO: Need to detect overflow here
	combineLength := 17
	if left.Len() + right.Len() < combineLength {
		return NewFlatRopeFromString(left.ToString() + right.ToString())
	}
	cLeft, ok1 := left.(*ConcatenationRope)
	cRight, ok2 := right.(*ConcatenationRope)

	if !ok1 && ok2 {
		if left.Len() + cRight.Left().Len() < combineLength {
			return r.autoRebalance(NewConcatenationRope( NewFlatRopeFromString(left.ToString() + cRight.Left().ToString()), cRight.Right() ))
		}
	}

	if !ok2 && ok1 {
		if right.Len() + cLeft.Left().Len() < combineLength {
			return r.autoRebalance(NewConcatenationRope( cLeft.Left(),  NewFlatRopeFromString(cLeft.Right().ToString() + right.ToString()) ))
		}
	}

	return r.autoRebalance(NewConcatenationRope(left, right))
}

func (r *RopeUtilities) depth(rope Rope) byte {
	return rope.Depth()
}

func (r *RopeUtilities) autoRebalance(rope Rope) Rope {
	if rope.Depth() > MAX_ROPE_DEPTH {
		return r.rebalance(rope)
	}
	return rope
}

type traverseFunc func(r Rope)

func (r *RopeUtilities) inOrderTraverse(n Rope, f traverseFunc) {
	if n == nil {
		return
	}

	switch n.(type) {
	case *FlatRope:
		f(n)
	case *ConcatenationRope:
		r.inOrderTraverse(n.(*ConcatenationRope).left, f)
		r.inOrderTraverse(n.(*ConcatenationRope).right, f)
	}
}

func (r *RopeUtilities) rebalance(rope Rope) Rope {
	var leafNodes = []Rope{}
	r.inOrderTraverse(rope, func(r1 Rope) {
		leafNodes = append(leafNodes, r1)
	})

	return r.merge(leafNodes, 0, len(leafNodes))
}

func (r *RopeUtilities) merge(leafNodes []Rope, start, end int) Rope {
	distance := end - start
	switch distance {
	case 1:
		return leafNodes[start]
	case 2:
		return NewConcatenationRope(leafNodes[start], leafNodes[start + 1])
	default:
		mid := start + ( distance >> 1 )
		return NewConcatenationRope(r.merge(leafNodes, start, mid), r.merge(leafNodes, mid, end))

	}
}