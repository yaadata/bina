package maps

import (
	"iter"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
)

type Map[K comparable, V any] interface {
	collection.Collection[K]
	collection.Aggregate[kv.Pair[K, V]]
	All() iter.Seq2[K, V]
	Delete(key K) Option[V]
	Get(key K) Option[V]
	Keys() iter.Seq[K]
	Merge(other Map[K, V], fn MapMergeFunc[K, V]) Map[K, V]
	Put(key K, value V) bool
	Values() iter.Seq[V]
}
