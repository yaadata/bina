package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type LinkedListNode[T any] interface {
	Value() T
	SetValue(value T)
	Next() Option[LinkedListNode[T]]
}
