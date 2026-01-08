package doublylinkedlist

import (
	"iter"

	. "codeberg.org/yaadata/opt"

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

func (s *linkedListFromComparable[T]) Append(item T) {
	node := &linkedListNode[T]{
		value: item,
		next:  nil,
	}
	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		s.tail.setNext(node)
		s.tail = node
	}
	s.len++
}

func (s *linkedListFromComparable[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for node := s.head; node != nil; node = node.next {
			if !yield(node.value) {
				return
			}
		}
	}
}

func (s *linkedListFromComparable[T]) All() iter.Seq2[int, T] {
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
	if index < 0 {
		return false
	}
	newNode := newLinkedListNode(item)
	if index == 0 {
		s.head, s.tail = newNode, newNode
		s.len++
		return true
	}
	currentIndex := 1
	previousNode := s.head
	for node := previousNode.next; node != nil; node = node.next {
		if index == currentIndex {
			previousNode.setNext(newNode)
			newNode.setNext(node)
			s.len++
			return true
		}
		previousNode = node
		currentIndex++
	}
	return false
}

func (s *linkedListFromComparable[T]) Retain(predicate predicate.Predicate[T]) {
	if s.len == 0 {
		return
	}
	previousNode := s.head
	for node := previousNode.next; node != nil; node = node.next {
		if !predicate(node.value) {
			previousNode.setNext(node.next)
			if node.next == nil {
				s.tail = previousNode
			}
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
			s.head.previous = nil
			s.len--
		}
	}
}

func (s *linkedListFromComparable[T]) Sort(fn func(a, b T) compare.Order) {
	s.head = mergeSort(s.head, fn)
	s.tail = tail(s.head)
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
		nextNode := &linkedListNode[T]{
			value: value,
			next:  nil,
		}
		if s.tail != nil {
			s.tail.setNext(nextNode)
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
			s.tail.setNext(nextNode)
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
	currentIndex := 0
	for node := s.head; node != nil; node = node.next {
		if currentIndex == index {
			var res sequence.DoublyLinkedListNode[T] = node
			return Some(res)
		}
		currentIndex++
	}
	return None[sequence.DoublyLinkedListNode[T]]()
}

func (s *linkedListFromComparable[T]) Head() Option[sequence.DoublyLinkedListNode[T]] {
	return optionalNode(s.head)
}

func (s *linkedListFromComparable[T]) Prepend(value T) {
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

func (s *linkedListFromComparable[T]) Tail() Option[sequence.DoublyLinkedListNode[T]] {
	return optionalNode(s.tail)
}
