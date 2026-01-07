package stack

import (
	"iter"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type builtinStack[T comparable] struct {
	inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]
}

func _[T comparable]() {
	var _ sequence.Stack[T] = (*builtinStack[T])(nil)
}

func StackFromBuiltin[T comparable](inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]) *builtinStack[T] {
	return &builtinStack[T]{inner: inner}
}

func (b *builtinStack[T]) Len() int {
	return b.inner.Len()
}

func (b *builtinStack[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *builtinStack[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *builtinStack[T]) Clear() {
	b.inner.Clear()
}

func (b *builtinStack[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *builtinStack[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *builtinStack[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *builtinStack[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *builtinStack[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *builtinStack[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *builtinStack[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *builtinStack[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *builtinStack[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *builtinStack[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *builtinStack[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *builtinStack[T]) ToSlice() []T {
	return b.inner.ToSlice()
}

func (b *builtinStack[T]) Push(element T) {
	b.inner.Append(element)
}

func (b *builtinStack[T]) Pop() Option[T] {
	return b.inner.RemoveAt(b.Len() - 1)
}

func (b *builtinStack[T]) Peek() Option[T] {
	tail := b.inner.Tail()
	if tail.IsNone() {
		return None[T]()
	}
	return Some(tail.Unwrap().Value())
}
