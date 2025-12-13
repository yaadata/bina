package slice

import (
	. "github.com/yaadata/optionsgo"
	"github.com/yaadata/optionsgo/core"

	"github.com/yaadata/bina/core/compare"
	"github.com/yaadata/bina/core/sequence"
	internal "github.com/yaadata/bina/internal/slice"
	"github.com/yaadata/bina/sequence/builder"
)

func NewBuiltinBuilder[T comparable]() builder.Builder[T] {
	return &builtinBuilder[T]{
		from:     None[[]T](),
		capacity: None[int](),
	}
}

type builtinBuilder[T comparable] struct {
	from     Option[[]T]
	capacity Option[int]
}

func (b *builtinBuilder[T]) From(items ...T) builder.Builder[T] {
	b.from = Some(items)
	return b
}

func (b *builtinBuilder[T]) Capacity(cap int) builder.Builder[T] {
	b.capacity = Some(cap)
	return b
}

func (b *builtinBuilder[T]) Build() sequence.Sequence[T] {
	return internal.SliceFromBuiltin[T](b.from.OrElse(func() core.Option[[]T] {
		return b.from.Or(Some(make([]T, 0, b.capacity.UnwrapOrDefault())))
	}).Unwrap()...)
}

func NewComparableBuilder[T compare.Comparable[T]]() builder.Builder[T] {
	return &comparableBuilder[T]{
		from:     None[[]T](),
		capacity: None[int](),
	}
}

type comparableBuilder[T compare.Comparable[T]] struct {
	from     Option[[]T]
	capacity Option[int]
}

func (b *comparableBuilder[T]) From(items ...T) builder.Builder[T] {
	b.from = Some(items)
	return b
}

func (b *comparableBuilder[T]) Capacity(cap int) builder.Builder[T] {
	b.capacity = Some(cap)
	return b
}

func (b *comparableBuilder[T]) Build() sequence.Sequence[T] {
	return internal.SliceFromComparableInterface[T](b.from.OrElse(func() core.Option[[]T] {
		return b.from.Or(Some(make([]T, 0, b.capacity.UnwrapOrDefault())))
	}).Unwrap()...)
}
