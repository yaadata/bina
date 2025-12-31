package builder

import "codeberg.org/yaadata/bina/set"

type BaseBuilder[T any, Target set.Set[T], Self any] interface {
	Capacity(cap int) Self
	Build() Target
}
