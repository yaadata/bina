package builder

import "codeberg.org/yaadata/bina/maps"

type BaseBuilder[K comparable, V any, Target maps.Map[K, V], Self any] interface {
	Build() Target
	Capacity(cap int) Self
	From(builtin map[K]V) Self
}
