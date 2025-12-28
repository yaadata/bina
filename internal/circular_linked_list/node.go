package circularlinkedlist

import (
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type linkedListNode[T any] struct {
	next     *linkedListNode[T]
	previous *linkedListNode[T]
	value    T
}

func newLinkedListNode[T any](value T) *linkedListNode[T] {
	return &linkedListNode[T]{
		next:     nil,
		previous: nil,
		value:    value,
	}
}

var _ sequence.DoublyLinkedListNode[int] = (*linkedListNode[int])(nil)

func (l *linkedListNode[T]) Next() Option[sequence.DoublyLinkedListNode[T]] {
	return optionalNode(l.next)
}

func (l *linkedListNode[T]) Previous() Option[sequence.DoublyLinkedListNode[T]] {
	return optionalNode(l.previous)
}

func (l *linkedListNode[T]) SetValue(value T) {
	l.value = value
}

func (l *linkedListNode[T]) Value() T {
	return l.value
}

func connectNodes[T any](left, right *linkedListNode[T]) {
	left.next = right
	right.previous = left
}

func optionalNode[T any](l *linkedListNode[T]) Option[sequence.DoublyLinkedListNode[T]] {
	if l == nil {
		return None[sequence.DoublyLinkedListNode[T]]()
	}
	var node sequence.DoublyLinkedListNode[T] = l
	return Some(node)
}
