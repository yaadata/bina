package compare

type Comparable[T any] interface {
	Compare(other T) Order
	Equal(other T) bool
}
