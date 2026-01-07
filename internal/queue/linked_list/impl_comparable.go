package queue

import (
	"iter"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type comparableQueue[T compare.Comparable[T]] struct {
	inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]
}

func _[T compare.Comparable[T]]() {
	var _ sequence.Queue[T] = (*comparableQueue[T])(nil)
}

func QueueFromComparable[T compare.Comparable[T]](inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]) *comparableQueue[T] {
	return &comparableQueue[T]{inner: inner}
}

func (b *comparableQueue[T]) Len() int {
	return b.inner.Len()
}

func (b *comparableQueue[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *comparableQueue[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *comparableQueue[T]) Clear() {
	b.inner.Clear()
}

func (b *comparableQueue[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *comparableQueue[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *comparableQueue[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *comparableQueue[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *comparableQueue[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *comparableQueue[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *comparableQueue[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *comparableQueue[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *comparableQueue[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *comparableQueue[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *comparableQueue[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *comparableQueue[T]) ToSlice() []T {
	return b.inner.ToSlice()
}

func (b *comparableQueue[T]) Enqueue(element T) {
	b.inner.Append(element)
}

func (b *comparableQueue[T]) Dequeue() Option[T] {
	return b.inner.RemoveAt(0)
}

func (b *comparableQueue[T]) Peek() Option[T] {
	head := b.inner.Head()
	if head.IsNone() {
		return None[T]()
	}
	return Some(head.Unwrap().Value())
}
