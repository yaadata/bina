package deque_test

import (
	"testing"

	"github.com/shoenig/test/must"

	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/sequence"
	"codeberg.org/yaadata/bina/sequence/deque"
)

type ComparableInt int

func (c ComparableInt) Equal(other ComparableInt) bool {
	return c == other
}

func TestDequeFromBuiltin(t *testing.T) {
	testCases := []struct {
		name     string
		backedBy deque.DequeBackedBy
	}{
		{"Slice", deque.DequeBackedBySlice},
		{"DoublyLinkedList", deque.DequeBackedByDoublyLinkedList},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newDeque := func(items ...int) sequence.Deque[int] {
				builder := deque.NewBuiltinBuilder[int]()
				builder.BackedBy(tc.backedBy)
				return builder.From(items...).Build()
			}

			t.Run("Can build", func(t *testing.T) {
				// ========= [A]rrange =========
				builder := deque.NewBuiltinBuilder[int]()
				builder.BackedBy(tc.backedBy)
				d := builder.Build()
				// ========= [A]ssert  =========
				must.Eq(t, 0, d.Len())
			})

			t.Run("Can build from items", func(t *testing.T) {
				// ========= [A]rrange =========
				d := newDeque(1, 2, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 3, d.Len())
				must.True(t, d.Contains(1))
				must.True(t, d.Contains(2))
				must.True(t, d.Contains(3))
			})

			t.Run("Can build from items with duplicates", func(t *testing.T) {
				// ========= [A]rrange =========
				d := newDeque(1, 2, 2, 3, 3, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 6, d.Len())
			})

			t.Run("Collection methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				d := newDeque(1, 2, 3)

				// SCENARIO: Len
				t.Run("Len", func(t *testing.T) {
					// ========= [A]ct     =========
					length := d.Len()
					// ========= [A]ssert  =========
					must.Eq(t, 3, length)
				})

				// SCENARIO: Contains
				t.Run("Contains - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Contains(4)
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Contains - true", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Contains(2)
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: IsEmpty
				t.Run("IsEmpty - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.IsEmpty()
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: Clear
				t.Run("Clear", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					d.Clear()
					// ========= [A]ssert  =========
					must.Eq(t, 0, d.Len())
					must.True(t, d.IsEmpty())
				})
			})

			t.Run("Aggregate methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				d := newDeque(1, 2, 3, 4, 5)

				// SCENARIO: Any
				t.Run("Any - False", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Any(func(item int) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Any - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Any(func(item int) bool {
						return item > 3
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: Count
				t.Run("Count", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Count(func(item int) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, actual)
				})

				// SCENARIO: Every
				t.Run("Every - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Every(func(item int) bool {
						return item < 10
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				t.Run("Every - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Every(func(item int) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: ForEach
				t.Run("ForEach", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					d.ForEach(func(item int) {
						count++
					})
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})
			})

			t.Run("Deque operations work", func(t *testing.T) {
				// SCENARIO: PushFront
				t.Run("PushFront", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					d.PushFront(0)
					// ========= [A]ssert  =========
					must.Eq(t, 4, d.Len())
					must.True(t, d.PeekFront().IsSome())
					must.Eq(t, 0, d.PeekFront().Unwrap())
				})

				// SCENARIO: PushBack
				t.Run("PushBack", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					d.PushBack(4)
					// ========= [A]ssert  =========
					must.Eq(t, 4, d.Len())
					must.True(t, d.PeekBack().IsSome())
					must.Eq(t, 4, d.PeekBack().Unwrap())
				})

				// SCENARIO: PopFront
				t.Run("PopFront - non-empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					popped := d.PopFront()
					// ========= [A]ssert  =========
					must.True(t, popped.IsSome())
					must.Eq(t, 1, popped.Unwrap())
					must.Eq(t, 2, d.Len())
				})

				t.Run("PopFront - empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					popped := d.PopFront()
					// ========= [A]ssert  =========
					must.True(t, popped.IsNone())
				})

				// SCENARIO: PopBack
				t.Run("PopBack - non-empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					popped := d.PopBack()
					// ========= [A]ssert  =========
					must.True(t, popped.IsSome())
					must.Eq(t, 3, popped.Unwrap())
					must.Eq(t, 2, d.Len())
				})

				t.Run("PopBack - empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					popped := d.PopBack()
					// ========= [A]ssert  =========
					must.True(t, popped.IsNone())
				})

				// SCENARIO: PeekFront
				t.Run("PeekFront - non-empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					peeked := d.PeekFront()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsSome())
					must.Eq(t, 1, peeked.Unwrap())
					must.Eq(t, 3, d.Len()) // Length unchanged
				})

				t.Run("PeekFront - empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					peeked := d.PeekFront()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsNone())
				})

				// SCENARIO: PeekBack
				t.Run("PeekBack - non-empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					peeked := d.PeekBack()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsSome())
					must.Eq(t, 3, peeked.Unwrap())
					must.Eq(t, 3, d.Len()) // Length unchanged
				})

				t.Run("PeekBack - empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					peeked := d.PeekBack()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsNone())
				})

				// SCENARIO: Double-ended operations
				t.Run("Double-ended operations", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					d.PushBack(1)
					d.PushBack(2)
					d.PushFront(0)
					d.PushBack(3)
					d.PushFront(-1)
					// ========= [A]ssert  =========
					// Deque should be: -1, 0, 1, 2, 3
					must.Eq(t, -1, d.PopFront().Unwrap())
					must.Eq(t, 3, d.PopBack().Unwrap())
					must.Eq(t, 0, d.PopFront().Unwrap())
					must.Eq(t, 2, d.PopBack().Unwrap())
					must.Eq(t, 1, d.PopFront().Unwrap())
					must.True(t, d.PopFront().IsNone())
				})
			})

			t.Run("Sequence methods work", func(t *testing.T) {
				// SCENARIO: Values
				t.Run("Values", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range d.Values() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: All
				t.Run("All", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range d.All() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: Find
				t.Run("Find - found", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					found := d.Find(func(item int) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsSome())
					must.Eq(t, 2, found.Unwrap())
				})

				t.Run("Find - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					found := d.Find(func(item int) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsNone())
				})

				// SCENARIO: FindIndex
				t.Run("FindIndex - found", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					idx := d.FindIndex(func(item int) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsSome())
					must.Eq(t, 1, idx.Unwrap())
				})

				t.Run("FindIndex - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					idx := d.FindIndex(func(item int) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsNone())
				})

				// SCENARIO: Get
				t.Run("Get - valid index", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					item := d.Get(1)
					// ========= [A]ssert  =========
					must.True(t, item.IsSome())
					must.Eq(t, 2, item.Unwrap())
				})

				t.Run("Get - invalid index", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					item := d.Get(10)
					// ========= [A]ssert  =========
					must.True(t, item.IsNone())
				})

				// SCENARIO: Retain
				t.Run("Retain", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3, 4, 5)
					// ========= [A]ct     =========
					d.Retain(func(item int) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, d.Len())
					must.True(t, d.Contains(2))
					must.True(t, d.Contains(4))
				})

				// SCENARIO: Sort
				t.Run("Sort", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(3, 1, 4, 1, 5, 9, 2, 6)
					// ========= [A]ct     =========
					d.Sort(func(a, b int) compare.Order {
						if a < b {
							return compare.OrderLess
						} else if a > b {
							return compare.OrderGreater
						}
						return compare.OrderEqual
					})
					// ========= [A]ssert  =========
					slice := d.ToSlice()
					must.Eq(t, []int{1, 1, 2, 3, 4, 5, 6, 9}, slice)
				})

				// SCENARIO: ToSlice
				t.Run("ToSlice", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					slice := d.ToSlice()
					// ========= [A]ssert  =========
					must.Eq(t, []int{1, 2, 3}, slice)
				})
			})
		})
	}
}

