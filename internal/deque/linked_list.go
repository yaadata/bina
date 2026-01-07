package deque

import (
	"iter"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type dequeLinkedList[T any] struct {
	inner sequence.LinkedList[T, sequence.DoublyLinkedListNode[T]]
}

// Compile-time interface checks
func ___[T comparable]() {
	var _ sequence.Deque[T] = (*dequeLinkedList[T])(nil)
}

func ____[T compare.Comparable[T]]() {
	var _ sequence.Deque[T] = (*dequeLinkedList[T])(nil)
}

func LinkedListBackedDequeFromBuiltin[T comparable](inner sequence.LinkedList[T, sequence.DoublyLinkedListNode[T]]) *dequeLinkedList[T] {
	return &dequeLinkedList[T]{inner: inner}
}

func LinkedListBackedDequeFromComparable[T compare.Comparable[T]](inner sequence.LinkedList[T, sequence.DoublyLinkedListNode[T]]) *dequeLinkedList[T] {
	return &dequeLinkedList[T]{inner: inner}
}

func (b *dequeLinkedList[T]) Len() int {
	return b.inner.Len()
}

func (b *dequeLinkedList[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *dequeLinkedList[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *dequeLinkedList[T]) Clear() {
	b.inner.Clear()
}

func (b *dequeLinkedList[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *dequeLinkedList[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *dequeLinkedList[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *dequeLinkedList[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *dequeLinkedList[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *dequeLinkedList[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *dequeLinkedList[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *dequeLinkedList[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *dequeLinkedList[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *dequeLinkedList[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *dequeLinkedList[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *dequeLinkedList[T]) PushFront(element T) {
	b.inner.Prepend(element)
}

func (b *dequeLinkedList[T]) PushBack(element T) {
	b.inner.Append(element)
}

func (b *dequeLinkedList[T]) PopFront() Option[T] {
	return b.inner.RemoveAt(0)
}

func (b *dequeLinkedList[T]) PopBack() Option[T] {
	return b.inner.RemoveAt(b.Len() - 1)
}

func (b *dequeLinkedList[T]) PeekFront() Option[T] {
	head := b.inner.Head()
	if head.IsNone() {
		return None[T]()
	}
	return Some(head.Unwrap().Value())
}

func (b *dequeLinkedList[T]) PeekBack() Option[T] {
	tail := b.inner.Tail()
	if tail.IsNone() {
		return None[T]()
	}
	return Some(tail.Unwrap().Value())
}
