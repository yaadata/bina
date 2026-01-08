package sequence

import (
	. "codeberg.org/yaadata/opt"
)

type Queue[T any] interface {
	Sequence[T]
	Enqueue(item T)
	Dequeue() Option[T]
	Peek() Option[T]
}
