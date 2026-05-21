package array

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/sequence/builder"
)

// Builder is a [builder.BaseBuilder] for [collection.Array] implementations.
type Builder[T any, Target collection.Sequence[T], Self Builder[T, Target, Self]] interface {
	builder.BaseBuilder[T, Target, Self]
	// Size sets the fixed size of the array.
	Size(cap int) Self
}
