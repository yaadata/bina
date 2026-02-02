package hashable

// Hashable defines types that can produce a hash key.
// Used by hash-based collections when the type itself is not comparable.
type Hashable[T comparable] interface {
	// Hash returns a comparable key representing this value.
	// Equal values must return equal hash keys.
	Hash() T
}
