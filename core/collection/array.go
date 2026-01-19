package collection

import (
	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/predicate"
	core_range "codeberg.org/yaadata/bina/core/range"
)

type Array[T any] interface {
	Sequence[T]
	Filter(predicate predicate.Predicate[T]) Array[T]
	First() Option[T]
	Last() Option[T]
	Offer(element T, index int) bool
	OfferRange(elements []T, cfgs ...core_range.RangeConfig[int]) bool
}
