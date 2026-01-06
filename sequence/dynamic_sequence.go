package sequence

import (
	. "github.com/yaadata/optionsgo"
)

type DynamicSequence[T any] interface {
	Sequence[T]
	Insert(index int, item T) bool
	RemoveAt(index int) Option[T]
}
