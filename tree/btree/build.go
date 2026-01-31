package btree

import (
	"cmp"

	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/kv"
	"codeberg.org/yaadata/bina/internal/btree"
)

func NewBuiltinBuilder[K cmp.Ordered, V any]() Builder[K, V, collection.BTree[K, V], *builtinBuilder[K, V]] {
	return &builtinBuilder[K, V]{
		order: None[int](),
		from:  None[[]kv.Pair[K, V]](),
	}

}

type builtinBuilder[K cmp.Ordered, V any] struct {
	order Option[int]
	from  Option[[]kv.Pair[K, V]]
}

func (d *builtinBuilder[K, V]) From(items ...kv.Pair[K, V]) *builtinBuilder[K, V] {
	d.from = Some(items)
	return d
}

func (d *builtinBuilder[K, V]) Order(order int) *builtinBuilder[K, V] {
	d.order = Some(order)
	return d
}

func (d *builtinBuilder[K, V]) Build() collection.BTree[K, V] {
	resp := btree.NewBuiltinImpl[K, V](d.order.UnwrapOrElse(func() int {
		return 5
	}))
	if d.from.IsSome() {
		for _, kv := range d.from.Unwrap() {
			resp.Put(kv.Key(), kv.Value())
		}
	}
	return resp
}
