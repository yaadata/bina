package queue

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/sequence/builder"
)

// QueueBackedBy specifies the underlying data structure for the queue.
type QueueBackedBy int

const (
	// QueueBackedBySlice uses a slice as the backing store.
	QueueBackedBySlice QueueBackedBy = iota
	// QueueBackedBySinglyLinkedList uses a singly linked list as the backing store.
	QueueBackedBySinglyLinkedList
)

// Builder is a [builder.BaseBuilder] for [collection.Queue] implementations.
type Builder[T any, Target collection.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	// BackedBy sets the underlying data structure.
	BackedBy(ds QueueBackedBy)
	// From initializes the queue with the given items.
	From(items ...T) Self
}
