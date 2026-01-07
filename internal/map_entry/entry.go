package mapentry

import "codeberg.org/yaadata/bina/maps"

type entry[K any, V any] struct {
	key   K
	value V
}

func New[K comparable, V any](key K, value V) maps.MapEntry[K, V] {
	return &entry[K, V]{
		key:   key,
		value: value,
	}
}

func (e *entry[K, V]) Key() K {
	return e.key
}

func (e *entry[K, V]) Value() V {
	return e.value
}
