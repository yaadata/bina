package circularlinkedlist

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

var _ sequence.LinkedList[int, sequence.DoublyLinkedListNode[int]] = (*linkedlistFromBuiltin[int])(nil)

func LinkedListFromBuiltin[T comparable]() sequence.LinkedList[T, sequence.DoublyLinkedListNode[T]] {
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
	if position < 0 || position >= s.len {
		return None[T]()
	}
	var previous *linkedListNode[T]
	node := s.head
	for range position {
		previous = node
		node = node.next
	}
	value := node.value
	if previous == nil {
		s.head = node.next
		if s.tail != nil {
			connectNodes(s.tail, s.head)
		}
	} else {
		previous.next = node.next
		if node == s.tail {
			s.tail = previous
		}
	}
	s.len--
	if s.len == 0 {
		s.head = nil
		s.tail = nil
	}
	return Some(value)
}

func (s *linkedlistFromBuiltin[T]) Append(item T) {
	node := newLinkedListNode(item)
	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		connectNodes(s.tail, node)
		connectNodes(node, s.head)
		s.tail = node
	}
	s.len++
}

func (s *linkedlistFromBuiltin[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		node := s.head
		for i := 0; i < s.len; i++ {
			if !yield(node.value) {
				return
			}
			node = node.next
		}
	}
}

func (s *linkedlistFromBuiltin[T]) Enumerate() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		node := s.head
		for i := 0; i < s.len; i++ {
			if !yield(i, node.value) {
				return
			}
			node = node.next
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
	if index < 0 || index > s.len {
		panic("index out of bounds")
	}
	newNode := newLinkedListNode(item)
	if index == 0 {
		if s.head == nil {
			s.head = newNode
			s.tail = newNode
			connectNodes(s.head, s.tail)
		} else {
			connectNodes(newNode, s.head)
			connectNodes(s.tail, newNode)
			s.head = newNode
		}
		s.len++
		return
	}
	if index == s.len {
		connectNodes(s.tail, newNode)
		connectNodes(newNode, s.head)
		s.tail = newNode
		return
	}
	previousNode := s.head
	for i := 0; i < index-1; i++ {
		previousNode = previousNode.next
	}
	connectNodes(newNode, previousNode.next)
	connectNodes(previousNode, newNode)
	s.len++
}

func (s *linkedlistFromBuiltin[T]) Retain(predicate predicate.Predicate[T]) {
	if s.len == 0 {
		return
	}
	// Handle head removal first - find new head
	for s.head != nil && !predicate(s.head.value) {
		if s.head == s.tail {
			s.head, s.tail = nil, nil
			s.len = 0
			return
		}
		s.head = s.head.next
		s.len--
	}
	if s.head == nil {
		return
	}
	// Now iterate through the rest
	previousNode := s.head
	count := s.len - 1
	node := s.head.next
	for range count {
		next := node.next
		if !predicate(node.value) {
			previousNode.next = next
			if node == s.tail {
				s.tail = previousNode
			}
			s.len--
		} else {
			previousNode = node
		}
		node = next
	}
	// Restore circular link
	if s.tail != nil {
		connectNodes(s.tail, s.head)
	}
}

func (s *linkedlistFromBuiltin[T]) Sort(fn func(a, b T) compare.Order) {
	if s.tail != nil {
		s.tail.next = nil // Break circular link before sorting
	}
	s.head = mergeSort(s.head, fn)
	s.tail = tail(s.head)
	if s.tail != nil {
		connectNodes(s.tail, s.head) // Restore circular link
	}
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
			next:     nil,
			previous: nil,
			value:    value,
		}
		if s.tail != nil {
			connectNodes(s.tail, nextNode)
			connectNodes(nextNode, s.head)
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
			next:     nil,
			previous: nil,
			value:    value,
		}
		if s.tail != nil {
			connectNodes(s.tail, nextNode)
			connectNodes(nextNode, s.head)
			s.tail = nextNode
			s.len++
		} else {
			s.head = nextNode
			s.tail = nextNode
			s.len = 1
		}
	}
}

func (s *linkedlistFromBuiltin[T]) GetNodeAt(index int) Option[sequence.DoublyLinkedListNode[T]] {
	if index < 0 || index >= s.len {
		return None[sequence.DoublyLinkedListNode[T]]()
	}
	node := s.head
	for range index {
		node = node.next
	}
	var res sequence.DoublyLinkedListNode[T] = node
	return Some(res)
}

func (s *linkedlistFromBuiltin[T]) Head() Option[sequence.DoublyLinkedListNode[T]] {
	return optionalNode(s.head)
}

func (s *linkedlistFromBuiltin[T]) Prepend(value T) {
	newHead := newLinkedListNode(value)
	if s.head != nil {
		connectNodes(newHead, s.head)
		connectNodes(s.tail, newHead) // Maintain circular link
		s.head = newHead
	} else {
		s.head = newHead
		s.tail = newHead
	}
	s.len++
}

func (s *linkedlistFromBuiltin[T]) Tail() Option[sequence.DoublyLinkedListNode[T]] {
	return optionalNode(s.tail)
}
