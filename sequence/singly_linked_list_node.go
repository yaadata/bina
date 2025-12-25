package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type SinglyLinkedListNode[T any] interface {
	LinkedListNode[T]
	Next() Option[SinglyLinkedListNode[T]]
}
