package builder

import "codeberg.org/yaadata/bina/core/collection"

type BaseBuilder[K comparable, V any, Target collection.Map[K, V], Self any] interface {
	Build() Target
	Capacity(cap int) Self
	From(builtin map[K]V) Self
}
