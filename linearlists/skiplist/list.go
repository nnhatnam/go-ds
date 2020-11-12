package skiplist

import (
	"github.com/nnhatnam/hydrogen"
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

	key hydrogen.Comparator
	value interface{}

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
	rootElement := Element{
		key:   nil,
		value: nil,
		next:  make([]*Element, maxLevel),
	}
	return &SkipList{
		root : &rootElement,
		length:   0,
		prob: prob,
		maxLevel: maxLevel,

		rand : rand.New(source),
	}
}

func New() *SkipList {
	return NewWithConfig(DefaultMaxLevel, DefaultProb)
}

func (l *SkipList) pickHeight() int {

	if l.maxLevel <= 1 {
		return 1
	}
	level := 0

	var x uint64 = rand.Uint64() & ((1 << uint(l.maxLevel-1)) - 1)
	zeroes := bits.TrailingZeros64(x)
	if zeroes <= l.maxLevel {
		level = zeroes
	}
	return level
}

func (l *SkipList) Insert(key hydrogen.Comparator, value interface{}) {

	height := l.pickHeight()
	prevCache := make([]*Element, height - 1)
	cur := l.root
	l.length++
	for r := l.maxLevel - 1; r >= 0; r-- {

		for cur.next[r] != nil && cur.next[r].key.Cmp(key) < 0 {
			cur = cur.next[r]
		}

		if r <= height {
			prevCache[r] = cur
		}
	}

	if cur.key.Cmp(key) == 0 {
		cur.value = value
	}

	el := &Element{
		key:   key,
		value: value,

		next:  make([]*Element, height),
		prev:  make([]*Element, height),
	}

	for row, pointer := range prevCache {
		el.next[row] = pointer.next[row]
		el.prev[row] = pointer
		pointer.next[row].prev[row] = el
		pointer.next[row] = el
	}

}

func (l *SkipList) Size() int {
	return l.length
}

func (l *SkipList) IsEmpty() bool {
	return l.length == 0
}

func (l *SkipList) Find(key hydrogen.Comparator) interface{} {
	cur := l.root
	for r := l.maxLevel - 1; r >= 0; r-- {

		for cur.next[r] != nil && cur.next[r].key.Cmp(key) < 0 {
			cur = cur.next[r]
		}

	}
	if cur.key.Cmp(key) == 0 {
		return cur.value
	}
	return nil
}

func (l *SkipList) Remove(key hydrogen.Comparator) bool {
	cur := l.root
	for r := l.maxLevel - 1; r >= 0; r-- {

		for cur.next[r] != nil && cur.next[r].key.Cmp(key) < 0 {
			cur = cur.next[r]
		}

	}
	if cur.key.Cmp(key) == 0 {

		for row, nextElem := range cur.next {
			nextElem.prev[row] = cur.prev[row]
			cur.prev[row].next[row] = nextElem
			cur.prev[row] = nil
			cur.next[row] = nil
		}

		cur.list = nil

		return true
	}
	return false
}