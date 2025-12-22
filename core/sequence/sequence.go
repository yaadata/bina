package sequence

import (
	"iter"

	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/shared"
)

type Sequence[T any] interface {
	collection.Collection[T]
	collection.Aggregate[T]
	Append(item T)
	All() iter.Seq[T]
	Enumerate() iter.Seq2[int, T]
	Find(predicate shared.Predicate[T]) Option[T]
	FindIndex(predicate shared.Predicate[T]) Option[int]
	Get(index int) Option[T]
	Insert(index int, item T)
	RemoveAt(index int) Option[T]
	Retain(predicate shared.Predicate[T])

	Sort(fn func(a, b T) compare.Order)
	ToSlice() []T
}
