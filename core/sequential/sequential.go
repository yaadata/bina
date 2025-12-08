package sequential

import (
	"iter"

	"github.com/yaadata/bina/core/collection"
	"github.com/yaadata/bina/core/shared"
	. "github.com/yaadata/optionsgo"
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
	ToSlice() []T
}
