package compare

import (
	. "codeberg.org/yaadata/opt"
)

// Order represents the result of comparing two values.
//
//	OrderLess    (-1): first value < second value
//	OrderEqual   (0):  first value == second value
//	OrderGreater (1):  first value > second value
type Order int

const (
	OrderLess    Order = -1
	OrderEqual   Order = 0
	OrderGreater Order = 1
)

func (s Order) Int() int {
	return int(s)
}

func (s Order) IsLess() bool {
	return s == OrderLess
}

func (s Order) IsLessThanOrEqualTo() bool {
	return s == OrderEqual || s == OrderLess
}

func (s Order) IsEqual() bool {
	return s == OrderEqual
}

func (s Order) IsGreater() bool {
	return s == OrderGreater
}

func (s Order) IsGreaterThanOrEqualTo() bool {
	return s == OrderEqual || s == OrderGreater
}

func (s Order) String() Option[string] {
	switch s {
	case OrderLess:
		return Some("Less")
	case OrderEqual:
		return Some("Equal")
	case OrderGreater:
		return Some("Greater")
	}
	return None[string]()
}
