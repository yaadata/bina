package sequence

import (
	"iter"

	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
)

type Sequence[T any] interface {
	collection.Collection[T]
	collection.Aggregate[T]
	All() iter.Seq[T]
	Enumerate() iter.Seq2[int, T]
	Find(predicate predicate.Predicate[T]) Option[T]
	FindIndex(predicate predicate.Predicate[T]) Option[int]
	Get(index int) Option[T]
	Retain(predicate predicate.Predicate[T])
	Sort(fn func(a, b T) compare.Order)
	ToSlice() []T
}
