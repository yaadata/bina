package orderedhashmap

import (
	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/collection"
	orderedhashmap "codeberg.org/yaadata/bina/internal/ordered_hashmap"
)

func BuiltinBuilder[K comparable, V any]() Builder[K, V, collection.OrderedMap[K, V], *build[K, V]] {
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

func (b *build[K, V]) Build() collection.OrderedMap[K, V] {
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
