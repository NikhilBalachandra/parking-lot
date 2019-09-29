package common

// IntMinHeap implements Heap interface from container/heap.
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Peek () interface{} {
	if h.Len() > 0 {
		return (*h)[0]
	}
	return nil
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}