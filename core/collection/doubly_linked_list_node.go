package collection

import (
	. "codeberg.org/yaadata/opt"
)

type DoublyLinkedListNode[T any] interface {
	LinkedListNode[T]
	Next() Option[DoublyLinkedListNode[T]]
	Previous() Option[DoublyLinkedListNode[T]]
}
