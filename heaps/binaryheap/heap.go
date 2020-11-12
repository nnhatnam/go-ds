package binaryheap

import "fmt"

//max-heap implementation
type Heap struct {
	Data []int
	Len int
}


func NewHeap(cap int) *Heap {
	h := new(Heap)
	h.Len = 0
	h.Data = make([]int, cap + 1)
	return h
}

func (h *Heap) Size() int {
	return h.Len
}

func (h *Heap) IsEmpty() bool {
	return h.Len == 0
}

func left(h *Heap, pos int) int {
	if isLeap(h, pos) {
		return 0
	}

	return 2*pos
}

func (h *Heap) left(pos int) int {
	return left(h, pos)
}

func right(h *Heap, pos int) int {
	if isLeap(h, pos) {
		return 0
	}

	return 2*pos + 1
}

func (h *Heap) right(pos int) int {
	return right(h, pos)
}

func (h *Heap) data(pos int) int {
	return h.Data[pos]
}

func (h *Heap) swap(pos1, pos2 int) {
	h.Data[pos1], h.Data[pos2] = h.Data[pos2], h.Data[pos1]
}

func parent(h *Heap, pos int) (int, bool){
	if pos == 1 {
		return 0, false
	}
	return pos / 2, true
}

func isLeap(h *Heap, pos int) bool {
	if pos >= h.Len / 2 && pos <= h.Len {
		return true
	}
	return false
}

func (h *Heap) insert(data int) {
	pos := h.Len + 1
	h.Data[pos] = data
	for ; pos / 2 != 0 && h.Data[pos] > h.Data[pos/2]  ; {
		h.Data[pos/2], h.Data[pos] = h.Data[pos], h.Data[pos/2]
		pos = pos / 2
	}
	h.Len++
}

func (h *Heap) Insert(data int) {
	h.insert(data)
}

func (h *Heap) Print()  {
	for i := 1; i < h.Len / 2 + 1; i++ {
		fmt.Println(fmt.Sprintf(`
			PARENT : + %v 
			LEFT CHILD : %v 
			RIGHT CHILD : %v 
		`,h.data(i), h.data(h.left(i)), h.data(h.right(i))))
	}

}

func maxHeapify(h *Heap, pos int) {
	if !isLeap(h, pos) {
		if h.data(pos) < h.data(h.left(pos)) {
			h.swap(pos, h.left(pos))
			maxHeapify( h, h.left(pos))
		} else if h.data(pos) < h.data(h.right(pos)) {
			h.swap(pos, h.right(pos))
			maxHeapify( h, h.left(pos))
		}
	}
}
