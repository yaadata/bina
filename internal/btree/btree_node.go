package btree

import (
	"cmp"
	"iter"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
	"codeberg.org/yaadata/opt/core"
	"codeberg.org/yaadata/opt/extension"
)

type Node[K cmp.Ordered, V any] struct {
	inner          kv.Pair[K, V]
	parent         Option[Node[K, V]]
	lessSibling    Option[Node[K, V]]
	greaterSibling Option[Node[K, V]]
	lessBranch     collection.Slice[Node[K, V]]
	greaterBranch  collection.Slice[Node[K, V]]
}

// compile time interface guard check
var _ collection.BTreeNode[int] = (*Node[int, int])(nil)

func (n *Node[K, V]) LessBranch() iter.Seq[V] {
	return func(yield func(V) bool) {
		for child := range n.lessBranch.Values() {
			if !yield(child.inner.Value()) {
				return
			}
		}
	}
}

func (n *Node[K, V]) GreaterBranch() iter.Seq[V] {
	return func(yield func(V) bool) {
		for child := range n.greaterBranch.Values() {
			if !yield(child.inner.Value()) {
				return
			}
		}
	}
}

func (n *Node[K, V]) Parent() Option[V] {
	return extension.OptionAndThen(n.parent, func(node Node[K, V]) core.Option[V] {
		return Some(node.inner.Value())
	})
}

func (n *Node[K, V]) Value() V {
	return n.inner.Value()
}
