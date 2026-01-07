package set

import (
	"iter"

	. "github.com/yaadata/optionsgo"
)

type OrderedSet[T any] interface {
	Set[T]
	All() iter.Seq2[int, T]
	First() Option[T]
	Last() Option[T]
	AsSlice() []T
}
