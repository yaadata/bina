package queue_test

import (
	"testing"

	"github.com/shoenig/test/must"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/sequence"
	"codeberg.org/yaadata/bina/sequence/queue"
)

type ComparableInt int

func (c ComparableInt) Equal(other ComparableInt) bool {
	return c == other
}

func TestQueueFromBuiltin(t *testing.T) {
	testCases := []struct {
		name     string
		backedBy queue.QueueBackedBy
	}{
		{"Slice", queue.QueueBackedBySlice},
		{"SinglyLinkedList", queue.QueueBackedBySinglyLinkedList},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newQueue := func(items ...int) sequence.Queue[int] {
				builder := queue.NewBuiltinBuilder[int]()
				builder.BackedBy(tc.backedBy)
				return builder.From(items...).Build()
			}

			t.Run("Can build", func(t *testing.T) {
				// ========= [A]rrange =========
				builder := queue.NewBuiltinBuilder[int]()
				builder.BackedBy(tc.backedBy)
				q := builder.Build()
				// ========= [A]ssert  =========
				must.Eq(t, 0, q.Len())
			})

			t.Run("Can build from items", func(t *testing.T) {
				// ========= [A]rrange =========
				q := newQueue(1, 2, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 3, q.Len())
				must.True(t, q.Contains(1))
				must.True(t, q.Contains(2))
				must.True(t, q.Contains(3))
			})

			t.Run("Can build from items with duplicates", func(t *testing.T) {
				// ========= [A]rrange =========
				q := newQueue(1, 2, 2, 3, 3, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 6, q.Len())
			})

			t.Run("Collection methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				q := newQueue(1, 2, 3)

				// SCENARIO: Len
				t.Run("Len", func(t *testing.T) {
					// ========= [A]ct     =========
					length := q.Len()
					// ========= [A]ssert  =========
					must.Eq(t, 3, length)
				})

				// SCENARIO: Contains
				t.Run("Contains - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Contains(4)
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Contains - true", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Contains(2)
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: IsEmpty
				t.Run("IsEmpty - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.IsEmpty()
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: Clear
				t.Run("Clear", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					q.Clear()
					// ========= [A]ssert  =========
					must.Eq(t, 0, q.Len())
					must.True(t, q.IsEmpty())
				})
			})

			t.Run("Aggregate methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				q := newQueue(1, 2, 3, 4, 5)

				// SCENARIO: Any
				t.Run("Any - False", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Any(func(item int) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Any - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Any(func(item int) bool {
						return item > 3
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: Count
				t.Run("Count", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Count(func(item int) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, actual)
				})

				// SCENARIO: Every
				t.Run("Every - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Every(func(item int) bool {
						return item < 10
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				t.Run("Every - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Every(func(item int) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: ForEach
				t.Run("ForEach", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					q.ForEach(func(item int) {
						count++
					})
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})
			})

			t.Run("Queue operations work", func(t *testing.T) {
				// SCENARIO: Enqueue
				t.Run("Enqueue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					q.Enqueue(4)
					// ========= [A]ssert  =========
					must.Eq(t, 4, q.Len())
					must.True(t, q.Peek().IsSome())
					must.Eq(t, 1, q.Peek().Unwrap()) // First element unchanged
				})

				// SCENARIO: Dequeue
				t.Run("Dequeue - non-empty queue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					dequeued := q.Dequeue()
					// ========= [A]ssert  =========
					must.True(t, dequeued.IsSome())
					must.Eq(t, 1, dequeued.Unwrap()) // First element
					must.Eq(t, 2, q.Len())
				})

				t.Run("Dequeue - empty queue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue()
					// ========= [A]ct     =========
					dequeued := q.Dequeue()
					// ========= [A]ssert  =========
					must.True(t, dequeued.IsNone())
				})

				// SCENARIO: Peek
				t.Run("Peek - non-empty queue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					peeked := q.Peek()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsSome())
					must.Eq(t, 1, peeked.Unwrap()) // First element
					must.Eq(t, 3, q.Len())         // Length unchanged
				})

				t.Run("Peek - empty queue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue()
					// ========= [A]ct     =========
					peeked := q.Peek()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsNone())
				})

				// SCENARIO: FIFO order
				t.Run("FIFO order", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue()
					// ========= [A]ct     =========
					q.Enqueue(1)
					q.Enqueue(2)
					q.Enqueue(3)
					// ========= [A]ssert  =========
					must.Eq(t, 1, q.Dequeue().Unwrap())
					must.Eq(t, 2, q.Dequeue().Unwrap())
					must.Eq(t, 3, q.Dequeue().Unwrap())
					must.True(t, q.Dequeue().IsNone())
				})
			})

			t.Run("Sequence methods work", func(t *testing.T) {
				// SCENARIO: Values
				t.Run("Values", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range q.Values() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: All
				t.Run("All", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range q.All() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: Find
				t.Run("Find - found", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					found := q.Find(func(item int) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsSome())
					must.Eq(t, 2, found.Unwrap())
				})

				t.Run("Find - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					found := q.Find(func(item int) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsNone())
				})

				// SCENARIO: FindIndex
				t.Run("FindIndex - found", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					idx := q.FindIndex(func(item int) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsSome())
					must.Eq(t, 1, idx.Unwrap())
				})

				t.Run("FindIndex - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					idx := q.FindIndex(func(item int) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsNone())
				})

				// SCENARIO: Get
				t.Run("Get - valid index", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					item := q.Get(1)
					// ========= [A]ssert  =========
					must.True(t, item.IsSome())
					must.Eq(t, 2, item.Unwrap())
				})

				t.Run("Get - invalid index", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					item := q.Get(10)
					// ========= [A]ssert  =========
					must.True(t, item.IsNone())
				})

				// SCENARIO: Retain
				t.Run("Retain", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3, 4, 5)
					// ========= [A]ct     =========
					q.Retain(func(item int) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, q.Len())
					must.True(t, q.Contains(2))
					must.True(t, q.Contains(4))
				})

				// SCENARIO: Sort
				t.Run("Sort", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(3, 1, 4, 1, 5, 9, 2, 6)
					// ========= [A]ct     =========
					q.Sort(func(a, b int) compare.Order {
						if a < b {
							return compare.OrderLess
						} else if a > b {
							return compare.OrderGreater
						}
						return compare.OrderEqual
					})
					// ========= [A]ssert  =========
					slice := q.ToSlice()
					must.Eq(t, []int{1, 1, 2, 3, 4, 5, 6, 9}, slice)
				})

				// SCENARIO: ToSlice
				t.Run("ToSlice", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					slice := q.ToSlice()
					// ========= [A]ssert  =========
					must.Eq(t, []int{1, 2, 3}, slice)
				})
			})
		})
	}
}

