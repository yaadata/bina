package orderedhashset

import (
	"iter"

	"codeberg.org/yaadata/bina/core/hashable"
	"codeberg.org/yaadata/bina/core/predicate"
	hashset "codeberg.org/yaadata/bina/internal/hash_set"
	"codeberg.org/yaadata/bina/set"
	. "github.com/yaadata/optionsgo"
)

type orderedHashSetFromHashable[K comparable, T hashable.Hashable[K]] struct {
	ordered []T
	deleted []bool
	set     map[K]int
	size    int
}

func _[K comparable, T hashable.Hashable[K]]() {
	var _ set.OrderedSet[T] = (*orderedHashSetFromHashable[K, T])(nil)
}

func OrderedHashSetFromHashable[K comparable, T hashable.Hashable[K]](capacity int) set.Set[T] {
	return &orderedHashSetFromHashable[K, T]{
		ordered: make([]T, 0, capacity),
		deleted: make([]bool, capacity/2),
		set:     make(map[K]int, capacity),
		size:    0,
	}
}

func (s *orderedHashSetFromHashable[K, T]) Len() int {
	return s.size
}

func (s *orderedHashSetFromHashable[K, T]) Contains(element T) bool {
	_, contains := s.set[element.Hash()]
	return contains
}

func (s *orderedHashSetFromHashable[K, T]) IsEmpty() bool {
	return s.size == 0
}

func (s *orderedHashSetFromHashable[K, T]) Clear() {
	clear(s.set)
	s.size = 0
	s.ordered = make([]T, 0)
	s.deleted = make([]bool, 0)
}

func (s *orderedHashSetFromHashable[K, T]) Any(pred predicate.Predicate[T]) bool {
	for element := range s.All() {
		if pred(element) {
			return true
		}
	}
	return false
}

func (s *orderedHashSetFromHashable[K, T]) Count(pred predicate.Predicate[T]) int {
	var count int
	for element := range s.All() {
		if pred(element) {
			count++
		}
	}
	return count
}

func (s *orderedHashSetFromHashable[K, T]) Every(pred predicate.Predicate[T]) bool {
	for element := range s.All() {
		if !pred(element) {
			return false
		}
	}
	return true
}

func (s *orderedHashSetFromHashable[K, T]) ForEach(fn func(element T)) {
	for element := range s.All() {
		fn(element)
	}
}

func (s *orderedHashSetFromHashable[K, T]) Add(element T) bool {
	if s.Contains(element) {
		return false
	}
	s.set[element.Hash()] = len(s.set)
	s.ordered = append(s.ordered, element)
	s.deleted = append(s.deleted, false)
	s.size++
	return true
}

func (s *orderedHashSetFromHashable[K, T]) All() iter.Seq[T] {
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

func (s *orderedHashSetFromHashable[K, T]) Difference(other set.Set[T]) Option[set.Set[T]] {
	var res = hashset.HashSetFromHashable[K, T](s.Len())
	for element := range s.All() {
		if !other.Contains(element) {
			res.Add(element)
		}
	}

	if res.IsEmpty() {
		return None[set.Set[T]]()
	}
	return Some(res)
}

func (s *orderedHashSetFromHashable[K, T]) Enumerate() iter.Seq2[int, T] {
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

func (s *orderedHashSetFromHashable[K, T]) Extend(elements ...T) {
	for _, element := range elements {
		_ = s.Add(element)
	}
}

func (s *orderedHashSetFromHashable[K, T]) Intersect(other set.Set[T]) Option[set.Set[T]] {
	res := hashset.HashSetFromHashable[K, T](s.Len())
	for element := range s.All() {
		if other.Contains(element) {
			res.Add(element)
		}
	}
	if res.IsEmpty() {
		return None[set.Set[T]]()
	}
	return Some(res)
}

func (s *orderedHashSetFromHashable[K, T]) IsSubsetOf(other set.Set[T]) bool {
	if s.Len() > other.Len() {
		return false
	}
	for element := range s.All() {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

func (s *orderedHashSetFromHashable[K, T]) IsSupersetOf(other set.Set[T]) bool {
	if s.Len() < other.Len() {
		return false
	}

	for element := range other.All() {
		if !s.Contains(element) {
			return false
		}
	}
	return true
}

func (s *orderedHashSetFromHashable[K, T]) Remove(element T) bool {
	index, contains := s.set[element.Hash()]
	if !contains {
		return false
	}
	s.deleted[index] = true
	delete(s.set, element.Hash())
	s.size--
	if s.size >= len(s.set)/2 {
		s.Compact()
	}
	return true
}

func (s *orderedHashSetFromHashable[K, T]) SymmetricDifference(other set.Set[T]) Option[set.Set[T]] {
	var res = hashset.HashSetFromHashable[K, T](s.Len())
	for element := range s.All() {
		if !other.Contains(element) {
			res.Add(element)
		}
	}

	for element := range other.All() {
		if !s.Contains(element) {
			res.Add(element)
		}
	}

	if res.IsEmpty() {
		return None[set.Set[T]]()
	}
	return Some(res)
}

func (s *orderedHashSetFromHashable[K, T]) Union(other set.Set[T]) set.Set[T] {
	var res = hashset.HashSetFromHashable[K, T](s.Len() + other.Len())
	return res
}

func (s *orderedHashSetFromHashable[K, T]) Compact() {
	updatedOrder := make([]T, 0, s.size)
	for index, element := range s.ordered {
		if !s.deleted[index] {
			s.set[element.Hash()] = len(updatedOrder)
			updatedOrder = append(updatedOrder, element)
		}
	}
	s.ordered = updatedOrder
	s.deleted = make([]bool, len(updatedOrder))
}

func (s *orderedHashSetFromHashable[K, T]) First() Option[T] {
	if s.Len() > 0 {
		return Some(s.ordered[0])
	}
	return None[T]()
}

func (s *orderedHashSetFromHashable[K, T]) Last() Option[T] {
	l := s.Len()
	if l > 0 {
		return Some(s.ordered[l-1])
	}
	return None[T]()
}
