package collection

import (
	. "codeberg.org/yaadata/opt"
)

// Deque is a double-ended queue supporting insertion and removal at both ends.
type Deque[T any] interface {
	Sequence[T]
	// PushFront adds an element to the front.
	PushFront(item T)
	// PushBack adds an element to the back.
	PushBack(item T)
	// PopFront removes and returns the element at the front, or None if empty.
	PopFront() Option[T]
	// PopBack removes and returns the element at the back, or None if empty.
	PopBack() Option[T]
	// PeekFront returns the element at the front without removing it, or None if empty.
	PeekFront() Option[T]
	// PeekBack returns the element at the back without removing it, or None if empty.
	PeekBack() Option[T]
}
