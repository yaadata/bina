package orderedhashmap

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/maps/builder"
)

// Builder is a [builder.BaseBuilder] for [collection.OrderedMap] implementations.
type Builder[K comparable, V any, Target collection.OrderedMap[K, V], Self any] interface {
	builder.BaseBuilder[K, V, Target, Self]
}
