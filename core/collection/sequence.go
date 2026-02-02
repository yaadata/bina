package collection

import (
	"iter"

	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
)

// Sequence is a positionally-ordered collection of elements accessible by index.
type Sequence[T any] interface {
	Collection[T]
	Aggregate[T]

	// All returns an iterator over index-value pairs in order.
	All() iter.Seq2[int, T]

	// Values returns an iterator over values in order.
	Values() iter.Seq[T]

	// Find returns the first element matching the predicate, or None if not found.
	Find(predicate predicate.Predicate[T]) Option[T]

	// FindIndex returns the index of the first element matching the predicate, or None if not found.
	FindIndex(predicate predicate.Predicate[T]) Option[int]

	// Get returns the element at the given index, or None if out of bounds.
	Get(index int) Option[T]

	// Retain keeps only elements satisfying the predicate, removing others in-place.
	Retain(predicate predicate.Predicate[T])

	// Sort orders elements in-place using the provided comparison function.
	Sort(fn func(a, b T) compare.Order)
}
