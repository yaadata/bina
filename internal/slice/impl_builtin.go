package slice

import (
	"iter"
	"slices"
	"sort"

	. "github.com/yaadata/optionsgo"

	"github.com/yaadata/bina/core/compare"
	"github.com/yaadata/bina/core/sequential"
	"github.com/yaadata/bina/core/shared"
)

type sliceFromBuiltin[T comparable] struct {
	inner []T
}

var _ sequential.Sequence[int] = (*sliceFromBuiltin[int])(nil)

func SliceFromBuiltin[T comparable](items ...T) *sliceFromBuiltin[T] {
	return &sliceFromBuiltin[T]{inner: items}
}

func (s *sliceFromBuiltin[T]) Len() int {
	return len(s.inner)
}

func (s *sliceFromBuiltin[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *sliceFromBuiltin[T]) Clear() {
	if !s.IsEmpty() {
		s.inner = make([]T, 0)
	}
}

func (s *sliceFromBuiltin[T]) Contains(element T) bool {
	return slices.Contains(s.inner, element)
}

func (s *sliceFromBuiltin[T]) Any(predicate shared.Predicate[T]) bool {
	return slices.ContainsFunc(s.inner, predicate)
}

func (s *sliceFromBuiltin[T]) Count(predicate shared.Predicate[T]) int {
	var count int
	for _, item := range s.inner {
		if predicate(item) {
			count++
		}
	}
	return count
}

func (s *sliceFromBuiltin[T]) Every(predicate shared.Predicate[T]) bool {
	for _, item := range s.inner {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (s *sliceFromBuiltin[T]) ForEach(fn func(T)) {
	for _, item := range s.inner {
		fn(item)
	}
}

func (s *sliceFromBuiltin[T]) Append(item T) sequential.Sequence[T] {
	s.inner = append(s.inner, item)
	return s
}

func (s *sliceFromBuiltin[T]) All() iter.Seq[T] {
	return func(yield func(item T) bool) {
		for _, item := range s.inner {
			if !yield(item) {
				return
			}
		}
	}
}

func (s *sliceFromBuiltin[T]) Enumerate() iter.Seq2[int, T] {
	return func(yield func(index int, item T) bool) {
		for index, item := range s.inner {
			if !yield(index, item) {
				return
			}
		}
	}
}

func (s *sliceFromBuiltin[T]) Extend(items ...T) sequential.Sequence[T] {
	s.inner = append(s.inner, items...)
	return s
}

func (s *sliceFromBuiltin[T]) ExtendFromSequence(sequence sequential.Sequence[T]) sequential.Sequence[T] {
	s.inner = append(s.inner, sequence.ToSlice()...)
	return s
}

func (s *sliceFromBuiltin[T]) Last() Option[T] {
	length := len(s.inner)
	if length == 0 {
		return None[T]()
	}
	return Some(s.inner[length-1])
}

func (s *sliceFromBuiltin[T]) Filter(predicate shared.Predicate[T]) sequential.Sequence[T] {
	filtered := make([]T, 0, len(s.inner))
	for _, item := range s.inner {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return &sliceFromBuiltin[T]{inner: filtered}
}

func (s *sliceFromBuiltin[T]) Find(predicate shared.Predicate[T]) Option[T] {
	for _, item := range s.inner {
		if predicate(item) {
			return Some(item)
		}
	}
	return None[T]()
}

func (s *sliceFromBuiltin[T]) FindIndex(predicate shared.Predicate[T]) Option[int] {
	for index, item := range s.inner {
		if predicate(item) {
			return Some(index)
		}
	}
	return None[int]()
}

func (s *sliceFromBuiltin[T]) First() Option[T] {
	if len(s.inner) == 0 {
		return None[T]()
	}
	return Some(s.inner[0])
}

func (s *sliceFromBuiltin[T]) Get(index int) Option[T] {
	length := len(s.inner)
	if index < 0 || index >= length {
		return None[T]()
	}
	return Some(s.inner[index])
}

func (s *sliceFromBuiltin[T]) Insert(index int, item T) sequential.Sequence[T] {
	s.inner = append(s.inner[:index], append([]T{item}, s.inner[index:]...)...)
	return s
}

func (s *sliceFromBuiltin[T]) RemoveAt(index int) T {
	var item T
	s.inner, item = slices.Delete(s.inner, index, index+1), s.inner[index]
	return item
}

func (s *sliceFromBuiltin[T]) Retain(predicate shared.Predicate[T]) sequential.Sequence[T] {
	var retained = make([]T, 0, len(s.inner))
	for _, item := range s.inner {
		if predicate(item) {
			retained = append(retained, item)
		}
	}
	s.inner = retained
	return s
}

func (s *sliceFromBuiltin[T]) Sort(fn func(a, b T) compare.Order) sequential.Sequence[T] {
	sort.SliceStable(s.inner, func(i, j int) bool {
		a := s.inner[i]
		b := s.inner[j]
		return fn(a, b).IsLess()
	})
	return s
}

func (s *sliceFromBuiltin[T]) ToSlice() []T {
	return s.inner
}
