package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type LinkedListNode[T any] interface {
	Next() Option[LinkedListNode[T]]
	SetValue(value T)
	Value() T
}
