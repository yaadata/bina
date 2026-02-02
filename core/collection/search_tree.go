package collection

import (
	"iter"

	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/kv"
	"codeberg.org/yaadata/bina/core/where"
)

// SearchTree is an ordered key-value collection that maintains keys in sorted order.
type SearchTree[K any, V any] interface {
	Collection[K]
	Aggregate[kv.Pair[K, V]]

	// Delete removes the entry for the given key, returning the value or None if not found.
	Delete(key K) Option[V]

	// Get returns the value for the given key, or None if not found.
	Get(key K) Option[V]

	// Put inserts or updates a key-value pair.
	Put(key K, value V)

	// Height returns the height of the tree.
	Height() int

	// Min returns the entry with the smallest key, or None if empty.
	Min() Option[kv.Pair[K, V]]

	// Max returns the entry with the largest key, or None if empty.
	Max() Option[kv.Pair[K, V]]

	// Floor returns the entry with the largest key less than or equal to the given key, or None if no such entry exists.
	Floor(key K) Option[kv.Pair[K, V]]

	// Ceiling returns the entry with the smallest key greater than or equal to the given key, or None if no such entry exists.
	Ceiling(key K) Option[kv.Pair[K, V]]

	// Range returns an iterator over entries within the specified key bounds.
	Range(opts ...where.WhereOption[K]) iter.Seq2[K, V]

	// All returns an iterator over all entries using the specified traversal strategy.
	All(opts ...SearchTreeTraversalOption) iter.Seq2[K, V]
}
