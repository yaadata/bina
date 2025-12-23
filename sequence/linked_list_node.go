package sequence

type LinkedListNode[T any] interface {
	Value() T
	SetValue(value T)
}
