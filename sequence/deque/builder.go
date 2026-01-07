package deque

import (
	"codeberg.org/yaadata/bina/sequence"
	"codeberg.org/yaadata/bina/sequence/builder"
)

type DequeBackedBy int

const (
	DequeBackedBySlice DequeBackedBy = iota
	DequeBackedByDoublyLinkedList
)

type Builder[T any, Target sequence.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	BackedBy(ds DequeBackedBy)
	From(items ...T) Self
}
