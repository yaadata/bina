package maps

import (
	. "github.com/yaadata/optionsgo"
)

type OrderedMap[K any, V any] interface {
	Map[K, V]
	First() Option[MapEntry[K, V]]
	Last() Option[MapEntry[K, V]]
}
