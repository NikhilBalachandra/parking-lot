package unsafelinkedlist

import (
	"testing"
)

func TestIntDoublyLinkedList_AppendEmptyList(t *testing.T) {
	list := IntDoublyLinkedList{}
	list.Append(10)
	if list.size != 1 {
		t.Errorf("Append() got %d, want %d", list.size, 1)
	}
	if list.first.value != 10 || list.last.value != 10 {
		t.Errorf("Append() got %d, got2 %d, want %d", list.first.value, list.last.value, 10)
	}
}

func TestIntDoublyLinkedList_AppendNonEmptyList(t *testing.T) {
	list := IntDoublyLinkedList{}
	list.Append(10)
	list.Append(20)
	if list.size != 2 {
		t.Errorf("Append() got %d, want %d", list.size, 2)
	}
	if list.first.value != 10 || list.last.value != 20 {
		t.Errorf("Append() got %d, got2 %d, want %d, want2 %d", list.first.value, list.last.value, 10, 20)
	}

	if list.first.prev != nil || list.last.next != nil {
		t.Errorf("Append() First elem shouldn't have previous and last element shouldn't have next")
	}

	if list.first.next != list.last || list.last.prev != list.first {
		t.Errorf("Append() First elem and/or last element don't point correctly")
	}
}

func TestIntDoublyLinkedList_RemoveFirstElement(t *testing.T) {
	list := IntDoublyLinkedList{}
	elem := list.Append(10)
	elem2 := list.Append(20)
	list.Remove(elem)

	if list.first != elem2 || elem2.prev != nil || list.last != elem2 || list.size != 1 || list.first.value != 20 {
		t.Errorf("Remove() Invalid internal state %+v", list)
	}
}

func TestIntDoublyLinkedList_RemoveLastElement(t *testing.T) {
	list := IntDoublyLinkedList{}
	elem := list.Append(10)
	elem2 := list.Append(20)
	list.Remove(elem2)

	if list.first != elem || elem.next != nil || list.last != elem || list.size != 1 || list.first.value != 10 {
		t.Errorf("Remove() Invalid internal state %+v", list)
	}
}

func TestIntDoublyLinkedList_RemoveOnlyAvailableElement(t *testing.T) {
	list := IntDoublyLinkedList{}
	elem := list.Append(10)
	list.Remove(elem)

	if list.first != nil || list.last != nil || list.size != 0 {
		t.Errorf("Remove() Invalid internal state %+v", list)
	}
}

func TestIntDoublyLinkedList_RemoveMiddleElement(t *testing.T) {
	list := IntDoublyLinkedList{}
	elem := list.Append(10)
	elem2 := list.Append(20)
	elem3 := list.Append(30)
	list.Remove(elem2)

	if list.first != elem || list.last != elem3 || list.size != 2 ||
		list.first.value != 10 || list.last.value != 30 ||
		list.first.next != list.last || list.last.prev != list.first ||
		list.first.prev != nil || list.last.next != nil {
		t.Errorf("Remove() Invalid internal state %+v", list)
	}
}
