package builder

import "codeberg.org/yaadata/bina/core/collection"

type BaseBuilder[T any, Target collection.Set[T], Self any] interface {
	Build() Target
	Capacity(cap int) Self
	From(values ...T) Self
}
