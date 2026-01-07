package circularlinkedlist

import (
	"iter"

	. "github.com/yaadata/optionsgo"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/sequence"
)

type linkedListFromComparable[T compare.Comparable[T]] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]
	len  int
}

// Compile-time interface implementation check for sliceComparableInterface
func _[T compare.Comparable[T]]() {
	var _ sequence.LinkedList[T, sequence.DoublyLinkedListNode[T]] = (*linkedListFromComparable[T])(nil)
}

func LinkedListFromComparable[T compare.Comparable[T]]() sequence.LinkedList[T, sequence.DoublyLinkedListNode[T]] {
	return &linkedListFromComparable[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (s *linkedListFromComparable[T]) Len() int {
	return s.len
}

func (s *linkedListFromComparable[T]) IsEmpty() bool {
	return s.len == 0
}

func (s *linkedListFromComparable[T]) Clear() {
	s.head = nil
	s.tail = nil
	s.len = 0
}

func (s *linkedListFromComparable[T]) Contains(element T) bool {
	for value := range s.Values() {
		if value.Equal(element) {
			return true
		}
	}
	return false
}

func (s *linkedListFromComparable[T]) Any(predicate predicate.Predicate[T]) bool {
	for value := range s.Values() {
		if predicate(value) {
			return true
		}
	}
	return false
}

func (s *linkedListFromComparable[T]) Count(predicate predicate.Predicate[T]) int {
	var result int
	for value := range s.Values() {
		if predicate(value) {
			result++
		}
	}
	return result
}

func (s *linkedListFromComparable[T]) Every(predicate predicate.Predicate[T]) bool {
	for value := range s.Values() {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func (s *linkedListFromComparable[T]) ForEach(fn func(value T)) {
	for value := range s.Values() {
		fn(value)
	}
}

func (s *linkedListFromComparable[T]) RemoveAt(position int) Option[T] {
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

func (s *linkedListFromComparable[T]) Append(item T) {
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

func (s *linkedListFromComparable[T]) Values() iter.Seq[T] {
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

func (s *linkedListFromComparable[T]) All() iter.Seq2[int, T] {
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

func (s *linkedListFromComparable[T]) Find(predicate predicate.Predicate[T]) Option[T] {
	for value := range s.Values() {
		if predicate(value) {
			return Some(value)
		}
	}
	return None[T]()
}

func (s *linkedListFromComparable[T]) FindIndex(predicate predicate.Predicate[T]) Option[int] {
	for index, value := range s.All() {
		if predicate(value) {
			return Some(index)
		}
	}
	return None[int]()
}

func (s *linkedListFromComparable[T]) Get(targetIndex int) Option[T] {
	for index, value := range s.All() {
		if index == targetIndex {
			return Some(value)
		}
	}
	return None[T]()
}

func (s *linkedListFromComparable[T]) Insert(index int, item T) bool {
	if index < 0 || index > s.len {
		return false
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
		return true
	}
	if index == s.len {
		connectNodes(s.tail, newNode)
		connectNodes(newNode, s.head)
		s.tail = newNode
		s.len++
		return true
	}
	previousNode := s.head
	for i := 0; i < index-1; i++ {
		previousNode = previousNode.next
	}
	connectNodes(newNode, previousNode.next)
	connectNodes(previousNode, newNode)
	s.len++
	return true
}

func (s *linkedListFromComparable[T]) Retain(predicate predicate.Predicate[T]) {
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

func (s *linkedListFromComparable[T]) Sort(fn func(a, b T) compare.Order) {
	if s.tail != nil {
		s.tail.next = nil // Break circular link before sorting
	}
	s.head = mergeSort(s.head, fn)
	s.tail = tail(s.head)
	if s.tail != nil {
		connectNodes(s.tail, s.head) // Restore circular link
	}
}

func (s *linkedListFromComparable[T]) ToSlice() []T {
	res := make([]T, 0, s.len)
	for value := range s.Values() {
		res = append(res, value)
	}
	return res
}

func (s *linkedListFromComparable[T]) Extend(values ...T) {
	for _, value := range values {
		nextNode := newLinkedListNode(value)
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

func (s *linkedListFromComparable[T]) ExtendFromSequence(seq sequence.Sequence[T]) {
	for value := range seq.Values() {
		nextNode := &linkedListNode[T]{
			value: value,
			next:  nil,
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

func (s *linkedListFromComparable[T]) GetNodeAt(index int) Option[sequence.DoublyLinkedListNode[T]] {
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

func (s *linkedListFromComparable[T]) Head() Option[sequence.DoublyLinkedListNode[T]] {
	return optionalNode(s.head)
}

func (s *linkedListFromComparable[T]) Prepend(value T) {
	newHead := newLinkedListNode(value)
	if s.head != nil {
		connectNodes(newHead, s.head)
		connectNodes(s.tail, newHead)
		s.head = newHead
	} else {
		s.head = newHead
		s.tail = newHead
	}
	s.len++
}

func (s *linkedListFromComparable[T]) Tail() Option[sequence.DoublyLinkedListNode[T]] {
	return optionalNode(s.tail)
}
