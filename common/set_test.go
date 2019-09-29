package common

import (
	"reflect"
	"testing"
)

func TestLinkedHashIntSet_Add(t *testing.T) {
	set := NewLinkedHashIntSet()
	err := set.Add(10)
	if err != nil {
		t.Errorf("Add() Error %+v encoutered", err)
	}
}

func TestLinkedHashIntSet_AddReturnsErrSetMemberExistsOnDuplicateMember(t *testing.T) {
	set := NewLinkedHashIntSet()
	_ = set.Add(10)
	err := set.Add(10)
	if err != ErrSetMemberExists {
		t.Errorf("Add() Expected Error got %+v want %+v", err, ErrSetMemberExists)
	}
}

func TestLinkedHashIntSet_Remove(t *testing.T) {
	set := NewLinkedHashIntSet()
	_ = set.Add(10)
	_ = set.Add(20)
	err := set.Remove(10)
	if err != nil {
		t.Errorf("Remove() Error got %+v encountered ", err)
	}

	expected := []int{20}
	actual := set.Members()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Remove() got %v want %v", actual, expected)
	}
}

func TestLinkedHashIntSet_RemoveReturnsErrSetMemberNotExistsWithoutAMember(t *testing.T) {
	set := NewLinkedHashIntSet()
	err := set.Remove(10)
	if err != ErrSetMemberNotExists {
		t.Errorf("Remove() Expected Error got %+v want %+v", err, ErrSetMemberNotExists)
	}
}

func TestLinkedHashIntSet_Members(t *testing.T) {
	set := NewLinkedHashIntSet()
	_ = set.Add(10)
	_ = set.Add(9)
	_ = set.Add(11)

	actual := set.Members()
	expected := []int{10, 9, 11}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Members() got %+v want %+v", actual, expected)
	}
}
