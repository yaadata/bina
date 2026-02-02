package collection

import (
	"iter"

	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
)

// Map is a collection of key-value pairs with unique keys.
type Map[K comparable, V any] interface {
	Collection[K]
	Aggregate[kv.Pair[K, V]]
	// All returns an iterator over key-value pairs.
	All() iter.Seq2[K, V]
	// Delete removes the entry for the given key, returning the value or None if not found.
	Delete(key K) Option[V]
	// Get returns the value for the given key, or None if not found.
	Get(key K) Option[V]
	// Keys returns an iterator over all keys.
	Keys() iter.Seq[K]
	// Merge combines this map with another, using the provided function to resolve conflicts.
	Merge(other Map[K, V], fn MapMergeFunc[K, V]) Map[K, V]
	// Put inserts or updates a key-value pair, returning true if the key was newly inserted.
	Put(key K, value V) bool
	// Values returns an iterator over all values.
	Values() iter.Seq[V]
}
