package set

import (
	"iter"

	. "codeberg.org/yaadata/opt"
)

type OrderedSet[T any] interface {
	Set[T]
	All() iter.Seq2[int, T]
	First() Option[T]
	Last() Option[T]
}
