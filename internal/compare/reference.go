package compare

import "codeberg.org/yaadata/bina/core/compare"

type ReferenceComparableInt int

var _ compare.Comparable[ReferenceComparableInt] = (*ReferenceComparableInt)(nil)
var _ compare.Orderable[ReferenceComparableInt] = (*ReferenceComparableInt)(nil)

func (c *ReferenceComparableInt) Order(other ReferenceComparableInt) compare.Order {
	if *c < other {
		return compare.OrderLess
	}
	if *c > other {
		return compare.OrderGreater
	}
	return compare.OrderEqual
}

func (c *ReferenceComparableInt) Equal(other ReferenceComparableInt) bool {
	return c.Order(other).IsEqual()
}
