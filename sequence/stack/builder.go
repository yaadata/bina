package stack

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/sequence/builder"
)

// StackBackedBy specifies the underlying data structure for the stack.
type StackBackedBy int

const (
	// StackBackedBySlice uses a slice as the backing store.
	StackBackedBySlice StackBackedBy = iota
	// StackBackedBySinglyLinkedList uses a singly linked list as the backing store.
	StackBackedBySinglyLinkedList
)

// Builder is a [builder.BaseBuilder] for [collection.Stack] implementations.
type Builder[T any, Target collection.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	// BackedBy sets the underlying data structure.
	BackedBy(ds StackBackedBy)
	// From initializes the stack with the given items.
	From(items ...T) Self
}
