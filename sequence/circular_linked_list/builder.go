package circularlinkedlist

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/sequence/builder"
)

// Builder is a [builder.BaseBuilder] for [collection.LinkedList] implementations.
type Builder[T any, Target collection.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	// From initializes the linked list with the given items.
	From(items ...T) Self
}
