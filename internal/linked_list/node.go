package linkedlist

import "codeberg.org/yaadata/bina/sequence"

type linkedListNode[T any] struct {
	value T
	next  *linkedListNode[T]
}

var _ sequence.LinkedListNode[int] = (*linkedListNode[int])(nil)

func (s *linkedListNode[T]) Value() T {
	return s.value
}

func (s *linkedListNode[T]) SetValue(value T) {
	s.value = value
}
