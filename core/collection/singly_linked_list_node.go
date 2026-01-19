package collection

import (
	. "codeberg.org/yaadata/opt"
)

type SinglyLinkedListNode[T any] interface {
	LinkedListNode[T]
	Next() Option[SinglyLinkedListNode[T]]
}
