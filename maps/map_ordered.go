package maps

import (
	"codeberg.org/yaadata/bina/core/kv"
	. "github.com/yaadata/optionsgo"
)

type OrderedMap[K comparable, V any] interface {
	Map[K, V]
	First() Option[kv.Pair[K, V]]
	Last() Option[kv.Pair[K, V]]
}
