package common

import (
	"errors"
	"parking_lot/common/unsafelinkedlist"
)

var (
	ErrSetMemberExists = errors.New("ERR_SET_MEMBER_EXISTS")
	ErrSetMemberNotExists = errors.New("ERR_SET_MEMBER_NOT_EXISTS")
)

// LinkedHashIntSet is a set data-structure for int data type.
// This set maintains insert order of the members of the set.
type LinkedHashIntSet struct {
	list       unsafelinkedlist.IntDoublyLinkedList
	membership map[int]*unsafelinkedlist.IntDLLElement
}

func NewLinkedHashIntSet() *LinkedHashIntSet {
	i := LinkedHashIntSet{}
	i.membership = make(map[int]*unsafelinkedlist.IntDLLElement)
	return &i
}

func (s *LinkedHashIntSet) Add(member int) error {
	if _, ok := s.membership[member]; ok {
		return ErrSetMemberExists
	}
	elem := s.list.Append(member)
	s.membership[member] = elem
	return nil
}

func (s *LinkedHashIntSet) Remove(member int) error {
	if elem, ok := s.membership[member]; ok {
		s.list.Remove(elem)
		delete(s.membership, member)
		return nil
	} else {
		return ErrSetMemberNotExists
	}
}

func (s *LinkedHashIntSet) Members() []int {
	return s.list.Values()
}
