package rope

type Rope interface {
	Len() int
	Depth() byte
	ToString() string
	CharAt(index int) string
}
