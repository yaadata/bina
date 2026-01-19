package linkedlist

import (
	"codeberg.org/yaadata/bina/core/collection"
	. "codeberg.org/yaadata/opt"
)

type linkedListNode[T any] struct {
	next  *linkedListNode[T]
	value T
}

var _ collection.LinkedListNode[int] = (*linkedListNode[int])(nil)

func (l *linkedListNode[T]) Value() T {
	return l.value
}

func (l *linkedListNode[T]) SetValue(value T) {
	l.value = value
}

func (l *linkedListNode[T]) Next() Option[collection.SinglyLinkedListNode[T]] {
	return optionalNode(l.next)
}

func optionalNode[T any](l *linkedListNode[T]) Option[collection.SinglyLinkedListNode[T]] {
	if l == nil {
		return None[collection.SinglyLinkedListNode[T]]()
	}
	var node collection.SinglyLinkedListNode[T] = l
	return Some(node)
}
