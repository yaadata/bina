package collection

import "github.com/yaadata/bina/core/shared"

type Aggregate[T any] interface {
	Any(predicate shared.Predicate[T]) bool
	Count(predicate shared.Predicate[T]) int
	Every(predicate shared.Predicate[T]) bool
	ForEach(fn func(T))
}
