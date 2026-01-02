package orderedhashset

import (
	"codeberg.org/yaadata/bina/set"
	"codeberg.org/yaadata/bina/set/builder"
)

type Builder[T any, Target set.Set[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
}
