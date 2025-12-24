package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type DoublyLinkedListNode[T any] interface {
	Next() Option[DoublyLinkedListNode[T]]
	Previous() Option[DoublyLinkedListNode[T]]
	SetValue(value T)
	Value() T
}
