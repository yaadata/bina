package collection

import (
	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
)

// OrderedMap is a [Map] that preserves insertion order.
type OrderedMap[K comparable, V any] interface {
	Map[K, V]
	// First returns the earliest inserted entry, or None if empty.
	First() Option[kv.Pair[K, V]]
	// Last returns the most recently inserted entry, or None if empty.
	Last() Option[kv.Pair[K, V]]
}
