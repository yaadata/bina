package doublylinkedlist

import (
	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/compare"
	linkedlist "codeberg.org/yaadata/bina/internal/doubly_linked_list"
)

func NewBuiltinBuilder[T comparable]() Builder[T, collection.LinkedList[T, collection.DoublyLinkedListNode[T]], *builtinBuilder[T]] {
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

func (b *builtinBuilder[T]) Build() collection.LinkedList[T, collection.DoublyLinkedListNode[T]] {
	ll := linkedlist.LinkedListFromBuiltin[T]()
	ll.Extend(b.from.UnwrapOrDefault()...)
	return ll
}

func NewComparableBuilder[T compare.Comparable[T]]() Builder[T, collection.LinkedList[T, collection.DoublyLinkedListNode[T]], *comparableBuilder[T]] {
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

func (b *comparableBuilder[T]) Build() collection.LinkedList[T, collection.DoublyLinkedListNode[T]] {
	ll := linkedlist.LinkedListFromComparable[T]()
	ll.Extend(b.from.UnwrapOrDefault()...)
	return ll
}
