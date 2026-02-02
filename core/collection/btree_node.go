package collection

import (
	"iter"

	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
)

// BTreeNode represents a node in a [BTree] containing keys and child pointers.
type BTreeNode[K any, V any] interface {
	// Children returns an iterator over this node's child nodes.
	Children() iter.Seq[BTreeNode[K, V]]

	// Parent returns this node's parent, or None if this is the root.
	Parent() Option[BTreeNode[K, V]]

	// Values returns the key-value pairs stored in this node.
	Values() []kv.Pair[K, V]
}
