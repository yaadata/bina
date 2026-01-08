package hashset

import (
	"iter"
	"maps"

	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/set"
	. "codeberg.org/yaadata/opt"
)

type hashSetFromBuiltin[T comparable] struct {
	set map[T]bool
}

func _[T comparable]() {
	var _ set.Set[T] = (*hashSetFromBuiltin[T])(nil)
}

func HashSetFromBuiltin[T comparable](capacity int) set.Set[T] {
	return &hashSetFromBuiltin[T]{
		set: make(map[T]bool, capacity),
	}
}

func (s *hashSetFromBuiltin[T]) Len() int {
	return len(s.set)
}

func (s *hashSetFromBuiltin[T]) Contains(element T) bool {
	return s.set[element]
}

func (s *hashSetFromBuiltin[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *hashSetFromBuiltin[T]) Clear() {
	clear(s.set)
}

func (s *hashSetFromBuiltin[T]) Any(pred predicate.Predicate[T]) bool {
	for element := range s.Values() {
		if pred(element) {
			return true
		}
	}
	return false
}

func (s *hashSetFromBuiltin[T]) Count(pred predicate.Predicate[T]) int {
	var count int
	for element := range s.Values() {
		if pred(element) {
			count++
		}
	}
	return count
}

func (s *hashSetFromBuiltin[T]) Every(pred predicate.Predicate[T]) bool {
	for element := range s.Values() {
		if !pred(element) {
			return false
		}
	}
	return true
}

func (s *hashSetFromBuiltin[T]) ForEach(fn func(element T)) {
	for element := range s.Values() {
		fn(element)
	}
}

func (s *hashSetFromBuiltin[T]) Add(element T) bool {
	if s.Contains(element) {
		return false
	}
	s.set[element] = true
	return true
}

func (s *hashSetFromBuiltin[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for key := range s.set {
			if !yield(key) {
				return
			}
		}
	}
}

func (s *hashSetFromBuiltin[T]) Difference(other set.Set[T]) Option[set.Set[T]] {
	res := HashSetFromBuiltin[T](s.Len())
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

func (s *hashSetFromBuiltin[T]) Extend(values ...T) {
	for _, value := range values {
		s.set[value] = true
	}
}

func (s *hashSetFromBuiltin[T]) Intersect(other set.Set[T]) Option[set.Set[T]] {
	res := HashSetFromBuiltin[T](s.Len())
	for element := range s.Values() {
		if other.Contains(element) {
			res.Add(element)
		}
	}
	if res.IsEmpty() {
		return None[set.Set[T]]()
	}
	return Some(res)
}

func (s *hashSetFromBuiltin[T]) IsSubsetOf(other set.Set[T]) bool {
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

func (s *hashSetFromBuiltin[T]) IsSupersetOf(other set.Set[T]) bool {
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

func (s *hashSetFromBuiltin[T]) Remove(element T) bool {
	if s.Contains(element) {
		delete(s.set, element)
		return true
	}
	return false
}

func (s *hashSetFromBuiltin[T]) SymmetricDifference(other set.Set[T]) Option[set.Set[T]] {
	var res = HashSetFromBuiltin[T](s.Len())
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
		return None[set.Set[T]]()
	}
	return Some(res)
}

func (s *hashSetFromBuiltin[T]) Union(other set.Set[T]) set.Set[T] {
	var resp = maps.Clone(s.set)
	for element := range other.Values() {
		resp[element] = true
	}
	return &hashSetFromBuiltin[T]{
		set: resp,
	}
}
