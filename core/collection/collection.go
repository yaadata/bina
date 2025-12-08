package collection

type Collection[T any] interface {
	Len() int
	Contains(element T) bool
	IsEmpty() bool
	Clear()
}
