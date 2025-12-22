package maps

import (
	"iter"

	"codeberg.org/yaadata/bina/core/sequence"
	. "github.com/yaadata/optionsgo"
)

type OrderedMap[K any, V any] interface {
	Map[K, V]
	ToSequence() sequence.Sequence[K]
	First() Option[MapEntry[K, V]]
	Last() Option[MapEntry[K, V]]
	Reversed() iter.Seq2[K, V]
}
