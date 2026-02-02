package collection

import (
	. "codeberg.org/yaadata/opt"
)

// BTree is a self-balancing [SearchTree] where nodes can have multiple keys and children.
type BTree[K any, V any] interface {
	SearchTree[K, V]

	// GetNode returns the node containing the given key, or None if not found.
	GetNode(key K) Option[BTreeNode[K, V]]

	// Order returns the branching factor of the tree.
	Order() int
}
