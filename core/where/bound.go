package where

// Bound specifies whether a range endpoint includes or excludes the boundary value.
type Bound int

const (
	// BoundInclusive includes the boundary value in the range.
	BoundInclusive Bound = iota
	// BoundExclusive excludes the boundary value from the range.
	BoundExclusive
)
