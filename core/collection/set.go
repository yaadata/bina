package collection

import (
	"iter"

	. "codeberg.org/yaadata/opt"
)

// Set defines operations for an unordered collection of unique elements.
type Set[T any] interface {
	Collection[T]
	Aggregate[T]

	// Add inserts value into the set.
	// Returns true if value was added, false if it already existed.
	Add(value T) bool

	// Values returns an iterator over all elements in the set.
	Values() iter.Seq[T]

	// Extend adds multiple values to the set.
	// Duplicate values are ignored.
	Extend(values ...T)

	// Difference returns elements in this set but not in other.
	// Returns None if the result is empty.
	Difference(other Set[T]) Option[Set[T]]

	// Intersect returns elements present in both sets.
	// Returns None if the intersection is empty.
	Intersect(other Set[T]) Option[Set[T]]

	// IsSubsetOf reports whether all elements in this set exist in other.
	IsSubsetOf(other Set[T]) bool

	// IsSupersetOf reports whether all elements in other exist in this set.
	IsSupersetOf(other Set[T]) bool

	// Remove deletes value from the set.
	// Returns true if value was removed, false if it did not exist.
	Remove(value T) bool

	// SymmetricDifference returns elements in either set but not both.
	// Returns None if the result is empty.
	SymmetricDifference(other Set[T]) Option[Set[T]]

	// Union returns a new set containing all elements from both sets.
	Union(other Set[T]) Set[T]
}
