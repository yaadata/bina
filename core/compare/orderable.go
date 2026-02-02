package compare

// Orderable defines ordering comparison for sortable types.
type Orderable[T any] interface {
	// Order returns OrderLess, OrderEqual, or OrderGreater.
	Order(other T) Order
}
