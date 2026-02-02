package builder

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/kv"
)

// BaseBuilder defines a fluent builder for [collection.SearchTree] implementations.
// The Self type parameter enables method chaining.
type BaseBuilder[K any, V any, Target collection.SearchTree[K, V], Self any] interface {
	// Build constructs and returns the target search tree.
	Build() Target
	// Order sets the branching factor of the tree. Default is 5.
	Order(order int) Self
	// From initializes the tree with the given key-value pairs.
	From(pairs ...kv.Pair[K, V]) Self
}
