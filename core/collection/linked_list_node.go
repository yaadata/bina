package collection

type LinkedListNode[T any] interface {
	SetValue(value T)
	Value() T
}
