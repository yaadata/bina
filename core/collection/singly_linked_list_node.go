package collection

import (
	. "codeberg.org/yaadata/opt"
)

// SinglyLinkedListNode is a [LinkedListNode] with a reference to the next node.
type SinglyLinkedListNode[T any] interface {
	LinkedListNode[T]

	// Next returns the next node, or None if this is the last node.
	Next() Option[SinglyLinkedListNode[T]]
}
