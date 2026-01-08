package queue

import (
	"iter"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
	. "codeberg.org/yaadata/opt"
)

type queueSlice[T any] struct {
	inner sequence.Slice[T]
}

// Compile-time interface checks
func _[T comparable]() {
	var _ sequence.Queue[T] = (*queueSlice[T])(nil)
}

func __[T compare.Comparable[T]]() {
	var _ sequence.Queue[T] = (*queueSlice[T])(nil)
}

func SliceBackedQueueFromBuiltin[T comparable](inner sequence.Slice[T]) *queueSlice[T] {
	return &queueSlice[T]{inner: inner}
}

func SliceBackedQueueFromComparable[T compare.Comparable[T]](inner sequence.Slice[T]) *queueSlice[T] {
	return &queueSlice[T]{inner: inner}
}

func (b *queueSlice[T]) Len() int {
	return b.inner.Len()
}

func (b *queueSlice[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *queueSlice[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *queueSlice[T]) Clear() {
	b.inner.Clear()
}

func (b *queueSlice[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *queueSlice[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *queueSlice[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *queueSlice[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *queueSlice[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *queueSlice[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *queueSlice[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *queueSlice[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *queueSlice[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *queueSlice[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *queueSlice[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *queueSlice[T]) Enqueue(element T) {
	b.inner.Append(element)
}

func (b *queueSlice[T]) Dequeue() Option[T] {
	return b.inner.RemoveAt(0)
}

func (b *queueSlice[T]) Peek() Option[T] {
	return b.inner.First()
}
