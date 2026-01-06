package builder

import (
	"codeberg.org/yaadata/bina/sequence"
)

type BaseBuilder[T any, Target sequence.Sequence[T], Self any] interface {
	Build() Target
}
