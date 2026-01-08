package sequence

import (
	"iter"

	"codeberg.org/yaadata/bina/core/predicate"
	. "codeberg.org/yaadata/opt"
)

type Slice[T any] interface {
	DynamicSequence[T]
	Append(element T)
	Capacity() int
	Extend(element ...T)
	ExtendFromSequence(sequence Sequence[T])
	Filter(predicate predicate.Predicate[T]) Slice[T]
	First() Option[T]
	Last() Option[T]
	Reverse() iter.Seq2[int, T]
}
