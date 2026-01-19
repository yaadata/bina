package array

import (
	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/internal/array"
)

type builtinBuilder[T comparable] struct {
	size Option[int]
}

func NewBuiltinBuilder[T comparable]() Builder[T, collection.Array[T], *builtinBuilder[T]] {
	return &builtinBuilder[T]{
		size: None[int](),
	}
}

func (b *builtinBuilder[T]) Size(size int) *builtinBuilder[T] {
	b.size = Some(size)
	return b
}

func (b *builtinBuilder[T]) Build() collection.Array[T] {
	return array.ArrayFromBuiltin[T](b.size.UnwrapOrElse(func() int {
		return 1
	}))
}

type comparableInterfaceBuilder[T compare.Comparable[T]] struct {
	size Option[int]
}

func NewComparableInterfaceBuilder[T compare.Comparable[T]]() Builder[T, collection.Array[T], *comparableInterfaceBuilder[T]] {
	return &comparableInterfaceBuilder[T]{
		size: None[int](),
	}
}

func (b *comparableInterfaceBuilder[T]) Size(size int) *comparableInterfaceBuilder[T] {
	b.size = Some(size)
	return b
}

func (b *comparableInterfaceBuilder[T]) Build() collection.Array[T] {
	return array.ArrayFromComparableInterface[T](b.size.UnwrapOrElse(func() int {
		return 1
	}))
}
