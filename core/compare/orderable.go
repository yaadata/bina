package compare

type Orderable[T any] interface {
	Order(other T) Order
}
