package hashmap

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/maps/builder"
)

type Builder[K comparable, V any, Target collection.Map[K, V], Self any] interface {
	builder.BaseBuilder[K, V, Target, Self]
}
