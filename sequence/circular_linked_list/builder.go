package circularlinkedlist

import (
	"codeberg.org/yaadata/bina/sequence"
	"codeberg.org/yaadata/bina/sequence/builder"
)

type Builder[T any, Target sequence.Sequence[T], Self any] interface {
	builder.BaseBuilder[T, Target, Self]
	From(items ...T) Self
}
