package btree

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/tree/builder"
)

// Builder is a [builder.BaseBuilder] for [collection.BTree] implementations.
type Builder[K any, V any, Target collection.BTree[K, V], Self any] interface {
	builder.BaseBuilder[K, V, Target, Self]
}
