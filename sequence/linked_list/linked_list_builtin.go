package linkedlist

import (
	"iter"

	"codeberg.org/yaadata/bina/core/shared"
	. "github.com/yaadata/optionsgo"
)

type linkedlistFromBuiltin[T comparable] struct {
	head *singlyLinkedListNode[T]
	tail *singlyLinkedListNode[T]
	len  int
}

func (s *linkedlistFromBuiltin[T]) Len() int {
	return s.len
}

func (s *linkedlistFromBuiltin[T]) IsEmpty() bool {
	return s.len == 0
}

func (s *linkedlistFromBuiltin[T]) Clear() {
	s.head = nil
}

func (s *linkedlistFromBuiltin[T]) Contains(element T) bool {
	for node := s.head; node != nil; node = node.next {
		if node.value == element {
			return true
		}
	}
	return false
}

func (s *linkedlistFromBuiltin[T]) Any(predicate shared.Predicate[T]) bool {
	for node := s.head; node != nil; node = node.next {
		if predicate(node.value) {
			return true
		}
	}
	return false
}

func (s *linkedlistFromBuiltin[T]) Count(predicate shared.Predicate[T]) int {
	var result int
	for node := s.head; node != nil; node = node.next {
		if predicate(node.value) {
			result++
		}
	}
	return result
}

func (s *linkedlistFromBuiltin[T]) Every(predicate shared.Predicate[T]) bool {
	for node := s.head; node != nil; node = node.next {
		if !predicate(node.value) {
			return false
		}
	}
	return true
}

func (s *linkedlistFromBuiltin[T]) ForEach(fn func(value T)) {
	for node := s.head; node != nil; node = node.next {
		fn(node.value)
	}
}

func (s *linkedlistFromBuiltin[T]) RemoveAt(position int) Option[T] {
	var currentIndex int
	var previous *singlyLinkedListNode[T]
	for node := s.head; node != nil; node = node.next {
		if currentIndex == position {
			if previous == nil {
				s.head = node.next
			} else {
				previous.next = node.next
			}
			return Some(node.value)
		}
		previous = node
		currentIndex++
	}
	return None[T]()
}

func (s *linkedlistFromBuiltin[T]) Append(item T) {
	node := &singlyLinkedListNode[T]{
		value: item,
		next:  nil,
	}
	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		s.tail.next = node
		s.tail = node
	}
}

func (s *linkedlistFromBuiltin[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for node := s.head; node != nil; node = node.next {
			if !yield(node.value) {
				return
			}
		}
	}
}

func (s *linkedlistFromBuiltin[T]) Enumerate() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		index := 0
		for node := s.head; node != nil; node = node.next {
			if !yield(index, node.value) {
				return
			}
			index++
		}
	}
}

func (s *linkedlistFromBuiltin[T]) Find(predicate shared.Predicate[T]) Option[T] {
	for node := s.head; node != nil; node = node.next {
		if predicate(node.value) {
			return Some(node.value)
		}
	}
	return None[T]()
}

func (s *linkedlistFromBuiltin[T]) FindIndex(predicate shared.Predicate[T]) Option[int] {
	index := 0
	for node := s.head; node != nil; node = node.next {
		if predicate(node.value) {
			return Some(index)
		}
		index++
	}
	return None[int]()
}
