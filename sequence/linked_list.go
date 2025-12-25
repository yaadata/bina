package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type LinkedList[T any, Node LinkedListNode[T]] interface {
	Sequence[T]
	Extend(values ...T)
	ExtendFromSequence(sequence Sequence[T])
	Head() Option[Node]
	GetNodeAt(index int) Option[Node]
	Prepend(value T)
	Tail() Option[Node]
}
