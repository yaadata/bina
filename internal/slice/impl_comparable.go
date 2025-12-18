package slice

import (
	"iter"
	"slices"
	"sort"

	. "github.com/yaadata/optionsgo"

	"github.com/yaadata/bina/core/compare"
	"github.com/yaadata/bina/core/sequence"
	"github.com/yaadata/bina/core/shared"
)

type sliceComparableInterface[T compare.Comparable[T]] struct {
	inner []T
}

func SliceFromComparableInterface[T compare.Comparable[T]](items ...T) *sliceComparableInterface[T] {
	return &sliceComparableInterface[T]{inner: items}
}

// Compile-time interface implementation check for sliceComparableInterface
func _[T compare.Comparable[T]]() {
	var _ sequence.Slice[T] = (*sliceComparableInterface[T])(nil)
}

func (s *sliceComparableInterface[T]) Len() int {
	return len(s.inner)
}

func (s *sliceComparableInterface[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *sliceComparableInterface[T]) Clear() {
	if !s.IsEmpty() {
		s.inner = make([]T, 0)
	}
}

func (s *sliceComparableInterface[T]) Contains(element T) bool {
	for _, item := range s.inner {
		if item.Equal(element) {
			return true
		}
	}
	return false
}

func (s *sliceComparableInterface[T]) Any(predicate shared.Predicate[T]) bool {
	return slices.ContainsFunc(s.inner, predicate)
}

func (s *sliceComparableInterface[T]) Count(predicate shared.Predicate[T]) int {
	var count int
	for _, item := range s.inner {
		if predicate(item) {
			count++
		}
	}
	return count
}

func (s *sliceComparableInterface[T]) Every(predicate shared.Predicate[T]) bool {
	for _, item := range s.inner {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (s *sliceComparableInterface[T]) ForEach(fn func(T)) {
	for _, item := range s.inner {
		fn(item)
	}
}

func (s *sliceComparableInterface[T]) Append(item T) {
	s.inner = append(s.inner, item)
}

func (s *sliceComparableInterface[T]) All() iter.Seq[T] {
	return func(yield func(item T) bool) {
		for _, item := range s.inner {
			if !yield(item) {
				return
			}
		}
	}
}

func (s *sliceComparableInterface[T]) Enumerate() iter.Seq2[int, T] {
	return func(yield func(index int, item T) bool) {
		for index, item := range s.inner {
			if !yield(index, item) {
				return
			}
		}
	}
}

func (s *sliceComparableInterface[T]) Extend(items ...T) {
	s.inner = append(s.inner, items...)
}

func (s *sliceComparableInterface[T]) ExtendFromSequence(sequence sequence.Sequence[T]) {
	s.inner = append(s.inner, sequence.ToSlice()...)
}

func (s *sliceComparableInterface[T]) Last() Option[T] {
	length := len(s.inner)
	if length == 0 {
		return None[T]()
	}
	return Some(s.inner[length-1])
}

func (s *sliceComparableInterface[T]) Filter(predicate shared.Predicate[T]) sequence.Slice[T] {
	filtered := make([]T, 0, len(s.inner))
	for _, item := range s.inner {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return &sliceComparableInterface[T]{inner: filtered}
}

func (s *sliceComparableInterface[T]) Find(predicate shared.Predicate[T]) Option[T] {
	for _, item := range s.inner {
		if predicate(item) {
			Some(item)
		}
	}
	return None[T]()
}

func (s *sliceComparableInterface[T]) FindIndex(predicate shared.Predicate[T]) Option[int] {
	for index, item := range s.inner {
		if predicate(item) {
			Some(index)
		}
	}
	return None[int]()
}

func (s *sliceComparableInterface[T]) First() Option[T] {
	if len(s.inner) == 0 {
		return None[T]()
	}
	return Some(s.inner[0])
}

func (s *sliceComparableInterface[T]) Get(index int) Option[T] {
	length := len(s.inner)
	if index < 0 || index >= length {
		return None[T]()
	}
	return Some(s.inner[index])
}

func (s *sliceComparableInterface[T]) Insert(index int, item T) {
	s.inner = append(s.inner[:index], append([]T{item}, s.inner[index:]...)...)
}

func (s *sliceComparableInterface[T]) RemoveAt(index int) T {
	var item T
	s.inner, item = slices.Delete(s.inner, index, index+1), s.inner[index]
	return item
}

func (s *sliceComparableInterface[T]) Retain(predicate shared.Predicate[T]) {
	var retained = make([]T, 0, len(s.inner))
	for _, item := range s.inner {
		if predicate(item) {
			retained = append(retained, item)
		}
	}
	s.inner = retained
}

func (s *sliceComparableInterface[T]) Sort(fn func(a, b T) compare.Order) {
	sort.SliceStable(s.inner, func(i, j int) bool {
		a := s.inner[i]
		b := s.inner[j]
		return fn(a, b).IsLess()
	})
}

func (s *sliceComparableInterface[T]) ToSlice() []T {
	return s.inner
}
