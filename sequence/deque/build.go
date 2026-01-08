package deque

import (
	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/compare"
	internal_deque "codeberg.org/yaadata/bina/internal/deque"
	linkedlist "codeberg.org/yaadata/bina/internal/doubly_linked_list"
	"codeberg.org/yaadata/bina/internal/slice"
	"codeberg.org/yaadata/bina/sequence"
)

func NewBuiltinBuilder[T comparable]() Builder[T, sequence.Deque[T], *builtinBuilder[T]] {
	return &builtinBuilder[T]{
		backedBy: DequeBackedBySlice,
		from:     None[[]T](),
	}
}

type builtinBuilder[T comparable] struct {
	backedBy DequeBackedBy
	from     Option[[]T]
}

func (b *builtinBuilder[T]) BackedBy(ds DequeBackedBy) {
	b.backedBy = ds
}

func (b *builtinBuilder[T]) From(items ...T) *builtinBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *builtinBuilder[T]) Build() sequence.Deque[T] {
	items := b.from.UnwrapOrDefault()
	switch b.backedBy {
	case DequeBackedByDoublyLinkedList:
		inner := linkedlist.LinkedListFromBuiltin[T]()
		inner.Extend(items...)
		return internal_deque.LinkedListBackedDequeFromBuiltin(inner)
	default:
		inner := slice.SliceFromBuiltin(items...)
		return internal_deque.SliceBackedDequeFromBuiltin(inner)
	}
}

func NewComparableBuilder[T compare.Comparable[T]]() Builder[T, sequence.Deque[T], *comparableBuilder[T]] {
	return &comparableBuilder[T]{
		backedBy: DequeBackedBySlice,
		from:     None[[]T](),
	}
}

type comparableBuilder[T compare.Comparable[T]] struct {
	backedBy DequeBackedBy
	from     Option[[]T]
}

func (b *comparableBuilder[T]) BackedBy(ds DequeBackedBy) {
	b.backedBy = ds
}

func (b *comparableBuilder[T]) From(items ...T) *comparableBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *comparableBuilder[T]) Build() sequence.Deque[T] {
	items := b.from.UnwrapOrDefault()
	switch b.backedBy {
	case DequeBackedByDoublyLinkedList:
		inner := linkedlist.LinkedListFromComparable[T]()
		inner.Extend(items...)
		return internal_deque.LinkedListBackedDequeFromComparable(inner)
	default:
		inner := slice.SliceFromComparableInterface(items...)
		return internal_deque.SliceBackedDequeFromComparable(inner)
	}
}
