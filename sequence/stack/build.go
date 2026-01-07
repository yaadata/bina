package stack

import (
	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/compare"
	linkedlist "codeberg.org/yaadata/bina/internal/linked_list"
	"codeberg.org/yaadata/bina/internal/slice"
	stacklinkedlist "codeberg.org/yaadata/bina/internal/stack/linked_list"
	stackslice "codeberg.org/yaadata/bina/internal/stack/slice"
	"codeberg.org/yaadata/bina/sequence"
)

func NewBuiltinBuilder[T comparable]() Builder[T, sequence.Stack[T], *builtinBuilder[T]] {
	return &builtinBuilder[T]{
		backedBy: StackBackedBySlice,
		from:     None[[]T](),
	}
}

type builtinBuilder[T comparable] struct {
	backedBy StackBackedBy
	from     Option[[]T]
}

func (b *builtinBuilder[T]) BackedBy(ds StackBackedBy) {
	b.backedBy = ds
}

func (b *builtinBuilder[T]) From(items ...T) *builtinBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *builtinBuilder[T]) Build() sequence.Stack[T] {
	items := b.from.UnwrapOrDefault()
	switch b.backedBy {
	case StackBackedBySinglyLinkedList:
		inner := linkedlist.LinkedListFromBuiltin[T]()
		inner.Extend(items...)
		return stacklinkedlist.StackFromBuiltin(inner)
	default:
		inner := slice.SliceFromBuiltin(items...)
		return stackslice.StackFromBuiltin(inner)
	}
}

func NewComparableBuilder[T compare.Comparable[T]]() Builder[T, sequence.Stack[T], *comparableBuilder[T]] {
	return &comparableBuilder[T]{
		backedBy: StackBackedBySlice,
		from:     None[[]T](),
	}
}

type comparableBuilder[T compare.Comparable[T]] struct {
	backedBy StackBackedBy
	from     Option[[]T]
}

func (b *comparableBuilder[T]) BackedBy(ds StackBackedBy) {
	b.backedBy = ds
}

func (b *comparableBuilder[T]) From(items ...T) *comparableBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *comparableBuilder[T]) Build() sequence.Stack[T] {
	items := b.from.UnwrapOrDefault()
	switch b.backedBy {
	case StackBackedBySinglyLinkedList:
		inner := linkedlist.LinkedListFromComparable[T]()
		inner.Extend(items...)
		return stacklinkedlist.StackFromComparable(inner)
	default:
		inner := slice.SliceFromComparableInterface(items...)
		return stackslice.StackFromComparable(inner)
	}
}
