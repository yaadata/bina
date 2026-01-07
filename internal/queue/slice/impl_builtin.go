package queue

import (
	"iter"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type builtinQueue[T comparable] struct {
	inner sequence.Slice[T]
}

func _[T comparable]() {
	var _ sequence.Queue[T] = (*builtinQueue[T])(nil)
}

func QueueFromBuiltin[T comparable](inner sequence.Slice[T]) *builtinQueue[T] {
	return &builtinQueue[T]{inner: inner}
}

func (b *builtinQueue[T]) Len() int {
	return b.inner.Len()
}

func (b *builtinQueue[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *builtinQueue[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *builtinQueue[T]) Clear() {
	b.inner.Clear()
}

func (b *builtinQueue[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *builtinQueue[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *builtinQueue[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *builtinQueue[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *builtinQueue[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *builtinQueue[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *builtinQueue[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *builtinQueue[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *builtinQueue[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *builtinQueue[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *builtinQueue[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *builtinQueue[T]) ToSlice() []T {
	return b.inner.ToSlice()
}

func (b *builtinQueue[T]) Enqueue(element T) {
	b.inner.Append(element)
}

func (b *builtinQueue[T]) Dequeue() Option[T] {
	return b.inner.RemoveAt(0)
}

func (b *builtinQueue[T]) Peek() Option[T] {
	return b.inner.First()
}
