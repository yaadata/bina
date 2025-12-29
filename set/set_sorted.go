package set

import (
	"iter"

	. "github.com/yaadata/optionsgo"
)

type SortedSet[T any] interface {
	OrderedSet[T]
	Range(from, to T) iter.Seq[T]
	Floor(value T) Option[T]
	Ceiling(value T) Option[T]
}
