package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type DoublyLinkedListNode[T any] interface {
	LinkedListNode[T]
	Next() Option[DoublyLinkedListNode[T]]
	Previous() Option[DoublyLinkedListNode[T]]
}
