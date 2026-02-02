package collection

import (
	"codeberg.org/yaadata/bina/core/predicate"
)

// Aggregate provides methods for querying and iterating over collection elements.
type Aggregate[T any] interface {
	// Any reports whether the collection contains at least one element
	// satisfying the predicate. Returns false for empty collections.
	Any(predicate predicate.Predicate[T]) bool

	// Count returns the number of elements in the collection that
	// satisfy the predicate.
	Count(predicate predicate.Predicate[T]) int

	// Every reports whether all elements in the collection satisfy
	// the predicate. Returns true for empty collections.
	Every(predicate predicate.Predicate[T]) bool

	// ForEach applies fn to each element in the collection.
	// The iteration order depends on the collection type and whether
	// the iteration maintains order.
	ForEach(fn func(T))
}
