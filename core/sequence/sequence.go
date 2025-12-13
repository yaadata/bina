package sequence

import (
	"iter"

	. "github.com/yaadata/optionsgo"

	"github.com/yaadata/bina/core/collection"
	"github.com/yaadata/bina/core/compare"
	"github.com/yaadata/bina/core/shared"
)

type Sequence[T any] interface {
	collection.Collection[T]
	collection.Aggregate[T]
	Append(item T) Sequence[T]
	All() iter.Seq[T]
	Enumerate() iter.Seq2[int, T]
	Extend(items ...T) Sequence[T]
	ExtendFromSequence(sequence Sequence[T]) Sequence[T]
	Last() Option[T]
	Filter(predicate shared.Predicate[T]) Sequence[T]
	Find(predicate shared.Predicate[T]) Option[T]
	FindIndex(predicate shared.Predicate[T]) Option[int]
	First() Option[T]
	Get(index int) Option[T]
	Insert(index int, item T) Sequence[T]
	RemoveAt(index int) T
	Retain(predicate shared.Predicate[T]) Sequence[T]
	Sort(fn func(a, b T) compare.Order) Sequence[T]
	ToSlice() []T
}
