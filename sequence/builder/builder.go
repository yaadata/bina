package builder

import "codeberg.org/yaadata/bina/core/collection"

type BaseBuilder[T any, Target collection.Sequence[T], Self any] interface {
	Build() Target
}
