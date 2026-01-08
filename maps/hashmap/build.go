package hashmap

import (
	maps0 "maps"

	. "github.com/yaadata/optionsgo"

	hashmap "codeberg.org/yaadata/bina/internal/hashmap"
	"codeberg.org/yaadata/bina/maps"
)

func BuiltinBuilder[K comparable, V any]() Builder[K, V, maps.Map[K, V], *build[K, V]] {
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

func (b *build[K, V]) Build() maps.Map[K, V] {
	if b.from.IsNone() {
		return hashmap.New(make(map[K]V))
	}
	from := b.from.Unwrap()
	m := make(map[K]V, max(b.capacity.UnwrapOrDefault(), len(from)))
	maps0.Copy(m, b.from.Unwrap())
	return hashmap.New(m)
}