func TestDequeFromComparable(t *testing.T) {
	testCases := []struct {
		name     string
		backedBy deque.DequeBackedBy
	}{
		{"Slice", deque.DequeBackedBySlice},
		{"DoublyLinkedList", deque.DequeBackedByDoublyLinkedList},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newDeque := func(items ...ComparableInt) sequence.Deque[ComparableInt] {
				builder := deque.NewComparableBuilder[ComparableInt]()
				builder.BackedBy(tc.backedBy)
				return builder.From(items...).Build()
			}

			t.Run("Can build", func(t *testing.T) {
				// ========= [A]rrange =========
				builder := deque.NewComparableBuilder[ComparableInt]()
				builder.BackedBy(tc.backedBy)
				d := builder.Build()
				// ========= [A]ssert  =========
				must.Eq(t, 0, d.Len())
			})

			t.Run("Can build from items", func(t *testing.T) {
				// ========= [A]rrange =========
				d := newDeque(1, 2, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 3, d.Len())
				must.True(t, d.Contains(1))
				must.True(t, d.Contains(2))
				must.True(t, d.Contains(3))
			})

			t.Run("Can build from items with duplicates", func(t *testing.T) {
				// ========= [A]rrange =========
				d := newDeque(1, 2, 2, 3, 3, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 6, d.Len())
			})

			t.Run("Collection methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				d := newDeque(1, 2, 3)

				// SCENARIO: Len
				t.Run("Len", func(t *testing.T) {
					// ========= [A]ct     =========
					length := d.Len()
					// ========= [A]ssert  =========
					must.Eq(t, 3, length)
				})

				// SCENARIO: Contains
				t.Run("Contains - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Contains(4)
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Contains - true", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Contains(2)
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: IsEmpty
				t.Run("IsEmpty - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.IsEmpty()
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: Clear
				t.Run("Clear", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					d.Clear()
					// ========= [A]ssert  =========
					must.Eq(t, 0, d.Len())
					must.True(t, d.IsEmpty())
				})
			})

			t.Run("Aggregate methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				d := newDeque(1, 2, 3, 4, 5)

				// SCENARIO: Any
				t.Run("Any - False", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Any(func(item ComparableInt) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Any - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Any(func(item ComparableInt) bool {
						return item > 3
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: Count
				t.Run("Count", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Count(func(item ComparableInt) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, actual)
				})

				// SCENARIO: Every
				t.Run("Every - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Every(func(item ComparableInt) bool {
						return item < 10
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				t.Run("Every - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := d.Every(func(item ComparableInt) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: ForEach
				t.Run("ForEach", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					d.ForEach(func(item ComparableInt) {
						count++
					})
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})
			})

			t.Run("Deque operations work", func(t *testing.T) {
				// SCENARIO: PushFront
				t.Run("PushFront", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					d.PushFront(0)
					// ========= [A]ssert  =========
					must.Eq(t, 4, d.Len())
					must.True(t, d.PeekFront().IsSome())
					must.Eq(t, ComparableInt(0), d.PeekFront().Unwrap())
				})

				// SCENARIO: PushBack
				t.Run("PushBack", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					d.PushBack(4)
					// ========= [A]ssert  =========
					must.Eq(t, 4, d.Len())
					must.True(t, d.PeekBack().IsSome())
					must.Eq(t, ComparableInt(4), d.PeekBack().Unwrap())
				})

				// SCENARIO: PopFront
				t.Run("PopFront - non-empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					popped := d.PopFront()
					// ========= [A]ssert  =========
					must.True(t, popped.IsSome())
					must.Eq(t, ComparableInt(1), popped.Unwrap())
					must.Eq(t, 2, d.Len())
				})

				t.Run("PopFront - empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					popped := d.PopFront()
					// ========= [A]ssert  =========
					must.True(t, popped.IsNone())
				})

				// SCENARIO: PopBack
				t.Run("PopBack - non-empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					popped := d.PopBack()
					// ========= [A]ssert  =========
					must.True(t, popped.IsSome())
					must.Eq(t, ComparableInt(3), popped.Unwrap())
					must.Eq(t, 2, d.Len())
				})

				t.Run("PopBack - empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					popped := d.PopBack()
					// ========= [A]ssert  =========
					must.True(t, popped.IsNone())
				})

				// SCENARIO: PeekFront
				t.Run("PeekFront - non-empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					peeked := d.PeekFront()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsSome())
					must.Eq(t, ComparableInt(1), peeked.Unwrap())
					must.Eq(t, 3, d.Len()) // Length unchanged
				})

				t.Run("PeekFront - empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					peeked := d.PeekFront()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsNone())
				})

				// SCENARIO: PeekBack
				t.Run("PeekBack - non-empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					peeked := d.PeekBack()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsSome())
					must.Eq(t, ComparableInt(3), peeked.Unwrap())
					must.Eq(t, 3, d.Len()) // Length unchanged
				})

				t.Run("PeekBack - empty deque", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					peeked := d.PeekBack()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsNone())
				})

				// SCENARIO: Double-ended operations
				t.Run("Double-ended operations", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque()
					// ========= [A]ct     =========
					d.PushBack(1)
					d.PushBack(2)
					d.PushFront(0)
					d.PushBack(3)
					d.PushFront(-1)
					// ========= [A]ssert  =========
					// Deque should be: -1, 0, 1, 2, 3
					must.Eq(t, ComparableInt(-1), d.PopFront().Unwrap())
					must.Eq(t, ComparableInt(3), d.PopBack().Unwrap())
					must.Eq(t, ComparableInt(0), d.PopFront().Unwrap())
					must.Eq(t, ComparableInt(2), d.PopBack().Unwrap())
					must.Eq(t, ComparableInt(1), d.PopFront().Unwrap())
					must.True(t, d.PopFront().IsNone())
				})
			})

			t.Run("Sequence methods work", func(t *testing.T) {
				// SCENARIO: Values
				t.Run("Values", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range d.Values() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: All
				t.Run("All", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range d.All() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: Find
				t.Run("Find - found", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					found := d.Find(func(item ComparableInt) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsSome())
					must.Eq(t, ComparableInt(2), found.Unwrap())
				})

				t.Run("Find - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					found := d.Find(func(item ComparableInt) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsNone())
				})

				// SCENARIO: FindIndex
				t.Run("FindIndex - found", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					idx := d.FindIndex(func(item ComparableInt) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsSome())
					must.Eq(t, 1, idx.Unwrap())
				})

				t.Run("FindIndex - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					idx := d.FindIndex(func(item ComparableInt) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsNone())
				})

				// SCENARIO: Get
				t.Run("Get - valid index", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					item := d.Get(1)
					// ========= [A]ssert  =========
					must.True(t, item.IsSome())
					must.Eq(t, ComparableInt(2), item.Unwrap())
				})

				t.Run("Get - invalid index", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					item := d.Get(10)
					// ========= [A]ssert  =========
					must.True(t, item.IsNone())
				})

				// SCENARIO: Retain
				t.Run("Retain", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3, 4, 5)
					// ========= [A]ct     =========
					d.Retain(func(item ComparableInt) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, d.Len())
					must.True(t, d.Contains(2))
					must.True(t, d.Contains(4))
				})

				// SCENARIO: Sort
				t.Run("Sort", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(3, 1, 4, 1, 5, 9, 2, 6)
					// ========= [A]ct     =========
					d.Sort(func(a, b ComparableInt) compare.Order {
						if a < b {
							return compare.OrderLess
						} else if a > b {
							return compare.OrderGreater
						}
						return compare.OrderEqual
					})
					// ========= [A]ssert  =========
					slice := d.ToSlice()
					must.Eq(t, []ComparableInt{1, 1, 2, 3, 4, 5, 6, 9}, slice)
				})

				// SCENARIO: ToSlice
				t.Run("ToSlice", func(t *testing.T) {
					// ========= [A]rrange =========
					d := newDeque(1, 2, 3)
					// ========= [A]ct     =========
					slice := d.ToSlice()
					// ========= [A]ssert  =========
					must.Eq(t, []ComparableInt{1, 2, 3}, slice)
				})
			})
		})
	}
}
