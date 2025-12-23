package sequence

import (
	"codeberg.org/yaadata/bina/core/predicate"
	. "github.com/yaadata/optionsgo"
)

type Slice[T any] interface {
	Sequence[T]
	Extend(items ...T)
	ExtendFromSequence(sequence Sequence[T])
	Filter(predicate predicate.Predicate[T]) Slice[T]
	First() Option[T]
	Last() Option[T]
}
