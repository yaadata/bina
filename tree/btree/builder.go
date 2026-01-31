package btree

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/tree/builder"
)

type Builder[K any, V any, Target collection.BTree[K, V], Self any] interface {
	builder.BaseBuilder[K, V, Target, Self]
}
