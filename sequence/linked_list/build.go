package linkedlist

import (
	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/compare"
	linkedlist "codeberg.org/yaadata/bina/internal/linked_list"
	"codeberg.org/yaadata/bina/sequence"
	"codeberg.org/yaadata/bina/sequence/builder"
)

func NewBuiltinBuilder[T comparable]() Builder[T, sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]], *builtinBuilder[T]] {
	return &builtinBuilder[T]{
		from: None[[]T](),
	}
}

type builtinBuilder[T comparable] struct {
	from Option[[]T]
}

func (b *builtinBuilder[T]) From(items ...T) *builtinBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *builtinBuilder[T]) Build() sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]] {
	ll := linkedlist.LinkedListFromBuiltin[T]()
	ll.Extend(b.from.UnwrapOrDefault()...)
	return ll
}

func NewComparableBuilder[T compare.Comparable[T]]() builder.BaseBuilder[T, sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]], *comparableBuilder[T]] {
	return &comparableBuilder[T]{
		from: None[[]T](),
	}
}

type comparableBuilder[T compare.Comparable[T]] struct {
	from Option[[]T]
}

func (b *comparableBuilder[T]) From(items ...T) *comparableBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *comparableBuilder[T]) Build() sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]] {
	ll := linkedlist.LinkedListFromComparable[T]()
	ll.Extend(b.from.UnwrapOrDefault()...)
	return ll
}
