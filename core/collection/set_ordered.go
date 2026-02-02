package collection

import (
	"iter"

	. "codeberg.org/yaadata/opt"
)

// OrderedSet extends Set with insertion-order preservation.
// Elements are iterated in the order they were added.
type OrderedSet[T any] interface {
	Set[T]

	// All returns an iterator of (index, element) pairs in insertion order.
	// Indices are contiguous starting from 0.
	All() iter.Seq2[int, T]

	// First returns the first element added to the set.
	// Returns None if the set is empty.
	First() Option[T]

	// Last returns the most recently added element.
	// Returns None if the set is empty.
	Last() Option[T]
}
