package collection

import (
	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/core/where"
)

type Array[T any] interface {
	Sequence[T]
	Filter(predicate predicate.Predicate[T]) Array[T]
	First() Option[T]
	Last() Option[T]
	Offer(element T, index int) bool
	OfferRange(elements []T, cfgs ...where.WhereOption[int]) bool
}
