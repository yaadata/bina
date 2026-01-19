package btree

import (
	"iter"

	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/collection"
)

type node[T any] struct {
	value    T
	parent   Option[T]
	children collection.Slice[T]
}

func (n *node[T]) Children() iter.Seq[T] {
	return n.children.Values()
}

func (n *node[T]) Parent() Option[T] {
	return n.parent
}

func (n *node[T]) Value() T {
	return n.value
}
