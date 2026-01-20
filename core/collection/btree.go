package collection

import (
	. "codeberg.org/yaadata/opt"
)

type BTree[K any, V any] interface {
	SearchTree[K, V]
	GetNode(key K) Option[BTreeNode[V]]
	Order() int
}
