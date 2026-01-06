package compare

type comparableWrapper[T any] struct {
	inner T
	fn    func(a, b T) Order
}

var _ Comparable[comparableWrapper[int]] = (*comparableWrapper[int])(nil)

func (e *comparableWrapper[T]) Compare(other comparableWrapper[T]) Order {
	return e.fn(e.inner, other.inner)
}

func (e *comparableWrapper[T]) Equal(other comparableWrapper[T]) bool {
	return e.fn(e.inner, other.inner) == OrderEqual
}

func (e *comparableWrapper[T]) Inner() T {
	return e.inner
}

func ToComparable[T any](inner T, fn func(a, b T) Order) *comparableWrapper[T] {
	return &comparableWrapper[T]{
		inner: inner,
		fn:    fn,
	}
}
