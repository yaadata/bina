package hashset

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/set/builder"
)

// Builder defines the fluent interface for constructing hash sets.
// Use [NewBuiltinBuilder] for comparable types or [NewHashableBuilder]
// for types implementing the Hashable interface.
type Builder[T any, Target collection.Set[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
}
