package hashmap

import (
	"iter"
	"maps"

	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/kv"
	"codeberg.org/yaadata/bina/core/predicate"
	bina_maps "codeberg.org/yaadata/bina/maps"
)

// compile time check
var _ bina_maps.Map[int, int] = (*impl[int, int])(nil)

type impl[K comparable, V any] struct {
	m map[K]V
}

func New[K comparable, V any](m map[K]V) *impl[K, V] {
	return &impl[K, V]{
		m: m,
	}
}

func (i *impl[K, V]) Len() int {
	return len(i.m)
}

func (i *impl[K, V]) Contains(element K) bool {
	_, ok := i.m[element]
	return ok
}

func (i *impl[K, V]) IsEmpty() bool {
	return len(i.m) == 0
}

func (i *impl[K, V]) Clear() {
	clear(i.m)
}

func (i *impl[K, V]) Any(predicate predicate.Predicate[kv.Pair[K, V]]) bool {
	for key, value := range i.All() {
		if predicate(kv.New(key, value)) {
			return true
		}
	}
	return false
}

func (i *impl[K, V]) Count(predicate predicate.Predicate[kv.Pair[K, V]]) int {
	var count int
	for key, value := range i.All() {
		if predicate(kv.New(key, value)) {
			count++
		}
	}
	return count
}

func (i *impl[K, V]) Every(predicate predicate.Predicate[kv.Pair[K, V]]) bool {
	for key, value := range i.All() {
		if !predicate(kv.New(key, value)) {
			return false
		}
	}
	return true
}

func (i *impl[K, V]) ForEach(fn func(element kv.Pair[K, V])) {
	for key, value := range i.All() {
		fn(kv.New(key, value))
	}
}

func (i *impl[K, V]) All() iter.Seq2[K, V] {
	return maps.All(i.m)
}

func (i *impl[K, V]) Delete(key K) Option[V] {
	value, ok := i.m[key]
	if !ok {
		return None[V]()
	}
	delete(i.m, key)
	return Some(value)
}

func (i *impl[K, V]) Get(key K) Option[V] {
	value, ok := i.m[key]
	if !ok {
		return None[V]()
	}
	return Some(value)
}

func (i *impl[K, V]) Keys() iter.Seq[K] {
	return maps.Keys(i.m)
}

func (i *impl[K, V]) Merge(other bina_maps.Map[K, V], fn bina_maps.MapMergeFunc[K, V]) bina_maps.Map[K, V] {
	res := New(make(map[K]V, i.Len()+other.Len()))
	for key, incoming := range other.All() {
		current := i.Get(key)
		if !current.IsSome() {
			res.m[key] = incoming
		} else {
			res.m[key] = fn(key, current.Unwrap(), incoming)
		}
	}
	return res
}

func (i *impl[K, V]) Put(key K, value V) bool {
	i.m[key] = value
	return true
}

func (i *impl[K, V]) Values() iter.Seq[V] {
	return maps.Values(i.m)
}
