package compare

// Comparable defines equality comparison for types that cannot use
// the built-in == operator. Implementations must be reflexive,
// symmetric, and transitive.
type Comparable[T any] interface {
	// Equal reports whether this value equals other.
	Equal(other T) bool
}
