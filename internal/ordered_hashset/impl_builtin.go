package orderedhashset

import (
	"iter"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/predicate"
	hashset "codeberg.org/yaadata/bina/internal/hashset"
	. "codeberg.org/yaadata/opt"
)

type orderedHashSetFromBuiltin[T comparable] struct {
	ordered []T
	deleted []bool
	set     map[T]int
	size    int
}

func _[T comparable]() {
	var _ collection.OrderedSet[T] = (*orderedHashSetFromBuiltin[T])(nil)
}

func OrderedHashSetFromBuiltin[T comparable](capacity int) collection.OrderedSet[T] {
	return &orderedHashSetFromBuiltin[T]{
		ordered: make([]T, 0, capacity),
		deleted: make([]bool, capacity/2),
		set:     make(map[T]int, capacity),
		size:    0,
	}
}

func (s *orderedHashSetFromBuiltin[T]) Len() int {
	return s.size
}

func (s *orderedHashSetFromBuiltin[T]) Contains(element T) bool {
	_, contains := s.set[element]
	return contains
}

func (s *orderedHashSetFromBuiltin[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *orderedHashSetFromBuiltin[T]) Clear() {
	clear(s.set)
	s.size = 0
	s.ordered = make([]T, 0)
	s.deleted = make([]bool, 0)
}

func (s *orderedHashSetFromBuiltin[T]) Any(pred predicate.Predicate[T]) bool {
	for element := range s.Values() {
		if pred(element) {
			return true
		}
	}
	return false
}

func (s *orderedHashSetFromBuiltin[T]) Count(pred predicate.Predicate[T]) int {
	var count int
	for element := range s.Values() {
		if pred(element) {
			count++
		}
	}
	return count
}

func (s *orderedHashSetFromBuiltin[T]) Every(pred predicate.Predicate[T]) bool {
	for element := range s.Values() {
		if !pred(element) {
			return false
		}
	}
	return true
}

func (s *orderedHashSetFromBuiltin[T]) ForEach(fn func(element T)) {
	for element := range s.Values() {
		fn(element)
	}
}

func (s *orderedHashSetFromBuiltin[T]) Add(element T) bool {
	if s.Contains(element) {
		return false
	}
	s.set[element] = len(s.set)
	s.ordered = append(s.ordered, element)
	s.deleted = append(s.deleted, false)
	s.size++
	return true
}

func (s *orderedHashSetFromBuiltin[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for index, key := range s.ordered {
			if s.deleted[index] {
				continue
			}
			if !yield(key) {
				return
			}
		}
	}
}

func (s *orderedHashSetFromBuiltin[T]) Difference(other collection.Set[T]) Option[collection.Set[T]] {
	var res = hashset.HashSetFromBuiltin[T](s.Len())
	for element := range s.Values() {
		if !other.Contains(element) {
			res.Add(element)
		}
	}

	if res.IsEmpty() {
		return None[collection.Set[T]]()
	}
	return Some(res)
}

func (s *orderedHashSetFromBuiltin[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		trueIndex := 0
		for index, key := range s.ordered {
			if s.deleted[index] {
				continue
			}
			if !yield(trueIndex, key) {
				return
			}
			trueIndex++
		}
	}
}

func (s *orderedHashSetFromBuiltin[T]) Extend(elements ...T) {
	for _, element := range elements {
		_ = s.Add(element)
	}
}

func (s *orderedHashSetFromBuiltin[T]) Intersect(other collection.Set[T]) Option[collection.Set[T]] {
	res := hashset.HashSetFromBuiltin[T](s.Len())
	for element := range s.Values() {
		if other.Contains(element) {
			res.Add(element)
		}
	}
	if res.IsEmpty() {
		return None[collection.Set[T]]()
	}
	return Some(res)
}

func (s *orderedHashSetFromBuiltin[T]) IsSubsetOf(other collection.Set[T]) bool {
	if s.Len() > other.Len() {
		return false
	}
	for element := range s.Values() {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

func (s *orderedHashSetFromBuiltin[T]) IsSupersetOf(other collection.Set[T]) bool {
	if s.Len() < other.Len() {
		return false
	}

	for element := range other.Values() {
		if !s.Contains(element) {
			return false
		}
	}
	return true
}

func (s *orderedHashSetFromBuiltin[T]) Remove(element T) bool {
	index, contains := s.set[element]
	if !contains {
		return false
	}
	s.deleted[index] = true
	delete(s.set, element)
	s.size--
	if s.size >= len(s.set)/2 {
		s.compact()
	}
	return true
}

func (s *orderedHashSetFromBuiltin[T]) SymmetricDifference(other collection.Set[T]) Option[collection.Set[T]] {
	var res = hashset.HashSetFromBuiltin[T](s.Len())
	for element := range s.Values() {
		if !other.Contains(element) {
			res.Add(element)
		}
	}

	for element := range other.Values() {
		if !s.Contains(element) {
			res.Add(element)
		}
	}

	if res.IsEmpty() {
		return None[collection.Set[T]]()
	}
	return Some(res)
}

func (s *orderedHashSetFromBuiltin[T]) Union(other collection.Set[T]) collection.Set[T] {
	var res = hashset.HashSetFromBuiltin[T](s.Len() + other.Len())
	for element := range s.Values() {
		res.Add(element)
	}
	for element := range other.Values() {
		res.Add(element)
	}
	return res
}

func (s *orderedHashSetFromBuiltin[T]) compact() {
	updatedOrder := make([]T, 0, s.size)
	for index, element := range s.ordered {
		if !s.deleted[index] {
			s.set[element] = len(updatedOrder)
			updatedOrder = append(updatedOrder, element)
		}
	}
	s.ordered = updatedOrder
	s.deleted = make([]bool, len(updatedOrder))
}

func (s *orderedHashSetFromBuiltin[T]) First() Option[T] {
	if s.Len() > 0 {
		return Some(s.ordered[0])
	}
	return None[T]()
}

func (s *orderedHashSetFromBuiltin[T]) Last() Option[T] {
	l := s.Len()
	if l > 0 {
		return Some(s.ordered[l-1])
	}
	return None[T]()
}

func (s *orderedHashSetFromBuiltin[T]) AsSlice() []T {
	res := make([]T, 0, s.size)
	for element := range s.Values() {
		res = append(res, element)
	}
	return res
}
