package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type Stack[T any] interface {
	Sequence[T]
	Push(element T)
	Pop() Option[T]
	Peek() Option[T]
}
