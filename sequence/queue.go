package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type Queue[T any] interface {
	Sequence[T]
	Enqueue(item T)
	Dequeue() Option[T]
	Peek() Option[T]
}
