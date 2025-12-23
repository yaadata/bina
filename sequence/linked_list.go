package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type LinkedList[T any] interface {
	Sequence[T]
	GetNodeAt(index int) Option[LinkedListNode[T]]
	Prepend(value T)
}
