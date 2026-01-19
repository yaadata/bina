package orderedhashset

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/set/builder"
)

type Builder[T any, Target collection.Set[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
}
