package maps

import (
	. "github.com/yaadata/optionsgo"
)

type OrderedMap[K comparable, V any] interface {
	Map[K, V]
	First() Option[MapEntry[K, V]]
	Last() Option[MapEntry[K, V]]
}
