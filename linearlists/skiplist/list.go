package skiplist

import (
	"github.com/nnhatnam/go-ds/types"
	"math"
	"math/bits"
	"math/rand"
	"time"
)

const (
	DefaultMaxLevel int = 32
	DefaultProb float64 = 1 / math.E
)

type nodes struct {
	next []*Element
}

type Element struct {

	value types.Comparator

	list *SkipList

	next []*Element
	prev []*Element
}

//func (e *Element) Key() float64 {
//	return e.key
//}

func (e *Element) Value() interface{} {
	return e.value
}

func (e *Element) Next() *Element {
	return e.next[0]
}

//need to re-order
type SkipList struct {

	root *Element

	maxLevel int
	prob float64

	rand *rand.Rand

	length int

}

func NewWithConfig(maxLevel int, prob float64) *SkipList {

	if maxLevel < 1 || maxLevel > 64 {
		panic("maxLevel for a SkipList must be a positive integer <= 64")
	}

	source := rand.NewSource(time.Now().UnixNano())

	return &SkipList{
		length:   0,
		prob: prob,
		maxLevel: maxLevel,

		rand : rand.New(source),
	}
}

func New() *SkipList {
	return NewWithConfig(DefaultMaxLevel, DefaultProb)
}

func (list *SkipList) generateLevel() int {

	if list.maxLevel <= 1 {
		return 1
	}
	level := 0

	var x uint64 = rand.Uint64() & ((1 << uint(list.maxLevel-1)) - 1)
	zeroes := bits.TrailingZeros64(x)
	if zeroes <= list.maxLevel {
		level = zeroes
	}
	return level
}

func (list *SkipList) Insert(value types.Comparator) {

}
