package compare

import "github.com/yaadata/bina/core/compare"

type ReferenceComparableInt int

var _ compare.Comparable[ReferenceComparableInt] = (*ReferenceComparableInt)(nil)

func (c *ReferenceComparableInt) Compare(other ReferenceComparableInt) compare.Order {
	if *c < other {
		return compare.OrderLess
	}
	if *c > other {
		return compare.OrderGreater
	}
	return compare.OrderEqual
}

func (c *ReferenceComparableInt) Equal(other ReferenceComparableInt) bool {
	return c.Compare(other).IsEqual()
}
