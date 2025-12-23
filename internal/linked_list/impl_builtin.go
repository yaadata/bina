package linkedlist

import (
	"iter"

	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
)

type linkedlistFromBuiltin[T comparable] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]
	len  int
}

var _ sequence.LinkedList[int] = (*linkedlistFromBuiltin[int])(nil)

func LinkedListFromBuiltin[T comparable]() sequence.LinkedList[T] {
	return &linkedlistFromBuiltin[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (s *linkedlistFromBuiltin[T]) Len() int {
	return s.len
}

func (s *linkedlistFromBuiltin[T]) IsEmpty() bool {
	return s.len == 0
}

func (s *linkedlistFromBuiltin[T]) Clear() {
	s.head = nil
	s.tail = nil
	s.len = 0
}

func (s *linkedlistFromBuiltin[T]) Contains(element T) bool {
	for value := range s.All() {
		if value == element {
			return true
		}
	}
	return false
}

func (s *linkedlistFromBuiltin[T]) Any(predicate predicate.Predicate[T]) bool {
	for value := range s.All() {
		if predicate(value) {
			return true
		}
	}
	return false
}

func (s *linkedlistFromBuiltin[T]) Count(predicate predicate.Predicate[T]) int {
	var result int
	for value := range s.All() {
		if predicate(value) {
			result++
		}
	}
	return result
}

func (s *linkedlistFromBuiltin[T]) Every(predicate predicate.Predicate[T]) bool {
	for value := range s.All() {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func (s *linkedlistFromBuiltin[T]) ForEach(fn func(value T)) {
	for value := range s.All() {
		fn(value)
	}
}

func (s *linkedlistFromBuiltin[T]) RemoveAt(position int) Option[T] {
	var currentIndex int
	var previous *linkedListNode[T]
	for node := s.head; node != nil; node = node.next {
		if currentIndex == position {
			if previous == nil {
				s.head = node.next
			} else {
				previous.next = node.next
			}
			s.len--
			return Some(node.value)
		}
		previous = node
		currentIndex++
	}
	return None[T]()
}

func (s *linkedlistFromBuiltin[T]) Append(item T) {
	node := &linkedListNode[T]{
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
	s.len++
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

func (s *linkedlistFromBuiltin[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	for value := range s.All() {
		if predicate(value) {
			return Some(value)
		}
	}
	return None[T]()
}

func (s *linkedlistFromBuiltin[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	for index, value := range s.Enumerate() {
		if predicate(value) {
			return Some(index)
		}
	}
	return None[int]()
}

func (s *linkedlistFromBuiltin[T]) Get(targetIndex int) Option[T] {
	for index, value := range s.Enumerate() {
		if index == targetIndex {
			return Some(value)
		}
	}
	return None[T]()
}

func (s *linkedlistFromBuiltin[T]) Insert(index int, item T) {
	if index < 0 {
		panic("index cannot be less than zero")
	}
	newNode := &linkedListNode[T]{
		value: item,
		next:  nil,
	}
	if index == 0 {
		s.head, s.tail = newNode, newNode
		s.len++
		return
	}
	currentIndex := 1
	previousNode := s.head
	for node := previousNode.next; node != nil; node = node.next {
		if index == currentIndex {
			previousNode.next = newNode
			newNode.next = node
			s.len++
			return
		}
		previousNode = node
		currentIndex++
	}
}

func (s *linkedlistFromBuiltin[T]) Retain(predicate predicate.Predicate[T]) {
	if s.len == 0 {
		return
	}
	previousNode := s.head
	for node := previousNode.next; node != nil; node = node.next {
		if !predicate(node.value) {
			previousNode.next = node.next
			s.len--
		} else {
			previousNode = node
		}
	}
	if !predicate(s.head.value) {
		if s.head == s.tail {
			s.head, s.tail = nil, nil
			s.len = 0
		} else {
			s.head = s.head.next
			s.len--
		}
	}
}

func (s *linkedlistFromBuiltin[T]) Sort(fn func(a, b T) compare.Order) {
	s.head = mergeSort(s.head, fn)
	s.tail = tail(s.head)
}

func (s *linkedlistFromBuiltin[T]) ToSlice() []T {
	res := make([]T, 0, s.len)
	for value := range s.All() {
		res = append(res, value)
	}
	return res
}

func (s *linkedlistFromBuiltin[T]) Extend(values ...T) {
	for _, value := range values {
		nextNode := &linkedListNode[T]{
			value: value,
			next:  nil,
		}
		if s.tail != nil {
			s.tail.next = nextNode
			s.tail = nextNode
			s.len++
		} else {
			s.head = nextNode
			s.tail = nextNode
			s.len = 1
		}
	}
}

func (s *linkedlistFromBuiltin[T]) ExtendFromSequence(seq sequence.Sequence[T]) {
	for value := range seq.All() {
		nextNode := &linkedListNode[T]{
			value: value,
			next:  nil,
		}
		if s.tail != nil {
			s.tail.next = nextNode
			s.tail = nextNode
			s.len++
		} else {
			s.head = nextNode
			s.tail = nextNode
			s.len = 1
		}
	}
}

func (s *linkedlistFromBuiltin[T]) GetNodeAt(index int) Option[sequence.LinkedListNode[T]] {
	currentIndex := 0
	for node := s.head; node != nil; node = node.next {
		if currentIndex == index {
			var res sequence.LinkedListNode[T] = node
			return Some(res)
		}
		currentIndex++
	}
	return None[sequence.LinkedListNode[T]]()
}

func (s *linkedlistFromBuiltin[T]) Head() Option[sequence.LinkedListNode[T]] {
	return optionalNode(s.head)
}

func (s *linkedlistFromBuiltin[T]) Prepend(value T) {
	newHead := &linkedListNode[T]{
		value: value,
		next:  nil,
	}
	if s.head != nil {
		newHead.next = s.head
		s.head = newHead
	} else {
		s.head = newHead
		s.tail = newHead
	}
	s.len++
}

func (s *linkedlistFromBuiltin[T]) Tail() Option[sequence.LinkedListNode[T]] {
	return optionalNode(s.tail)
}
