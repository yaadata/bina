package compare

type Orderable[T any] interface {
	Compare(other T) Order
}
