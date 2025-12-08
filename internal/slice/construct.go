package slice

import (
	"github.com/yaadata/bina/core/compare"
	"github.com/yaadata/bina/core/sequential"
)

func NewSliceSequence[T compare.Comparable[T]](items ...T) sequential.Sequence[T] {
	return &sliceCompareInterface[T]{_data: items}
}

func NewSliceSequenceFromComparable[T comparable](items ...T) sequential.Sequence[T] {
	return &slice[T]{_data: items}
}
