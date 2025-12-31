package set

import (
	"iter"

	"codeberg.org/yaadata/bina/core/collection"
	. "github.com/yaadata/optionsgo"
)

type Set[T any] interface {
	collection.Collection[T]
	collection.Aggregate[T]
	Add(value T) bool
	All() iter.Seq[T]
	Extend(values ...T)
	Difference(other Set[T]) Option[Set[T]]
	Intersect(other Set[T]) Option[Set[T]]
	IsSubsetOf(other Set[T]) bool
	IsSupersetOf(other Set[T]) bool
	Remove(value T) bool
	SymmetricDifference(other Set[T]) Option[Set[T]]
	Union(other Set[T]) Set[T]
}
