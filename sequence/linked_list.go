package sequence

import (
	. "codeberg.org/yaadata/opt"
)

type LinkedList[T any, Node LinkedListNode[T]] interface {
	DynamicSequence[T]
	Append(item T)
	Extend(values ...T)
	ExtendFromSequence(sequence Sequence[T])
	Head() Option[Node]
	GetNodeAt(index int) Option[Node]
	Prepend(value T)
	Tail() Option[Node]
}
