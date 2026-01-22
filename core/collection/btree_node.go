package collection

import (
	"iter"

	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
)

type BTreeNode[K any, V any] interface {
	Children() iter.Seq[BTreeNode[K, V]]
	Parent() Option[BTreeNode[K, V]]
	Values() []kv.Pair[K, V]
}
