package collection

import (
	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/core/where"
)

// Array is a fixed-size [Sequence] where elements can be replaced but not inserted or removed.
type Array[T any] interface {
	Sequence[T]
	// Filter returns a new array containing only elements matching the predicate.
	Filter(predicate predicate.Predicate[T]) Array[T]
	// First returns the first element, or None if empty.
	First() Option[T]
	// Last returns the last element, or None if empty.
	Last() Option[T]
	// Offer replaces the element at the given index, returning true on success.
	Offer(element T, index int) bool
	// OfferRange replaces elements in the specified range, returning true on success.
	OfferRange(elements []T, cfgs ...where.WhereOption[int]) bool
}
