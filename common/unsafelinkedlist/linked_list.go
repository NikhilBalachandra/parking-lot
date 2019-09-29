package unsafelinkedlist

// IntDoublyLinkedList is a special purpose doubly-linked list that
// returns internal memory address to the caller. This allows fast random
// access to the linked list elements for removal or addition of a node.
// Incorrect usage of this can potentially crash the program.
type IntDoublyLinkedList struct {
	first *IntDLLElement
	last  *IntDLLElement
	size  int
}

// IntDLLElement represents a node in the doubly-linked list
// Members of this struct are unexported to limit access outside
// this package.
type IntDLLElement struct {
	value int
	prev  *IntDLLElement
	next  *IntDLLElement
}

// Append item to the set. Returns element inserted.
// IMP: Be careful not to modify returned element
func (list *IntDoublyLinkedList) Append(item int) *IntDLLElement {
	newElement := &IntDLLElement{value: item, prev: list.last}
	if list.size == 0 {
		list.first = newElement
		list.last = newElement
	} else {
		list.last.next = newElement
		list.last = newElement
	}
	list.size++
	return newElement
}

// Remove removes element passed in the argument from the linked list.
// It is the job of the caller to ensure this function isn't called
// with the invalid address.
func (list *IntDoublyLinkedList) Remove(elem *IntDLLElement) {
	prev := elem.prev
	next := elem.next

	// Element to be removed is the only element in the list.
	if prev == nil && next == nil {
		list.first = nil
		list.last = nil
		list.size = 0
		return
	}

	if prev == nil { // Element to be removed is first element.
		list.first = next
		next.prev = nil
	} else {
		prev.next = next
	}

	if next == nil { // Element to be removed is last element
		list.last = prev
		prev.next = nil
	} else {
		next.prev = prev
	}
	list.size--
}

// Values returns list of element values in the linked list.
func (list *IntDoublyLinkedList) Values() []int {
	values := make([]int, list.size, list.size)
	for e, item := 0, list.first; item != nil; e, item = e+1, item.next {
		values[e] = item.value
	}
	return values
}
