package collection

import (
	"iter"

	. "codeberg.org/yaadata/opt"
)

type Set[T any] interface {
	Collection[T]
	Aggregate[T]
	Add(value T) bool
	Values() iter.Seq[T]
	Extend(values ...T)
	Difference(other Set[T]) Option[Set[T]]
	Intersect(other Set[T]) Option[Set[T]]
	IsSubsetOf(other Set[T]) bool
	IsSupersetOf(other Set[T]) bool
	Remove(value T) bool
	SymmetricDifference(other Set[T]) Option[Set[T]]
	Union(other Set[T]) Set[T]
}
