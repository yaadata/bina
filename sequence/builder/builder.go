package builder

import "codeberg.org/yaadata/bina/core/collection"

// BaseBuilder defines a fluent builder for [collection.Sequence] implementations.
type BaseBuilder[T any, Target collection.Sequence[T], Self BaseBuilder[T, Target, Self]] interface {
	// Build constructs and returns the target sequence.
	Build() Target
}
