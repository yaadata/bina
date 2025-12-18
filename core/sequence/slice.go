package sequence

import (
	"github.com/yaadata/bina/core/shared"
	. "github.com/yaadata/optionsgo"
)

type Slice[T any] interface {
	Sequence[T]
	Extend(items ...T)
	ExtendFromSequence(sequence Sequence[T])
	Filter(predicate shared.Predicate[T]) Slice[T]
	First() Option[T]
	Last() Option[T]
	RemoveAt(index int) T
}
