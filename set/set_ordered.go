package sets

import (
	"iter"

	"codeberg.org/yaadata/bina/sequence"
	. "github.com/yaadata/optionsgo"
)

type OrderedSet[T any] interface {
	Set[T]
	ToSequence() sequence.Sequence[T]
	First() Option[T]
	Last() Option[T]
	Reversed() iter.Seq[T]
}