func TestQueueFromComparable(t *testing.T) {
	testCases := []struct {
		name     string
		backedBy queue.QueueBackedBy
	}{
		{"Slice", queue.QueueBackedBySlice},
		{"SinglyLinkedList", queue.QueueBackedBySinglyLinkedList},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newQueue := func(items ...ComparableInt) sequence.Queue[ComparableInt] {
				builder := queue.NewComparableBuilder[ComparableInt]()
				builder.BackedBy(tc.backedBy)
				return builder.From(items...).Build()
			}

			t.Run("Can build", func(t *testing.T) {
				// ========= [A]rrange =========
				builder := queue.NewComparableBuilder[ComparableInt]()
				builder.BackedBy(tc.backedBy)
				q := builder.Build()
				// ========= [A]ssert  =========
				must.Eq(t, 0, q.Len())
			})

			t.Run("Can build from items", func(t *testing.T) {
				// ========= [A]rrange =========
				q := newQueue(1, 2, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 3, q.Len())
				must.True(t, q.Contains(1))
				must.True(t, q.Contains(2))
				must.True(t, q.Contains(3))
			})

			t.Run("Can build from items with duplicates", func(t *testing.T) {
				// ========= [A]rrange =========
				q := newQueue(1, 2, 2, 3, 3, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 6, q.Len())
			})

			t.Run("Collection methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				q := newQueue(1, 2, 3)

				// SCENARIO: Len
				t.Run("Len", func(t *testing.T) {
					// ========= [A]ct     =========
					length := q.Len()
					// ========= [A]ssert  =========
					must.Eq(t, 3, length)
				})

				// SCENARIO: Contains
				t.Run("Contains - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Contains(4)
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Contains - true", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Contains(2)
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: IsEmpty
				t.Run("IsEmpty - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.IsEmpty()
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: Clear
				t.Run("Clear", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					q.Clear()
					// ========= [A]ssert  =========
					must.Eq(t, 0, q.Len())
					must.True(t, q.IsEmpty())
				})
			})

			t.Run("Aggregate methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				q := newQueue(1, 2, 3, 4, 5)

				// SCENARIO: Any
				t.Run("Any - False", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Any(func(item ComparableInt) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Any - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Any(func(item ComparableInt) bool {
						return item > 3
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: Count
				t.Run("Count", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Count(func(item ComparableInt) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, actual)
				})

				// SCENARIO: Every
				t.Run("Every - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Every(func(item ComparableInt) bool {
						return item < 10
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				t.Run("Every - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := q.Every(func(item ComparableInt) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: ForEach
				t.Run("ForEach", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					q.ForEach(func(item ComparableInt) {
						count++
					})
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})
			})

			t.Run("Queue operations work", func(t *testing.T) {
				// SCENARIO: Enqueue
				t.Run("Enqueue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					q.Enqueue(4)
					// ========= [A]ssert  =========
					must.Eq(t, 4, q.Len())
					must.True(t, q.Peek().IsSome())
					must.Eq(t, ComparableInt(1), q.Peek().Unwrap()) // First element unchanged
				})

				// SCENARIO: Dequeue
				t.Run("Dequeue - non-empty queue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					dequeued := q.Dequeue()
					// ========= [A]ssert  =========
					must.True(t, dequeued.IsSome())
					must.Eq(t, ComparableInt(1), dequeued.Unwrap()) // First element
					must.Eq(t, 2, q.Len())
				})

				t.Run("Dequeue - empty queue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue()
					// ========= [A]ct     =========
					dequeued := q.Dequeue()
					// ========= [A]ssert  =========
					must.True(t, dequeued.IsNone())
				})

				// SCENARIO: Peek
				t.Run("Peek - non-empty queue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					peeked := q.Peek()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsSome())
					must.Eq(t, ComparableInt(1), peeked.Unwrap()) // First element
					must.Eq(t, 3, q.Len())                        // Length unchanged
				})

				t.Run("Peek - empty queue", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue()
					// ========= [A]ct     =========
					peeked := q.Peek()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsNone())
				})

				// SCENARIO: FIFO order
				t.Run("FIFO order", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue()
					// ========= [A]ct     =========
					q.Enqueue(1)
					q.Enqueue(2)
					q.Enqueue(3)
					// ========= [A]ssert  =========
					must.Eq(t, ComparableInt(1), q.Dequeue().Unwrap())
					must.Eq(t, ComparableInt(2), q.Dequeue().Unwrap())
					must.Eq(t, ComparableInt(3), q.Dequeue().Unwrap())
					must.True(t, q.Dequeue().IsNone())
				})
			})

			t.Run("Sequence methods work", func(t *testing.T) {
				// SCENARIO: Values
				t.Run("Values", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range q.Values() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: All
				t.Run("All", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range q.All() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: Find
				t.Run("Find - found", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					found := q.Find(func(item ComparableInt) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsSome())
					must.Eq(t, ComparableInt(2), found.Unwrap())
				})

				t.Run("Find - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					found := q.Find(func(item ComparableInt) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsNone())
				})

				// SCENARIO: FindIndex
				t.Run("FindIndex - found", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					idx := q.FindIndex(func(item ComparableInt) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsSome())
					must.Eq(t, 1, idx.Unwrap())
				})

				t.Run("FindIndex - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					idx := q.FindIndex(func(item ComparableInt) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsNone())
				})

				// SCENARIO: Get
				t.Run("Get - valid index", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					item := q.Get(1)
					// ========= [A]ssert  =========
					must.True(t, item.IsSome())
					must.Eq(t, ComparableInt(2), item.Unwrap())
				})

				t.Run("Get - invalid index", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					item := q.Get(10)
					// ========= [A]ssert  =========
					must.True(t, item.IsNone())
				})

				// SCENARIO: Retain
				t.Run("Retain", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3, 4, 5)
					// ========= [A]ct     =========
					q.Retain(func(item ComparableInt) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, q.Len())
					must.True(t, q.Contains(2))
					must.True(t, q.Contains(4))
				})

				// SCENARIO: Sort
				t.Run("Sort", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(3, 1, 4, 1, 5, 9, 2, 6)
					// ========= [A]ct     =========
					q.Sort(func(a, b ComparableInt) compare.Order {
						if a < b {
							return compare.OrderLess
						} else if a > b {
							return compare.OrderGreater
						}
						return compare.OrderEqual
					})
					// ========= [A]ssert  =========
					slice := q.ToSlice()
					must.Eq(t, []ComparableInt{1, 1, 2, 3, 4, 5, 6, 9}, slice)
				})

				// SCENARIO: ToSlice
				t.Run("ToSlice", func(t *testing.T) {
					// ========= [A]rrange =========
					q := newQueue(1, 2, 3)
					// ========= [A]ct     =========
					slice := q.ToSlice()
					// ========= [A]ssert  =========
					must.Eq(t, []ComparableInt{1, 2, 3}, slice)
				})
			})
		})
	}
}
