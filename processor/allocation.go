package processor

import (
	"container/heap"
	"parking_lot/common"
)

// Allocator is a interface type to deal with the allocating slots for the parking.
type Allocator interface {
	// MarkAsAllocated marks the slot as allocated.
	MarkAsAllocated()
	// MarkAsAvailable frees up the slot.
	MarkAsAvailable(slotID int)
	// SelectCandidate returns the slot next to be allocated.
	SelectCandidate() int
	// SetSize sets the size of the parking lot.
	SetSize(size int)
	// GetSize returns the size of the parking lot.
	GetSize() int
}

// NearestAllocator allocates slot nearest to the entrance for the incoming car.
type NearestAllocator struct {
	size int
	heap common.IntMinHeap
}

// NewNearestAllocator builds and returns the NearestAllocator
func NewNearestAllocator() NearestAllocator {
	n := NearestAllocator{}
	n.heap = common.IntMinHeap{}
	heap.Init(&n.heap)
	return n
}

func (na *NearestAllocator) SetSize(size int) {
	na.size = size
	for i := 1; i <= size; i++ {
		heap.Push(&na.heap, i)
	}
}

func (na *NearestAllocator) GetSize() int {
	return na.size
}

func (na *NearestAllocator) MarkAsAllocated() {
	heap.Remove(&na.heap, 0)
}

func (na *NearestAllocator) MarkAsAvailable(slotID int) {
	heap.Push(&na.heap, slotID)
}

func (na *NearestAllocator) SelectCandidate() int {
	top := na.heap.Peek()
	if top != nil {
		return top.(int)
	}
	return 0
}
