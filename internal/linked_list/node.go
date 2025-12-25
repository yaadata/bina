package linkedlist

import (
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type linkedListNode[T any] struct {
	next  *linkedListNode[T]
	value T
}

var _ sequence.LinkedListNode[int] = (*linkedListNode[int])(nil)

func (l *linkedListNode[T]) Value() T {
	return l.value
}

func (l *linkedListNode[T]) SetValue(value T) {
	l.value = value
}

func (l *linkedListNode[T]) Next() Option[sequence.SinglyLinkedListNode[T]] {
	return optionalNode(l.next)
}

func optionalNode[T any](l *linkedListNode[T]) Option[sequence.SinglyLinkedListNode[T]] {
	if l == nil {
		return None[sequence.SinglyLinkedListNode[T]]()
	}
	var node sequence.SinglyLinkedListNode[T] = l
	return Some(node)
}
