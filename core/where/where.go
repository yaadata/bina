package where

import (
	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
)

// Where defines a range with optional start and end bounds.
type Where[K any] struct {
	from      Option[K]
	fromBound Bound
	to        Option[K]
	toBound   Bound
}

// Default returns an unbounded range (no start or end limits).
func Default[K any]() *Where[K] {
	return &Where[K]{
		from: None[K](),
		to:   None[K](),
	}
}

// From returns the start point and its bound type.
func (c *Where[K]) From() kv.Pair[Option[K], Bound] {
	return kv.New(c.from, c.fromBound)
}

// To returns the end point and its bound type.
func (c *Where[K]) To() kv.Pair[Option[K], Bound] {
	return kv.New(c.to, c.toBound)
}
