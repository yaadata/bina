package collection

import (
	. "codeberg.org/yaadata/opt"
)

// Queue is a first-in-first-out (FIFO) sequence.
type Queue[T any] interface {
	Sequence[T]

	// Enqueue adds an element to the back of the queue.
	Enqueue(item T)

	// Dequeue removes and returns the element at the front, or None if empty.
	Dequeue() Option[T]

	// Peek returns the element at the front without removing it, or None if empty.
	Peek() Option[T]
}
