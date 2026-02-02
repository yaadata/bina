package kv

// Pair represents an immutable key-value association.
type Pair[K any, V any] interface {
	// Key returns the key component of the pair.
	Key() K

	// Value returns the value component of the pair.
	Value() V
}
