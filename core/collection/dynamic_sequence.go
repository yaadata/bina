package collection

import (
	. "codeberg.org/yaadata/opt"
)

// DynamicSequence is a [Sequence] that supports insertion and removal at arbitrary positions.
type DynamicSequence[T any] interface {
	Sequence[T]
	// Insert adds an element at the given index, returning true on success.
	Insert(index int, item T) bool
	// RemoveAt removes and returns the element at the given index, or None if out of bounds.
	RemoveAt(index int) Option[T]
}
