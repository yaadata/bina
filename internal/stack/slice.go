package stack

import (
	"iter"
	"slices"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	. "codeberg.org/yaadata/opt"
)

// Slice-backed stack

type sliceStack[T any] struct {
	inner collection.Slice[T]
}

func _[T comparable]() {
	var _ collection.Stack[T] = (*sliceStack[T])(nil)
}

func __[T compare.Comparable[T]]() {
	var _ collection.Stack[T] = (*sliceStack[T])(nil)
}

func SliceStackFromBuiltin[T comparable](inner collection.Slice[T]) *sliceStack[T] {
	return &sliceStack[T]{inner: inner}
}

func SliceStackFromComparable[T compare.Comparable[T]](inner collection.Slice[T]) *sliceStack[T] {
	return &sliceStack[T]{inner: inner}
}

func (b *sliceStack[T]) Len() int {
	return b.inner.Len()
}

func (b *sliceStack[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *sliceStack[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *sliceStack[T]) Clear() {
	b.inner.Clear()
}

func (b *sliceStack[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *sliceStack[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *sliceStack[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *sliceStack[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *sliceStack[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *sliceStack[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *sliceStack[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *sliceStack[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *sliceStack[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *sliceStack[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *sliceStack[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *sliceStack[T]) ToSlice() []T {
	return slices.Collect(b.inner.Values())
}

func (b *sliceStack[T]) Push(element T) {
	b.inner.Append(element)
}

func (b *sliceStack[T]) Pop() Option[T] {
	return b.inner.RemoveAt(b.Len() - 1)
}

func (b *sliceStack[T]) Peek() Option[T] {
	return b.inner.Last()
}
