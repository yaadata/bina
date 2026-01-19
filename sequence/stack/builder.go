package stack

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/sequence/builder"
)

type StackBackedBy int

const (
	StackBackedBySlice StackBackedBy = iota
	StackBackedBySinglyLinkedList
)

type Builder[T any, Target collection.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	BackedBy(ds StackBackedBy)
	From(items ...T) Self
}
