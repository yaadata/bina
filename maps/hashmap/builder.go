package hashmap

import (
	"codeberg.org/yaadata/bina/maps"
	"codeberg.org/yaadata/bina/maps/builder"
)

type Builder[K comparable, V any, Target maps.Map[K, V], Self any] interface {
	builder.BaseBuilder[K, V, Target, Self]
}
