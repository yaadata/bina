package maps

import (
	"iter"

	"github.com/yaadata/bina/core/sequential"
	. "github.com/yaadata/optionsgo"
)

type OrderedMap[K any, V any] interface {
	Map[K, V]
	ToSequence() sequential.Sequence[K]
	First() Option[MapEntry[K, V]]
	Last() Option[MapEntry[K, V]]
	Reversed() iter.Seq2[K, V]
}
