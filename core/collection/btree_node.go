package collection

import (
	"iter"

	. "codeberg.org/yaadata/opt"
)

type BTreeNode[T any] interface {
	GreaterBranch() iter.Seq[T]
	LessBranch() iter.Seq[T]
	Parent() Option[T]
	Value() T
}
