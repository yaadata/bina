package kv

type Pair[K any, V any] interface {
	Key() K
	Value() V
}
