package rope

func FromString(s string) Rope {
	return NewFlatRopeFromString(s)
}

type Rope interface {
	Len() int
	Depth() byte
	ToString() string
	CharAt(index int) string
}
