package collection

import (
	. "codeberg.org/yaadata/opt"
)

// Stack is a last-in-first-out (LIFO) sequence.
type Stack[T any] interface {
	Sequence[T]
	// Push adds an element to the top of the stack.
	Push(element T)
	// Pop removes and returns the element at the top, or None if empty.
	Pop() Option[T]
	// Peek returns the element at the top without removing it, or None if empty.
	Peek() Option[T]
}
