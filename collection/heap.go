package collection

type Heap struct {
	elements []int
	maxHeap  bool
}

func NewMaxHeap() *Heap {
	return &Heap{
		elements: make([]int, 0),
		maxHeap:  true,
	}
}

func NewMinHeap() *Heap {
	return &Heap{
		elements: make([]int, 0),
		maxHeap:  false,
	}
}

func (h *Heap) Len() int {
	return len(h.elements)
}

func (h *Heap) IsEmpty() bool {
	return h.Len() == 0
}

func (h *Heap) Push(e int) {
	h.elements = append(h.elements, e)
	// if only one element, no need to heapify
	if h.Len() > 1 {
		h.up(h.Len() - 1)
	}
}

func (h *Heap) Pop() *int {
	if h.IsEmpty() {
		return nil
	}
	e := h.elements[0]
	h.elements[0] = h.elements[h.Len()-1]
	h.elements = h.elements[:h.Len()-1]
	h.down(0)
	return &e
}

func (h *Heap) Peek() *int {
	if h.IsEmpty() {
		return nil
	}
	return &h.elements[0]
}

func (h *Heap) up(idx int) {
	for {
		parentIdx := idx / 2
		if parentIdx < 0 || h.compare(parentIdx, idx) {
			break
		}
		h.swap(idx, parentIdx)
		idx = parentIdx
	}
}

func (h *Heap) down(i int) {
	for {
		child := 2*i + 1
		if child >= h.Len() {
			break
		}
		if child+1 < h.Len() && h.compare(child+1, child) {
			child++
		}
		if h.compare(i, child) {
			break
		}
		h.swap(i, child)
		i = child
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *Heap) compare(i, j int) bool {
	if h.maxHeap {
		return h.compareMax(i, j)
	} else {
		return h.compareMin(i, j)
	}
}

func (h *Heap) compareMax(i, j int) bool {
	return h.elements[i] >= h.elements[j]
}

func (h *Heap) compareMin(i, j int) bool {
	return h.elements[i] <= h.elements[j]
}
