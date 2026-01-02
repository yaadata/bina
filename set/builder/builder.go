package builder

import "codeberg.org/yaadata/bina/set"

type BaseBuilder[T any, Target set.Set[T], Self any] interface {
	Build() Target
	Capacity(cap int) Self
	From(values ...T) Self
}
