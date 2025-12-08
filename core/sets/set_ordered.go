package sets

import (
	"iter"

	"github.com/yaadata/bina/core/sequential"
	. "github.com/yaadata/optionsgo"
)

type OrderedSet[T any] interface {
	Set[T]
	ToSequence() sequential.Sequence[T]
	First() Option[T]
	Last() Option[T]
	Reversed() iter.Seq[T]
}
