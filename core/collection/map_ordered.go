package collection

import (
	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
)

type OrderedMap[K comparable, V any] interface {
	Map[K, V]
	First() Option[kv.Pair[K, V]]
	Last() Option[kv.Pair[K, V]]
}
