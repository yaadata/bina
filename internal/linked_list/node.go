package linkedlist

import (
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type linkedListNode[T any] struct {
	value T
	next  *linkedListNode[T]
}

var _ sequence.LinkedListNode[int] = (*linkedListNode[int])(nil)

func (l *linkedListNode[T]) Value() T {
	return l.value
}

func (l *linkedListNode[T]) SetValue(value T) {
	l.value = value
}

func (l *linkedListNode[T]) Next() Option[sequence.LinkedListNode[T]] {
	return optionalNode(l.next)
}

func optionalNode[T any](l *linkedListNode[T]) Option[sequence.LinkedListNode[T]] {
	if l == nil {
		return None[sequence.LinkedListNode[T]]()
	}
	var node sequence.LinkedListNode[T] = l
	return Some(node)
}
