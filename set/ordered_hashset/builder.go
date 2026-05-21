package orderedhashset

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/set/builder"
)

// Builder defines the fluent interface for constructing ordered hash sets.
// Use [NewBuiltinBuilder] for comparable types or [NewHashableBuilder]
// for types implementing the Hashable interface.
type Builder[T any, Target collection.Set[T], Self Builder[T, Target, Self]] interface {
	builder.BaseBuilder[T, Target, Self]
}
