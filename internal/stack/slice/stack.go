package stack

import (
	"iter"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type stack[T any] struct {
	inner sequence.Slice[T]
}

// Compile-time interface checks
func _[T comparable]() {
	var _ sequence.Stack[T] = (*stack[T])(nil)
}

func __[T compare.Comparable[T]]() {
	var _ sequence.Stack[T] = (*stack[T])(nil)
}

func StackFromBuiltin[T comparable](inner sequence.Slice[T]) *stack[T] {
	return &stack[T]{inner: inner}
}

func StackFromComparable[T compare.Comparable[T]](inner sequence.Slice[T]) *stack[T] {
	return &stack[T]{inner: inner}
}

func (b *stack[T]) Len() int {
	return b.inner.Len()
}

func (b *stack[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *stack[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *stack[T]) Clear() {
	b.inner.Clear()
}

func (b *stack[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *stack[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *stack[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *stack[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *stack[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *stack[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *stack[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *stack[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *stack[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *stack[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *stack[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *stack[T]) ToSlice() []T {
	return b.inner.ToSlice()
}

func (b *stack[T]) Push(element T) {
	b.inner.Append(element)
}

func (b *stack[T]) Pop() Option[T] {
	return b.inner.RemoveAt(b.Len() - 1)
}

func (b *stack[T]) Peek() Option[T] {
	return b.inner.Last()
}
