package collection

// LinkedListNode is the base interface for nodes in a linked list.
type LinkedListNode[T any] interface {
	// SetValue updates the value stored in the node.
	SetValue(value T)

	// Value returns the value stored in the node.
	Value() T
}
