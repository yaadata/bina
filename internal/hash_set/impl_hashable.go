package hashset

import (
	"iter"
	"maps"

	"codeberg.org/yaadata/bina/core/hashable"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/set"
	. "github.com/yaadata/optionsgo"
)

type hashSetFromHashable[K comparable, T hashable.Hashable[K]] struct {
	set map[K]T
}

func HashSetFromHashable[K comparable, T hashable.Hashable[K]](capacity int) set.Set[T] {
	return &hashSetFromHashable[K, T]{
		set: make(map[K]T, capacity),
	}
}

func (s *hashSetFromHashable[K, T]) Len() int {
	return len(s.set)
}

func (s *hashSetFromHashable[K, T]) Contains(element T) bool {
	_, ok := s.set[element.Hash()]
	return ok
}

func (s *hashSetFromHashable[K, T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *hashSetFromHashable[K, T]) Clear() {
	clear(s.set)
}

func (s *hashSetFromHashable[K, T]) Any(pred predicate.Predicate[T]) bool {
	for element := range s.Values() {
		if pred(element) {
			return true
		}
	}
	return false
}

func (s *hashSetFromHashable[K, T]) Count(pred predicate.Predicate[T]) int {
	var count int
	for element := range s.Values() {
		if pred(element) {
			count++
		}
	}
	return count
}

func (s *hashSetFromHashable[K, T]) Every(pred predicate.Predicate[T]) bool {
	for element := range s.Values() {
		if !pred(element) {
			return false
		}
	}
	return true
}

func (s *hashSetFromHashable[K, T]) ForEach(fn func(element T)) {
	for element := range s.Values() {
		fn(element)
	}
}

func (s *hashSetFromHashable[K, T]) Add(element T) bool {
	if s.Contains(element) {
		return false
	}
	s.set[element.Hash()] = element
	return true
}

func (s *hashSetFromHashable[K, T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, value := range s.set {
			if !yield(value) {
				return
			}
		}
	}
}

func (s *hashSetFromHashable[K, T]) Difference(other set.Set[T]) Option[set.Set[T]] {
	res := HashSetFromHashable[K, T](s.Len())
	for element := range s.Values() {
		if !other.Contains(element) {
			res.Add(element)
		}
	}
	if res.IsEmpty() {
		return None[set.Set[T]]()
	}
	return Some(res)
}

func (s *hashSetFromHashable[K, T]) Extend(values ...T) {
	for _, value := range values {
		s.set[value.Hash()] = value
	}
}

func (s *hashSetFromHashable[K, T]) Intersect(other set.Set[T]) Option[set.Set[T]] {
	m := make(map[K]T)
	for element := range other.Values() {
		if s.Contains(element) {
			m[element.Hash()] = element
		}
	}
	if len(m) == 0 {
		return None[set.Set[T]]()
	}
	var resp set.Set[T] = &hashSetFromHashable[K, T]{
		set: m,
	}
	return Some(resp)
}

func (s *hashSetFromHashable[K, T]) IsSubsetOf(other set.Set[T]) bool {
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

func (s *hashSetFromHashable[K, T]) IsSupersetOf(other set.Set[T]) bool {
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

func (s *hashSetFromHashable[K, T]) Remove(element T) bool {
	if s.Contains(element) {
		delete(s.set, element.Hash())
		return true
	}
	return false
}

func (s *hashSetFromHashable[K, T]) SymmetricDifference(other set.Set[T]) Option[set.Set[T]] {
	var m = make(map[K]T)
	for element := range s.Values() {
		if !other.Contains(element) {
			m[element.Hash()] = element
		}
	}

	for element := range other.Values() {
		if !s.Contains(element) {
			m[element.Hash()] = element
		}
	}

	if len(m) == 0 {
		return None[set.Set[T]]()
	}
	var resp set.Set[T] = &hashSetFromHashable[K, T]{
		set: m,
	}
	return Some(resp)
}

func (s *hashSetFromHashable[K, T]) Union(other set.Set[T]) set.Set[T] {
	var resp = maps.Clone(s.set)
	for element := range other.Values() {
		resp[element.Hash()] = element
	}
	return &hashSetFromHashable[K, T]{
		set: resp,
	}
}
