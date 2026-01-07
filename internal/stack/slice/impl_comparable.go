package stack

import (
	"iter"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type comparableStack[T compare.Comparable[T]] struct {
	inner sequence.Slice[T]
}

func _[T compare.Comparable[T]]() {
	var _ sequence.Stack[T] = (*comparableStack[T])(nil)
}

func StackFromComparable[T compare.Comparable[T]](inner sequence.Slice[T]) *comparableStack[T] {
	return &comparableStack[T]{inner: inner}
}

func (b *comparableStack[T]) Len() int {
	return b.inner.Len()
}

func (b *comparableStack[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *comparableStack[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *comparableStack[T]) Clear() {
	b.inner.Clear()
}

func (b *comparableStack[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *comparableStack[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *comparableStack[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *comparableStack[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *comparableStack[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *comparableStack[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *comparableStack[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *comparableStack[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *comparableStack[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *comparableStack[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *comparableStack[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *comparableStack[T]) ToSlice() []T {
	return b.inner.ToSlice()
}

func (b *comparableStack[T]) Push(element T) {
	b.inner.Append(element)
}

func (b *comparableStack[T]) Pop() Option[T] {
	return b.inner.RemoveAt(b.Len() - 1)
}

func (b *comparableStack[T]) Peek() Option[T] {
	return b.inner.Last()
}
