package hashset

import (
	"iter"
	"maps"

	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/set"
	. "github.com/yaadata/optionsgo"
)

type hashSetFromBuiltin[T comparable] struct {
	set map[T]bool
}

func _[T comparable]() {
	var _ set.Set[T] = (*hashSetFromBuiltin[T])(nil)
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
	for element := range s.All() {
		if pred(element) {
			return true
		}
	}
	return false
}

func (s *hashSetFromBuiltin[T]) Count(pred predicate.Predicate[T]) int {
	var count int
	for element := range s.All() {
		if pred(element) {
			count++
		}
	}
	return count
}

func (s *hashSetFromBuiltin[T]) Every(pred predicate.Predicate[T]) bool {
	for element := range s.All() {
		if !pred(element) {
			return false
		}
	}
	return true
}

func (s *hashSetFromBuiltin[T]) ForEach(fn func(element T)) {
	for element := range s.All() {
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

func (s *hashSetFromBuiltin[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for key := range s.set {
			if !yield(key) {
				return
			}
		}
	}
}

func (s *hashSetFromBuiltin[T]) Difference(other set.Set[T]) Option[set.Set[T]] {
	return None[set.Set[T]]()
}

func (s *hashSetFromBuiltin[T]) Remove(element T) bool {
	if s.Contains(element) {
		delete(s.set, element)
		return true
	}
	return false
}

func (s *hashSetFromBuiltin[T]) Intersect(other set.Set[T]) Option[set.Set[T]] {
	m := make(map[T]bool)
	for element := range other.All() {
		if s.Contains(element) {
			m[element] = true
		}
	}
	if len(m) == 0 {
		return None[set.Set[T]]()
	}
	var resp set.Set[T] = &hashSetFromBuiltin[T]{
		set: m,
	}
	return Some(resp)
}

func (s *hashSetFromBuiltin[T]) IsSubsetOf(other set.Set[T]) bool {
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

func (s *hashSetFromBuiltin[T]) IsSupersetOf(other set.Set[T]) bool {
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

func (s *hashSetFromBuiltin[T]) SymmetricDifference(other set.Set[T]) Option[set.Set[T]] {
	var m = make(map[T]bool)
	for element := range s.All() {
		if !other.Contains(element) {
			m[element] = true
		}
	}

	for element := range other.All() {
		if !s.Contains(element) {
			m[element] = true
		}
	}

	if len(m) == 0 {
		return None[set.Set[T]]()
	}
	var resp set.Set[T] = &hashSetFromBuiltin[T]{
		set: m,
	}
	return Some(resp)
}

func (s *hashSetFromBuiltin[T]) Union(other set.Set[T]) set.Set[T] {
	var resp = maps.Clone(s.set)
	for element := range other.All() {
		resp[element] = true
	}
	return &hashSetFromBuiltin[T]{
		set: resp,
	}
}
