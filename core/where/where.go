package where

import (
	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
)

type Where[K any] struct {
	from      Option[K]
	fromBound Bound
	to        Option[K]
	toBound   Bound
}

func Default[K any]() *Where[K] {
	return &Where[K]{
		from: None[K](),
		to:   None[K](),
	}
}

func (c *Where[K]) From() kv.Pair[Option[K], Bound] {
	return kv.New(c.from, c.fromBound)
}

func (c *Where[K]) To() kv.Pair[Option[K], Bound] {
	return kv.New(c.to, c.toBound)
}
