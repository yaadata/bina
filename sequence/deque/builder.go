package deque

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/sequence/builder"
)

// DequeBackedBy specifies the underlying data structure for the deque.
type DequeBackedBy int

const (
	// DequeBackedBySlice uses a slice as the backing store.
	DequeBackedBySlice DequeBackedBy = iota
	// DequeBackedByDoublyLinkedList uses a doubly linked list as the backing store.
	DequeBackedByDoublyLinkedList
)

// Builder is a [builder.BaseBuilder] for [collection.Deque] implementations.
type Builder[T any, Target collection.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	// BackedBy sets the underlying data structure.
	BackedBy(ds DequeBackedBy)
	// From initializes the deque with the given items.
	From(items ...T) Self
}
