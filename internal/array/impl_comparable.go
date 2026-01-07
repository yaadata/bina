package array

import (
	"iter"
	"reflect"
	"slices"
	"sort"

	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	core_range "codeberg.org/yaadata/bina/core/range"
	"codeberg.org/yaadata/bina/sequence"
)

type arrayComparableInterface[T compare.Comparable[T]] struct {
	inner []T
}

func ArrayFromComparableInterface[T compare.Comparable[T]](size int) *arrayComparableInterface[T] {
	return &arrayComparableInterface[T]{
		inner: make([]T, size),
	}
}

// Compile-time interface implementation check for arrayComparableInterface
func _[T compare.Comparable[T]]() {
	var _ sequence.Array[T] = (*arrayComparableInterface[T])(nil)
}

func (s *arrayComparableInterface[T]) Len() int {
	return len(s.inner)
}

func (s *arrayComparableInterface[T]) IsEmpty() bool {
	return s.Every(func(item T) bool {
		return reflect.ValueOf(item).IsZero()
	})
}

func (s *arrayComparableInterface[T]) Clear() {
	clear(s.inner)
}

func (s *arrayComparableInterface[T]) Contains(element T) bool {
	for _, item := range s.inner {
		if item.Equal(element) {
			return true
		}
	}
	return false
}

func (s *arrayComparableInterface[T]) Any(predicate predicate.Predicate[T]) bool {
	return slices.ContainsFunc(s.inner, predicate)
}

func (s *arrayComparableInterface[T]) Count(predicate predicate.Predicate[T]) int {
	var count int
	for _, item := range s.inner {
		if predicate(item) {
			count++
		}
	}
	return count
}

func (s *arrayComparableInterface[T]) Every(predicate predicate.Predicate[T]) bool {
	for _, item := range s.inner {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (s *arrayComparableInterface[T]) ForEach(fn func(T)) {
	for _, item := range s.inner {
		fn(item)
	}
}

func (s *arrayComparableInterface[T]) All() iter.Seq[T] {
	return func(yield func(item T) bool) {
		for _, item := range s.inner {
			if !yield(item) {
				return
			}
		}
	}
}

func (s *arrayComparableInterface[T]) Enumerate() iter.Seq2[int, T] {
	return func(yield func(index int, item T) bool) {
		for index, item := range s.inner {
			if !yield(index, item) {
				return
			}
		}
	}
}

func (s *arrayComparableInterface[T]) Last() Option[T] {
	length := len(s.inner)
	if length == 0 {
		return None[T]()
	}
	return Some(s.inner[length-1])
}

func (s *arrayComparableInterface[T]) Filter(predicate predicate.Predicate[T]) sequence.Array[T] {
	filtered := make([]T, 0, len(s.inner))
	for _, item := range s.inner {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return &arrayComparableInterface[T]{
		inner: filtered,
	}
}

func (s *arrayComparableInterface[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	for _, item := range s.inner {
		if predicate(item) {
			return Some(item)
		}
	}
	return None[T]()
}

func (s *arrayComparableInterface[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	for index, item := range s.inner {
		if predicate(item) {
			return Some(index)
		}
	}
	return None[int]()
}

func (s *arrayComparableInterface[T]) First() Option[T] {
	if len(s.inner) == 0 {
		return None[T]()
	}
	return Some(s.inner[0])
}

func (s *arrayComparableInterface[T]) Get(index int) Option[T] {
	length := len(s.inner)
	if index < 0 || index >= length {
		return None[T]()
	}
	return Some(s.inner[index])
}

func (s *arrayComparableInterface[T]) Offer(element T, index int) bool {
	if index < 0 || index >= s.Len() {
		return false
	}
	s.inner[index] = element
	return true
}

func (s *arrayComparableInterface[T]) OfferRange(elements []T, cfgs ...core_range.CoreRangeConfig) bool {
	r := core_range.New()
	for _, cfg := range cfgs {
		cfg(r)
	}
	length := s.Len()
	from := r.From().UnwrapOrDefault()
	if from < 0 {
		return false
	}
	end := r.End().UnwrapOrElse(func() int {
		return from + len(elements)
	})
	if end > length {
		return false
	}
	slices.Replace(s.inner, from, end, elements...)
	return true
}

func (s *arrayComparableInterface[T]) Retain(predicate predicate.Predicate[T]) {
	for index, element := range s.Enumerate() {
		if !predicate(element) {
			s.inner[index] = *new(T)
		}
	}
}

func (s *arrayComparableInterface[T]) Sort(fn func(a, b T) compare.Order) {
	sort.SliceStable(s.inner, func(i, j int) bool {
		a := s.inner[i]
		b := s.inner[j]
		return fn(a, b).IsLess()
	})
}

func (s *arrayComparableInterface[T]) ToSlice() []T {
	return s.inner
}
