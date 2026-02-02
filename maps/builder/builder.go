package builder

import "codeberg.org/yaadata/bina/core/collection"

// BaseBuilder defines a fluent builder for [collection.Map] implementations.
// The Self type parameter enables method chaining.
type BaseBuilder[K comparable, V any, Target collection.Map[K, V], Self any] interface {
	// Build constructs and returns the target map.
	Build() Target
	// Capacity sets the initial capacity for the underlying map.
	Capacity(cap int) Self
	// From initializes the builder with entries from a built-in Go map.
	From(builtin map[K]V) Self
}
