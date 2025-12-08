package slice

import (
	"iter"
	"slices"

	"github.com/yaadata/bina/core/compare"
	"github.com/yaadata/bina/core/sequential"
	"github.com/yaadata/bina/core/shared"
	. "github.com/yaadata/optionsgo"
)

type sliceCompareInterface[T compare.Comparable[T]] struct {
	_data []T
}

func NewSliceCompareInterface[T compare.Compare[T]](items ...T) sequential.Sequence[T] {
	return &sliceCompareInterface[T]{_data: items}
}

func (s *sliceCompareInterface[T]) Len() int {
	return len(s._data)
}

func (s *sliceCompareInterface[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *sliceCompareInterface[T]) Clear() {
	if !s.IsEmpty() {
		s._data = make([]T, 0)
	}
}

func (s *sliceCompareInterface[T]) Contains(element T) bool {
	for _, item := range s._data {
		if item.Compare(element).IsEqual() {
			return true
		}
	}
	return false
}

func (s *sliceCompareInterface[T]) Any(predicate shared.Predicate[T]) bool {
	return slices.ContainsFunc(s._data, predicate)
}

func (s *sliceCompareInterface[T]) Count(predicate shared.Predicate[T]) int {
	var count int
	for _, item := range s._data {
		if predicate(item) {
			count++
		}
	}
	return count
}

func (s *sliceCompareInterface[T]) Every(predicate shared.Predicate[T]) bool {
	for _, item := range s._data {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (s *sliceCompareInterface[T]) ForEach(fn func(T)) {
	for _, item := range s._data {
		fn(item)
	}
}

func (s *sliceCompareInterface[T]) Append(item T) sequential.Sequence[T] {
	s._data = append(s._data, item)
	return s
}

func (s *sliceCompareInterface[T]) All() iter.Seq[T] {
	return func(yield func(item T) bool) {
		for _, item := range s._data {
			if !yield(item) {
				return
			}
		}
	}
}

func (s *sliceCompareInterface[T]) Enumerate() iter.Seq2[int, T] {
	return func(yield func(index int, item T) bool) {
		for index, item := range s._data {
			if !yield(index, item) {
				return
			}
		}
	}
}

func (s *sliceCompareInterface[T]) Extend(items ...T) sequential.Sequence[T] {
	s._data = append(s._data, items...)
	return s
}

func (s *sliceCompareInterface[T]) ExtendFromSequence(sequence sequential.Sequence[T]) sequential.Sequence[T] {
	s._data = append(s._data, sequence.ToSlice()...)
	return s
}

func (s *sliceCompareInterface[T]) Last() Option[T] {
	length := len(s._data)
	if length == 0 {
		return None[T]()
	}
	return Some(s._data[length-1])
}

func (s *sliceCompareInterface[T]) Filter(predicate shared.Predicate[T]) sequential.Sequence[T] {
	filtered := make([]T, 0, len(s._data))
	for _, item := range s._data {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return &sliceCompareInterface[T]{_data: filtered}
}

func (s *sliceCompareInterface[T]) Find(predicate shared.Predicate[T]) Option[T] {
	for _, item := range s._data {
		if predicate(item) {
			Some(item)
		}
	}
	return None[T]()
}

func (s *sliceCompareInterface[T]) FindIndex(predicate shared.Predicate[T]) Option[int] {
	for index, item := range s._data {
		if predicate(item) {
			Some(index)
		}
	}
	return None[int]()
}

func (s *sliceCompareInterface[T]) First() Option[T] {
	if len(s._data) == 0 {
		return None[T]()
	}
	return Some(s._data[0])
}

func (s *sliceCompareInterface[T]) Get(index int) Option[T] {
	length := len(s._data)
	if index < 0 || index >= length {
		return None[T]()
	}
	return Some(s._data[index])
}

func (s *sliceCompareInterface[T]) Insert(index int, item T) sequential.Sequence[T] {
	s._data = append(s._data[:index], append([]T{item}, s._data[index:]...)...)
	return s
}

func (s *sliceCompareInterface[T]) RemoveAt(index int) T {
	var item T
	s._data, item = slices.Delete(s._data, index, index+1), s._data[index]
	return item
}

func (s *sliceCompareInterface[T]) Retain(predicate shared.Predicate[T]) sequential.Sequence[T] {
	var retained = make([]T, 0, len(s._data))
	for _, item := range s._data {
		if predicate(item) {
			retained = append(retained, item)
		}
	}
	s._data = retained
	return s
}

func (s *sliceCompareInterface[T]) ToSlice() []T {
	return s._data
}
