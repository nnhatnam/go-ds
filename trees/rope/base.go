package rope

type rope interface {
	concat(rope) rope
	index(int) byte
	len() int
	slice(int, int) rope
	//writeTo(w io.Writer) (int64, error)
}
