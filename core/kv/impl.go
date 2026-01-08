package kv

type pair[K any, V any] struct {
	key   K
	value V
}

func New[K comparable, V any](key K, value V) Pair[K, V] {
	return &pair[K, V]{
		key:   key,
		value: value,
	}
}

func (e *pair[K, V]) Key() K {
	return e.key
}

func (e *pair[K, V]) Value() V {
	return e.value
}
