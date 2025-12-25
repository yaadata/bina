package circularlinkedlist

import "codeberg.org/yaadata/bina/core/compare"

func mergeSort[T any](leftHead *linkedListNode[T], fn func(a, b T) compare.Order) *linkedListNode[T] {
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

func findMiddle[T any](head *linkedListNode[T]) *linkedListNode[T] {
	slow, fast := head, head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func merge[T any](left, right *linkedListNode[T], fn func(a, b T) compare.Order) *linkedListNode[T] {
	result := &linkedListNode[T]{}
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

func tail[T any](node *linkedListNode[T]) *linkedListNode[T] {
	if node == nil || node.next == nil {
		return node
	}
	return tail(node.next)
}
