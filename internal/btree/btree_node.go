package btree

import (
	"cmp"
	"iter"

	. "codeberg.org/yaadata/opt"
	"codeberg.org/yaadata/opt/core"
	"codeberg.org/yaadata/opt/extension"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/kv"
)

type Node[K cmp.Ordered, V any] struct {
	elements []kv.Pair[K, V]
	parent   Option[Node[K, V]]
	children collection.Slice[Node[K, V]]
}

// compile time interface guard check
var _ collection.BTreeNode[int, int] = (*Node[int, int])(nil)

func (n *Node[K, V]) Children() iter.Seq[collection.BTreeNode[K, V]] {
	return func(yield func(collection.BTreeNode[K, V]) bool) {
		for child := range n.children.Values() {
			if !yield(&child) {
				return
			}
		}
	}
}

func (n *Node[K, V]) Parent() Option[collection.BTreeNode[K, V]] {
	return extension.OptionAndThen(n.parent, func(node Node[K, V]) core.Option[collection.BTreeNode[K, V]] {
		var n collection.BTreeNode[K, V] = &node
		return Some(n)
	})
}

func (n *Node[K, V]) Value() []kv.Pair[K, V] {
	return n.elements
}
