package queue

import (
	"iter"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type queue[T any] struct {
	inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]
}

// Compile-time interface checks
func _[T comparable]() {
	var _ sequence.Queue[T] = (*queue[T])(nil)
}

func __[T compare.Comparable[T]]() {
	var _ sequence.Queue[T] = (*queue[T])(nil)
}

func QueueFromBuiltin[T comparable](inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]) *queue[T] {
	return &queue[T]{inner: inner}
}

func QueueFromComparable[T compare.Comparable[T]](inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]) *queue[T] {
	return &queue[T]{inner: inner}
}

func (b *queue[T]) Len() int {
	return b.inner.Len()
}

func (b *queue[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *queue[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *queue[T]) Clear() {
	b.inner.Clear()
}

func (b *queue[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *queue[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *queue[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *queue[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *queue[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *queue[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *queue[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *queue[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *queue[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *queue[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *queue[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *queue[T]) ToSlice() []T {
	return b.inner.ToSlice()
}

func (b *queue[T]) Enqueue(element T) {
	b.inner.Append(element)
}

func (b *queue[T]) Dequeue() Option[T] {
	return b.inner.RemoveAt(0)
}

func (b *queue[T]) Peek() Option[T] {
	head := b.inner.Head()
	if head.IsNone() {
		return None[T]()
	}
	return Some(head.Unwrap().Value())
}
