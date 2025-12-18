package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type LinkedList[T any] interface {
	Sequence[T]
	RemoveAt(index int) Option[T]
	GetNodeAt(index int) Option[LinkedListNode[T]]
	Prepend(value T)
	InsertAt(index int, value T) bool
}
