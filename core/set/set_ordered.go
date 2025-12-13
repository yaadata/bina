package sets

import (
	"iter"

	"github.com/yaadata/bina/core/sequence"
	. "github.com/yaadata/optionsgo"
)

type OrderedSet[T any] interface {
	Set[T]
	ToSequence() sequence.Sequence[T]
	First() Option[T]
	Last() Option[T]
	Reversed() iter.Seq[T]
}
