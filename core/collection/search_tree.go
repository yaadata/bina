package collection

import (
	"iter"

	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/kv"
	"codeberg.org/yaadata/bina/core/where"
)

type SearchTree[K any, V any] interface {
	Collection[K]
	Aggregate[K]
	Delete(key K) Option[V]
	Get(key K) Option[V]
	Put(key K, value V)
	Height() int
	Min() Option[kv.Pair[K, V]]
	Max() Option[kv.Pair[K, V]]
	Floor(key K) Option[kv.Pair[K, V]]
	Ceiling(key K) Option[kv.Pair[K, V]]
	Range(opts ...where.WhereOption[K]) iter.Seq2[K, V]
	All(opts ...SearchTreeTraversalOption) iter.Seq2[K, V]
}
