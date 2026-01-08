package hashset

import (
	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/hashable"
	hashset "codeberg.org/yaadata/bina/internal/hashset"
	"codeberg.org/yaadata/bina/set"
)

func NewBuiltinBuilder[T comparable]() Builder[T, set.Set[T], *builtinBuilder[T]] {
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

func (b *builtinBuilder[T]) Build() set.Set[T] {
	s := hashset.HashSetFromBuiltin[T](b.capacity.UnwrapOrDefault())
	s.Extend(b.from.UnwrapOrDefault()...)
	return s
}

func NewHashableBuilder[K comparable, T hashable.Hashable[K]]() Builder[T, set.Set[T], *hashableBuilder[K, T]] {
	return &hashableBuilder[K, T]{
		from:     None[[]T](),
		capacity: None[int](),
	}
}

type hashableBuilder[K comparable, T hashable.Hashable[K]] struct {
	from     Option[[]T]
	capacity Option[int]
}

func (b *hashableBuilder[K, T]) From(items ...T) *hashableBuilder[K, T] {
	b.from = Some(items)
	return b
}

func (b *hashableBuilder[K, T]) Capacity(cap int) *hashableBuilder[K, T] {
	b.capacity = Some(cap)
	return b
}

func (b *hashableBuilder[K, T]) Build() set.Set[T] {
	s := hashset.HashSetFromHashable[K, T](b.capacity.UnwrapOrDefault())
	s.Extend(b.from.UnwrapOrDefault()...)
	return s
}
