package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type LinkedList[T any] interface {
	Sequence[T]
	RemoveAt(index int) Option[T]
}
