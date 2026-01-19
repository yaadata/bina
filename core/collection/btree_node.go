package collection

import (
	"iter"

	. "codeberg.org/yaadata/opt"
)

type BTreeNode[T any] interface {
	Children() iter.Seq[T]
	Parent() Option[T]
	Value() T
}
