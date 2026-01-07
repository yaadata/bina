package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type Deque[T any] interface {
	Sequence[T]
	PushFront(item T)
	PushBack(item T)
	PopFront() Option[T]
	PopBack() Option[T]
	PeekFront() Option[T]
	PeekBack() Option[T]
}
