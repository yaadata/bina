package orderedhashmap

import (
	"iter"

	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/kv"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/internal/hashmap"
	bina_maps "codeberg.org/yaadata/bina/maps"
)

type orderedHashMapFromBuiltin[K comparable, V any] struct {
	ordered  []kv.Pair[K, V]
	deleted  []bool
	keyIndex map[K]int
	size     int
}

// compile time check
var _ bina_maps.OrderedMap[int, int] = (*orderedHashMapFromBuiltin[int, int])(nil)

func OrderedHashMapFromBuiltin[K comparable, V any](capacity int) bina_maps.OrderedMap[K, V] {
	return &orderedHashMapFromBuiltin[K, V]{
		ordered:  make([]kv.Pair[K, V], 0, capacity),
		deleted:  make([]bool, 0, capacity/2),
		keyIndex: make(map[K]int, capacity),
		size:     0,
	}
}

func (m *orderedHashMapFromBuiltin[K, V]) Len() int {
	return m.size
}

func (m *orderedHashMapFromBuiltin[K, V]) Contains(key K) bool {
	_, contains := m.keyIndex[key]
	return contains
}

func (m *orderedHashMapFromBuiltin[K, V]) IsEmpty() bool {
	return m.size == 0
}

func (m *orderedHashMapFromBuiltin[K, V]) Clear() {
	clear(m.keyIndex)
	m.size = 0
	m.ordered = make([]kv.Pair[K, V], 0)
	m.deleted = make([]bool, 0)
}

func (m *orderedHashMapFromBuiltin[K, V]) Any(pred predicate.Predicate[kv.Pair[K, V]]) bool {
	for key, value := range m.All() {
		if pred(kv.New(key, value)) {
			return true
		}
	}
	return false
}

func (m *orderedHashMapFromBuiltin[K, V]) Count(pred predicate.Predicate[kv.Pair[K, V]]) int {
	var count int
	for key, value := range m.All() {
		if pred(kv.New(key, value)) {
			count++
		}
	}
	return count
}

func (m *orderedHashMapFromBuiltin[K, V]) Every(pred predicate.Predicate[kv.Pair[K, V]]) bool {
	for key, value := range m.All() {
		if !pred(kv.New(key, value)) {
			return false
		}
	}
	return true
}

func (m *orderedHashMapFromBuiltin[K, V]) ForEach(fn func(element kv.Pair[K, V])) {
	for key, value := range m.All() {
		fn(kv.New(key, value))
	}
}

func (m *orderedHashMapFromBuiltin[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for index, entry := range m.ordered {
			if m.deleted[index] {
				continue
			}
			if !yield(entry.Key(), entry.Value()) {
				return
			}
		}
	}
}

func (m *orderedHashMapFromBuiltin[K, V]) Delete(key K) Option[V] {
	index, contains := m.keyIndex[key]
	if !contains {
		return None[V]()
	}
	value := m.ordered[index].Value()
	m.deleted[index] = true
	delete(m.keyIndex, key)
	m.size--
	if m.size >= len(m.keyIndex)/2 {
		m.compact()
	}
	return Some(value)
}

func (m *orderedHashMapFromBuiltin[K, V]) Get(key K) Option[V] {
	index, contains := m.keyIndex[key]
	if !contains {
		return None[V]()
	}
	return Some(m.ordered[index].Value())
}

func (m *orderedHashMapFromBuiltin[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for index, entry := range m.ordered {
			if m.deleted[index] {
				continue
			}
			if !yield(entry.Key()) {
				return
			}
		}
	}
}

func (m *orderedHashMapFromBuiltin[K, V]) Merge(other bina_maps.Map[K, V], fn bina_maps.MapMergeFunc[K, V]) bina_maps.Map[K, V] {
	res := hashmap.New(make(map[K]V, m.Len()+other.Len()))
	for key, value := range m.All() {
		res.Put(key, value)
	}
	for key, incoming := range other.All() {
		current := res.Get(key)
		if !current.IsSome() {
			res.Put(key, incoming)
		} else {
			res.Put(key, fn(key, current.Unwrap(), incoming))
		}
	}
	return res
}

func (m *orderedHashMapFromBuiltin[K, V]) Put(key K, value V) bool {
	if index, contains := m.keyIndex[key]; contains {
		m.ordered[index] = kv.New(key, value)
		return false
	}
	m.keyIndex[key] = len(m.ordered)
	m.ordered = append(m.ordered, kv.New(key, value))
	m.deleted = append(m.deleted, false)
	m.size++
	return true
}

func (m *orderedHashMapFromBuiltin[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for index, entry := range m.ordered {
			if m.deleted[index] {
				continue
			}
			if !yield(entry.Value()) {
				return
			}
		}
	}
}

func (m *orderedHashMapFromBuiltin[K, V]) compact() {
	updated := make([]kv.Pair[K, V], 0, m.size)
	for index, entry := range m.ordered {
		if !m.deleted[index] {
			m.keyIndex[entry.Key()] = len(updated)
			updated = append(updated, entry)
		}
	}
	m.ordered = updated
	m.deleted = make([]bool, len(updated))
}

func (m *orderedHashMapFromBuiltin[K, V]) First() Option[kv.Pair[K, V]] {
	for index, entry := range m.ordered {
		if !m.deleted[index] {
			return Some(entry)
		}
	}
	return None[kv.Pair[K, V]]()
}

func (m *orderedHashMapFromBuiltin[K, V]) Last() Option[kv.Pair[K, V]] {
	for i := len(m.ordered) - 1; i >= 0; i-- {
		if !m.deleted[i] {
			return Some(m.ordered[i])
		}
	}
	return None[kv.Pair[K, V]]()
}
