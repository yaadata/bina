package compare

type Comparable[T any] interface {
	Equal(other T) bool
}
