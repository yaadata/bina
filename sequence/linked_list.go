package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type LinkedList[T any] interface {
	Sequence[T]
	Extend(values ...T)
	ExtendFromSequence(sequence Sequence[T])
	Head() Option[LinkedListNode[T]]
	GetNodeAt(index int) Option[LinkedListNode[T]]
	Prepend(value T)
	Tail() Option[LinkedListNode[T]]
}
