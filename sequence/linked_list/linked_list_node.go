package linkedlist

import "github.com/yaadata/bina/core/sequence"

type singlyLinkedListNode[T any] struct {
	value T
	next  *singlyLinkedListNode[T]
}

var _ sequence.LinkedListNode[int] = (*singlyLinkedListNode[int])(nil)

func (s *singlyLinkedListNode[T]) Value() T {
	return s.value
}

func (s *singlyLinkedListNode[T]) SetValue(value T) {
	s.value = value
}
