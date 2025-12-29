package slice

import (
	. "github.com/yaadata/optionsgo"
	"github.com/yaadata/optionsgo/core"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/internal/slice"
	"codeberg.org/yaadata/bina/sequence"
)

func NewBuiltinBuilder[T comparable]() Builder[T, sequence.Slice[T], *builtinBuilder[T]] {
	return &builtinBuilder[T]{
		from:     None[[]T](),
		capacity: None[int](),
	}
}

type builtinBuilder[T comparable] struct {
	from     Option[[]T]
	capacity Option[int]
}

func (b *builtinBuilder[T]) From(items ...T) *builtinBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *builtinBuilder[T]) Capacity(cap int) *builtinBuilder[T] {
	b.capacity = Some(cap)
	return b
}

func (b *builtinBuilder[T]) Build() sequence.Slice[T] {
	return slice.SliceFromBuiltin(b.from.OrElse(func() core.Option[[]T] {
		return b.from.Or(Some(make([]T, 0, b.capacity.UnwrapOrDefault())))
	}).Unwrap()...)
}

func NewComparableBuilder[T compare.Comparable[T]]() Builder[T, sequence.Slice[T], *comparableBuilder[T]] {
	return &comparableBuilder[T]{
		from:     None[[]T](),
		capacity: None[int](),
	}
}

type comparableBuilder[T compare.Comparable[T]] struct {
	from     Option[[]T]
	capacity Option[int]
}

func (b *comparableBuilder[T]) From(items ...T) *comparableBuilder[T] {
	b.from = Some(items)
	return b
}

func (b *comparableBuilder[T]) Capacity(cap int) *comparableBuilder[T] {
	b.capacity = Some(cap)
	return b
}

func (b *comparableBuilder[T]) Build() sequence.Slice[T] {
	return slice.SliceFromComparableInterface(b.from.OrElse(func() core.Option[[]T] {
		return b.from.Or(Some(make([]T, 0, b.capacity.UnwrapOrDefault())))
	}).Unwrap()...)
}
