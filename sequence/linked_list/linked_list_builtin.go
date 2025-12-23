package linkedlist

import (
	"iter"

	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/sequence"
	"codeberg.org/yaadata/bina/core/shared"
)

type linkedlistFromBuiltin[T comparable] struct {
	head *singlyLinkedListNode[T]
	tail *singlyLinkedListNode[T]
	len  int
}

var _ sequence.LinkedList[int] = (*linkedlistFromBuiltin[int])(nil)

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
	for value := range s.All() {
		if value == element {
			return true
		}
	}
	return false
}

func (s *linkedlistFromBuiltin[T]) Any(predicate shared.Predicate[T]) bool {
	for value := range s.All() {
		if predicate(value) {
			return true
		}
	}
	return false
}

func (s *linkedlistFromBuiltin[T]) Count(predicate shared.Predicate[T]) int {
	var result int
	for value := range s.All() {
		if predicate(value) {
			result++
		}
	}
	return result
}

func (s *linkedlistFromBuiltin[T]) Every(predicate shared.Predicate[T]) bool {
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
	var previous *singlyLinkedListNode[T]
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

func (s *linkedlistFromBuiltin[T]) Find(predicate shared.Predicate[T]) Option[T] {
	for value := range s.All() {
		if predicate(value) {
			return Some(value)
		}
	}
	return None[T]()
}

func (s *linkedlistFromBuiltin[T]) FindIndex(predicate shared.Predicate[T]) Option[int] {
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
	newNode := &singlyLinkedListNode[T]{
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

func (s *linkedlistFromBuiltin[T]) Retain(predicate shared.Predicate[T]) {
	if s.len == 0 {
		return
	}
	previousNode := s.head
	for node := previousNode.next; node != nil; node = node.next {
		if !predicate(node.value) {
			previousNode.next = node.next
			s.len--
		}
	}
	if !predicate(s.head.value) {
		if s.head == s.tail {
			s.head, s.tail = nil, nil
			s.len = 0
		} else {
			s.head = s.head.next
		}
	}
}

func (s *linkedlistFromBuiltin[T]) Sort(fn func(a, b T) compare.Order) {
	s.head = mergeSort(s.head, fn)
	s.tail = tail(s.head)
}

func (s *linkedlistFromBuiltin[T]) ToSlice() []T {
	res := make([]T, s.len)
	for value := range s.All() {
		res = append(res, value)
	}
	return res
}

func (s *linkedlistFromBuiltin[T]) GetNodeAt(index int) Option[sequence.LinkedListNode[T]] {
	currentIndex := 0
	for node := s.head; node != nil; node = node.next {
		if currentIndex == index {
			var res sequence.LinkedListNode[T] = node
			return Some(res)
		}
	}
	return None[sequence.LinkedListNode[T]]()
}

func (s *linkedlistFromBuiltin[T]) Prepend(value T) {
	newHead := &singlyLinkedListNode[T]{
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
}

func mergeSort[T any](leftHead *singlyLinkedListNode[T], fn func(a, b T) compare.Order) *singlyLinkedListNode[T] {
	if leftHead == nil || leftHead.next == nil {
		return leftHead
	}
	mid := findMiddle(leftHead)
	rightHead := mid.next
	mid.next = nil
	left := mergeSort(leftHead, fn)
	right := mergeSort(rightHead, fn)
	return merge(left, right, fn)
}

func findMiddle[T any](head *singlyLinkedListNode[T]) *singlyLinkedListNode[T] {
	slow, fast := head, head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func merge[T any](left, right *singlyLinkedListNode[T], fn func(a, b T) compare.Order) *singlyLinkedListNode[T] {
	result := &singlyLinkedListNode[T]{}
	current := result

	for left != nil && right != nil {
		if fn(left.value, right.value).IsLessThanOrEqualTo() {
			current.next = left
			left = left.next
		} else {
			current.next = right
			right = right.next
		}
		current = current.next
	}

	if left != nil {
		current.next = left
	} else {
		current.next = right
	}
	return result.next
}

func tail[T any](node *singlyLinkedListNode[T]) *singlyLinkedListNode[T] {
	if node == nil || node.next == nil {
		return node
	}
	return tail(node.next)
}
