package collection

// Collection defines the core operations shared by all collection types.
type Collection[T any] interface {
	// Len returns the number of elements in the collection.
	Len() int

	// Contains reports whether element is present in the collection.
	Contains(element T) bool

	// IsEmpty reports whether the collection has no elements.
	// For fixed-size collections like arrays, reports whether all
	// elements are zero values.
	IsEmpty() bool

	// Clear removes all elements from the collection.
	// For fixed-size collections like arrays, sets all elements
	// to their zero value.
	Clear()
}
