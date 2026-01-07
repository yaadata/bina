package stack

import (
	"codeberg.org/yaadata/bina/sequence"
	"codeberg.org/yaadata/bina/sequence/builder"
)

type StackBackedBy int

const (
	StackBackedBySlice StackBackedBy = iota
	StackBackedBySinglyLinkedList
)

type Builder[T any, Target sequence.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	BackedBy(ds StackBackedBy)
	From(items ...T) Self
}
