package stack

import (
	"iter"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

// Slice-backed stack

type sliceStack[T any] struct {
	inner sequence.Slice[T]
}

func _[T comparable]() {
	var _ sequence.Stack[T] = (*sliceStack[T])(nil)
}

func __[T compare.Comparable[T]]() {
	var _ sequence.Stack[T] = (*sliceStack[T])(nil)
}

func SliceStackFromBuiltin[T comparable](inner sequence.Slice[T]) *sliceStack[T] {
	return &sliceStack[T]{inner: inner}
}

func SliceStackFromComparable[T compare.Comparable[T]](inner sequence.Slice[T]) *sliceStack[T] {
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
	return b.inner.ToSlice()
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

// LinkedList-backed stack

type linkedListStack[T any] struct {
	inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]
}

func ___[T comparable]() {
	var _ sequence.Stack[T] = (*linkedListStack[T])(nil)
}

func ____[T compare.Comparable[T]]() {
	var _ sequence.Stack[T] = (*linkedListStack[T])(nil)
}

func LinkedListStackFromBuiltin[T comparable](inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]) *linkedListStack[T] {
	return &linkedListStack[T]{inner: inner}
}

func LinkedListStackFromComparable[T compare.Comparable[T]](inner sequence.LinkedList[T, sequence.SinglyLinkedListNode[T]]) *linkedListStack[T] {
	return &linkedListStack[T]{inner: inner}
}

func (b *linkedListStack[T]) Len() int {
	return b.inner.Len()
}

func (b *linkedListStack[T]) Contains(value T) bool {
	return b.inner.Contains(value)
}

func (b *linkedListStack[T]) IsEmpty() bool {
	return b.inner.IsEmpty()
}

func (b *linkedListStack[T]) Clear() {
	b.inner.Clear()
}

func (b *linkedListStack[T]) Any(predicate predicate.Predicate[T]) bool {
	return b.inner.Any(predicate)
}

func (b *linkedListStack[T]) Count(predicate predicate.Predicate[T]) int {
	return b.inner.Count(predicate)
}

func (b *linkedListStack[T]) Every(predicate predicate.Predicate[T]) bool {
	return b.inner.Every(predicate)
}

func (b *linkedListStack[T]) ForEach(fn func(T)) {
	b.inner.ForEach(fn)
}

func (b *linkedListStack[T]) All() iter.Seq2[int, T] {
	return b.inner.All()
}

func (b *linkedListStack[T]) Values() iter.Seq[T] {
	return b.inner.Values()
}

func (b *linkedListStack[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	return b.inner.Find(predicate)
}

func (b *linkedListStack[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	return b.inner.FindIndex(predicate)
}

func (b *linkedListStack[T]) Get(index int) Option[T] {
	return b.inner.Get(index)
}

func (b *linkedListStack[T]) Retain(predicate predicate.Predicate[T]) {
	b.inner.Retain(predicate)
}

func (b *linkedListStack[T]) Sort(fn func(a, b T) compare.Order) {
	b.inner.Sort(fn)
}

func (b *linkedListStack[T]) ToSlice() []T {
	return b.inner.ToSlice()
}

func (b *linkedListStack[T]) Push(element T) {
	b.inner.Append(element)
}

func (b *linkedListStack[T]) Pop() Option[T] {
	return b.inner.RemoveAt(b.Len() - 1)
}

func (b *linkedListStack[T]) Peek() Option[T] {
	tail := b.inner.Tail()
	if tail.IsNone() {
		return None[T]()
	}
	return Some(tail.Unwrap().Value())
}
