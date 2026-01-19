package collection

import (
	. "codeberg.org/yaadata/opt"
)

type Stack[T any] interface {
	Sequence[T]
	Push(element T)
	Pop() Option[T]
	Peek() Option[T]
}
