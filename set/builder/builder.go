package builder

import "codeberg.org/yaadata/bina/core/collection"

// BaseBuilder defines the fluent interface for constructing sets.
// Methods return Self to enable method chaining.
type BaseBuilder[T any, Target collection.Set[T], Self any] interface {
	// Build constructs and returns the set with the configured options.
	// If From was called, the set is populated with those values.
	// If Capacity was called, the set is pre-allocated to that size.
	Build() Target

	// Capacity sets the initial capacity hint for the underlying storage.
	// Returns Self for method chaining.
	Capacity(cap int) Self

	// From specifies the initial values to populate the set with.
	// Duplicate values are automatically removed.
	// Returns Self for method chaining.
	From(values ...T) Self
}
