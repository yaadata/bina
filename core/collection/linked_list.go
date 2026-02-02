package collection

import (
	. "codeberg.org/yaadata/opt"
)

// LinkedList is a [DynamicSequence] implemented as a chain of nodes.
type LinkedList[T any, Node LinkedListNode[T]] interface {
	DynamicSequence[T]

	// Append adds an element to the end.
	Append(item T)

	// Extend adds multiple elements to the end.
	Extend(values ...T)

	// ExtendFromSequence adds all elements from the given sequence to the end.
	ExtendFromSequence(sequence Sequence[T])

	// Head returns the first node, or None if empty.
	Head() Option[Node]

	// GetNodeAt returns the node at the given index, or None if out of bounds.
	GetNodeAt(index int) Option[Node]

	// Prepend adds an element to the beginning.
	Prepend(value T)

	// Tail returns the last node, or None if empty.
	Tail() Option[Node]
}
