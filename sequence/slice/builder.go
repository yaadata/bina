package slice

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/sequence/builder"
)

// Builder is a [builder.BaseBuilder] for [collection.Slice] implementations.
type Builder[T any, Target collection.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	// From initializes the slice with the given items.
	From(items ...T) Self
	// Capacity sets the initial capacity of the slice.
	Capacity(cap int) Self
}
