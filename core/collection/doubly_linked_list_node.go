package collection

import (
	. "codeberg.org/yaadata/opt"
)

// DoublyLinkedListNode is a [LinkedListNode] with references to both adjacent nodes.
type DoublyLinkedListNode[T any] interface {
	LinkedListNode[T]

	// Next returns the next node, or None if this is the last node.
	Next() Option[DoublyLinkedListNode[T]]

	// Previous returns the previous node, or None if this is the first node.
	Previous() Option[DoublyLinkedListNode[T]]
}
