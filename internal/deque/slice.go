package deque

import (
	"iter"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	. "codeberg.org/yaadata/opt"
)

type dequeSlice[T any] struct {
	inner collection.Slice[T]
}

// Compile-time interface checks
func _[T comparable]() {
	var _ collection.Deque[T] = (*dequeSlice[T])(nil)
}

func __[T compare.Comparable[T]]() {
	var _ collection.Deque[T] = (*dequeSlice[T])(nil)
}

func SliceBackedDequeFromBuiltin[T comparable](inner collection.Slice[T]) *dequeSlice[T] {
	return &dequeSlice[T]{inner: inner}
}

func SliceBackedDequeFromComparable[T compare.Comparable[T]](inner collection.Slice[T]) *dequeSlice[T] {
	return &dequeSlice[T]{inner: inner}
}

func (b *dequeSlice[T]) Len() int {
	return b.inner.Len()
}

func (b *dequeSlice[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *dequeSlice[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *dequeSlice[T]) Clear() {
	b.inner.Clear()
}

func (b *dequeSlice[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *dequeSlice[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *dequeSlice[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *dequeSlice[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *dequeSlice[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *dequeSlice[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *dequeSlice[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *dequeSlice[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *dequeSlice[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *dequeSlice[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *dequeSlice[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *dequeSlice[T]) PushFront(element T) {
	b.inner.Insert(0, element)
}

func (b *dequeSlice[T]) PushBack(element T) {
	b.inner.Append(element)
}

func (b *dequeSlice[T]) PopFront() Option[T] {
	return b.inner.RemoveAt(0)
}

func (b *dequeSlice[T]) PopBack() Option[T] {
	return b.inner.RemoveAt(b.Len() - 1)
}

func (b *dequeSlice[T]) PeekFront() Option[T] {
	return b.inner.First()
}

func (b *dequeSlice[T]) PeekBack() Option[T] {
	return b.inner.Last()
}
