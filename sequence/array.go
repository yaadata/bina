package sequence

import (
	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/predicate"
	core_range "codeberg.org/yaadata/bina/core/range"
)

type Array[T any] interface {
	Sequence[T]
	InsertRange(elements []T, opt ...core_range.CoreRangeConfig) bool
	Filter(predicate predicate.Predicate[T]) Array[T]
	First() Option[T]
	Last() Option[T]
}
