package slice

import (
	"github.com/yaadata/bina/core/compare"
	"github.com/yaadata/bina/core/sequential"
	. "github.com/yaadata/optionsgo"
)

type Builder[T any] interface {
	From(items ...T) Builder[T]
	Capacity(cap int) Builder[T]
	Build() sequential.Sequence[T]
}

func NewBuiltinBuilder[T comparable]() Builder[T] {
	return &builtinBuilder[T]{
		from:     None[[]T](),
		capacity: None[int](),
	}
}

func NewComparableBuilder[T compare.Comparable[T]]() Builder[T] {
	return &comparableBuilder[T]{
		from:     None[[]T](),
		capacity: None[int](),
	}
}
