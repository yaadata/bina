package maps

import (
	"iter"

	. "github.com/yaadata/optionsgo"
)

type SortedMap[K any, V any] interface {
	OrderedMap[K, V]
	Range(from, to K) iter.Seq2[K, V]
	Floor(key K) Option[MapEntry[K, V]]
	Ceiling(key K) Option[MapEntry[K, V]]
}
