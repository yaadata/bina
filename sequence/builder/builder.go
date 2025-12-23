package builder

import (
	"codeberg.org/yaadata/bina/sequence"
)

type Builder[T any, Target sequence.Sequence[T]] interface {
	From(items ...T) Builder[T, Target]
	Capacity(cap int) Builder[T, Target]
	Build() Target
}
