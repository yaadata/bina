package queue

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/sequence/builder"
)

type QueueBackedBy int

const (
	QueueBackedBySlice QueueBackedBy = iota
	QueueBackedBySinglyLinkedList
)

type Builder[T any, Target collection.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	BackedBy(ds QueueBackedBy)
	From(items ...T) Self
}
