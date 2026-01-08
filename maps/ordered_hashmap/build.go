package orderedhashmap

import (
	. "github.com/yaadata/optionsgo"

	orderedhashmap "codeberg.org/yaadata/bina/internal/ordered_hashmap"
	"codeberg.org/yaadata/bina/maps"
)

func BuiltinBuilder[K comparable, V any]() Builder[K, V, maps.OrderedMap[K, V], *build[K, V]] {
	return &build[K, V]{
		capacity: None[int](),
		from:     None[map[K]V](),
	}
}

type build[K comparable, V any] struct {
	capacity Option[int]
	from     Option[map[K]V]
}

func (b *build[K, V]) Capacity(capacity int) *build[K, V] {
	b.capacity = Some(capacity)
	return b
}

func (b *build[K, V]) From(builtin map[K]V) *build[K, V] {
	b.from = Some(builtin)
	return b
}

func (b *build[K, V]) Build() maps.OrderedMap[K, V] {
	if b.from.IsNone() {
		return orderedhashmap.OrderedHashMapFromBuiltin[K, V](b.capacity.UnwrapOrDefault())
	}
	from := b.from.Unwrap()
	m := orderedhashmap.OrderedHashMapFromBuiltin[K, V](max(b.capacity.UnwrapOrDefault(), len(from)))
	for k, v := range from {
		m.Put(k, v)
	}
	return m
}
