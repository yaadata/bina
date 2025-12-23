package collection

import (
	"codeberg.org/yaadata/bina/core/predicate"
)

type Aggregate[T any] interface {
	Any(predicate predicate.Predicate[T]) bool
	Count(predicate predicate.Predicate[T]) int
	Every(predicate predicate.Predicate[T]) bool
	ForEach(fn func(T))
}
