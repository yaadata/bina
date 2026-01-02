package maps

import (
	"iter"

	"codeberg.org/yaadata/bina/core/collection"
	. "github.com/yaadata/optionsgo"
)

type Map[K any, V any] interface {
	collection.Collection[K]
	collection.Aggregate[K]
	Get(key K) Option[V]
	Put(key K, value V) bool
	Delete(key K) bool
	Enumerate() iter.Seq2[K, V]
	Keys() iter.Seq[K]
	values() iter.Seq[V]
}
