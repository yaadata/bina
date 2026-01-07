package queue

import (
	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/compare"
	linkedlist "codeberg.org/yaadata/bina/internal/linked_list"
	"codeberg.org/yaadata/bina/internal/queue"
	"codeberg.org/yaadata/bina/internal/slice"
	"codeberg.org/yaadata/bina/sequence"
)

func NewBuiltinBuilder[T comparable]() Builder[T, sequence.Queue[T], *builtinBuilder[T]] {
	return &builtinBuilder[T]{
		backedBy: QueueBackedBySlice,
		from:     None[[]T](),
	}
}

type builtinBuilder[T comparable] struct {
	backedBy QueueBackedBy
	from     Option[[]T]
}

func (b *builtinBuilder[T]) BackedBy(ds QueueBackedBy) {
	b.backedBy = ds
}

func (b *builtinBuilder[T]) From(items ...T) *builtinBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *builtinBuilder[T]) Build() sequence.Queue[T] {
	items := b.from.UnwrapOrDefault()
	switch b.backedBy {
	case QueueBackedBySinglyLinkedList:
		inner := linkedlist.LinkedListFromBuiltin[T]()
		inner.Extend(items...)
		return queue.LinkedListBackedQueueFromBuiltin(inner)
	default:
		inner := slice.SliceFromBuiltin(items...)
		return queue.SliceBackedQueueFromBuiltin(inner)
	}
}

func NewComparableBuilder[T compare.Comparable[T]]() Builder[T, sequence.Queue[T], *comparableBuilder[T]] {
	return &comparableBuilder[T]{
		backedBy: QueueBackedBySlice,
		from:     None[[]T](),
	}
}

type comparableBuilder[T compare.Comparable[T]] struct {
	backedBy QueueBackedBy
	from     Option[[]T]
}

func (b *comparableBuilder[T]) BackedBy(ds QueueBackedBy) {
	b.backedBy = ds
}

func (b *comparableBuilder[T]) From(items ...T) *comparableBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *comparableBuilder[T]) Build() sequence.Queue[T] {
	items := b.from.UnwrapOrDefault()
	switch b.backedBy {
	case QueueBackedBySinglyLinkedList:
		inner := linkedlist.LinkedListFromComparable[T]()
		inner.Extend(items...)
		return queue.LinkedListBackedQueueFromComparable(inner)
	default:
		inner := slice.SliceFromComparableInterface(items...)
		return queue.SliceBackedQueueFromComparable(inner)
	}
}
