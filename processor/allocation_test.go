package processor

import "testing"

func TestNewNearestAllocator_SelectCandidate(t *testing.T) {
	allocator := NewNearestAllocator()
	allocator.SetSize(6)
	slot := allocator.SelectCandidate()
	if slot != 1 {
		t.Errorf("SelectCandidate() got %d want %d", slot, 1)
	}
	allocator.MarkAsAllocated()

	slot = allocator.SelectCandidate()
	if slot != 2 {
		t.Errorf("SelectCandidate() got %d want %d", slot, 2)
	}
	allocator.MarkAsAllocated()

	slot = allocator.SelectCandidate()
	if slot != 3 {
		t.Errorf("SelectCandidate() got %d want %d", slot, 3)
	}
	allocator.MarkAsAllocated()

	allocator.MarkAsAvailable(2)

	slot = allocator.SelectCandidate()
	if slot != 2 {
		t.Errorf("SelectCandidate() got %d want %d", slot, 2)
	}
}

func TestNewNearestAllocator_SelectCandidateAllSlotsUsed(t *testing.T) {
	allocator := NewNearestAllocator()
	allocator.SetSize(2)
	allocator.MarkAsAllocated()
	allocator.MarkAsAllocated()
	slot := allocator.SelectCandidate()
	if slot != 0 {
		t.Errorf("SelectCandidate() got %d want %d", slot, 0)
	}
}