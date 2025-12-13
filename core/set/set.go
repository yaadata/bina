package sets

import (
	"iter"

	"github.com/yaadata/bina/core/collection"
	. "github.com/yaadata/optionsgo"
)

type Set[T any] interface {
	collection.Collection[T]
	Seq() iter.Seq[T]
	Add(value T) bool
	Remove(value T) bool
	All() iter.Seq[T]

	Union(other Set[T]) Option[Set[T]]
	Intersect(other Set[T]) Option[Set[T]]
	Difference(other Set[T]) Option[Set[T]]
	SymmetricDifference(other Set[T]) Option[Set[T]]
	IsSubsetOf(other Set[T]) bool
	IsSupersetOf(other Set[T]) bool
}
