package compare

type ComparableType[T any] struct {
	inner T
	fn    func(a, b T) Order
}

var _ Comparable[ComparableType[int]] = (*ComparableType[int])(nil)

func (e *ComparableType[T]) Compare(other ComparableType[T]) Order {
	return e.fn(e.inner, other.inner)
}

func (e *ComparableType[T]) Equal(other ComparableType[T]) bool {
	return e.fn(e.inner, other.inner) == OrderEqual
}

func (e *ComparableType[T]) Inner() T {
	return e.inner
}

func ToComparableType[T any](inner T, fn func(a, b T) Order) *ComparableType[T] {
	return &ComparableType[T]{
		inner: inner,
		fn:    fn,
	}
}
