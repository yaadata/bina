package collection

import (
	"iter"

	"codeberg.org/yaadata/bina/core/predicate"
	. "codeberg.org/yaadata/opt"
)

// Slice is a growable [DynamicSequence] backed by a contiguous array.
type Slice[T any] interface {
	DynamicSequence[T]

	// Append adds an element to the end.
	Append(element T)

	// Capacity returns the current capacity of the underlying array.
	Capacity() int

	// Extend adds multiple elements to the end.
	Extend(element ...T)

	// ExtendFromSequence adds all elements from the given sequence to the end.
	ExtendFromSequence(sequence Sequence[T])

	// Filter returns a new slice containing only elements matching the predicate.
	Filter(predicate predicate.Predicate[T]) Slice[T]

	// First returns the first element, or None if empty.
	First() Option[T]

	// Last returns the last element, or None if empty.
	Last() Option[T]

	// Reverse returns an iterator over index-value pairs in reverse order.
	Reverse() iter.Seq2[int, T]
}
