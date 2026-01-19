package stack_test

import (
	"slices"
	"testing"

	"github.com/shoenig/test/must"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/compare"
	"codeberg.org/yaadata/bina/sequence/stack"
)

type ComparableInt int

func (c ComparableInt) Equal(other ComparableInt) bool {
	return c == other
}

func TestStackFromBuiltin(t *testing.T) {
	testCases := []struct {
		name     string
		backedBy stack.StackBackedBy
	}{
		{"Slice", stack.StackBackedBySlice},
		{"SinglyLinkedList", stack.StackBackedBySinglyLinkedList},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newStack := func(items ...int) collection.Stack[int] {
				builder := stack.NewBuiltinBuilder[int]()
				builder.BackedBy(tc.backedBy)
				return builder.From(items...).Build()
			}

			t.Run("Can build", func(t *testing.T) {
				// ========= [A]rrange =========
				builder := stack.NewBuiltinBuilder[int]()
				builder.BackedBy(tc.backedBy)
				s := builder.Build()
				// ========= [A]ssert  =========
				must.Eq(t, 0, s.Len())
			})

			t.Run("Can build from items", func(t *testing.T) {
				// ========= [A]rrange =========
				s := newStack(1, 2, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 3, s.Len())
				must.True(t, s.Contains(1))
				must.True(t, s.Contains(2))
				must.True(t, s.Contains(3))
			})

			t.Run("Can build from items with duplicates", func(t *testing.T) {
				// ========= [A]rrange =========
				s := newStack(1, 2, 2, 3, 3, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 6, s.Len())
			})

			t.Run("Collection methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				s := newStack(1, 2, 3)

				// SCENARIO: Len
				t.Run("Len", func(t *testing.T) {
					// ========= [A]ct     =========
					length := s.Len()
					// ========= [A]ssert  =========
					must.Eq(t, 3, length)
				})

				// SCENARIO: Contains
				t.Run("Contains - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Contains(4)
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Contains - true", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Contains(2)
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: IsEmpty
				t.Run("IsEmpty - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.IsEmpty()
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: Clear
				t.Run("Clear", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					s.Clear()
					// ========= [A]ssert  =========
					must.Eq(t, 0, s.Len())
					must.True(t, s.IsEmpty())
				})
			})

			t.Run("Aggregate methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				s := newStack(1, 2, 3, 4, 5)

				// SCENARIO: Any
				t.Run("Any - False", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Any(func(item int) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Any - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Any(func(item int) bool {
						return item > 3
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: Count
				t.Run("Count", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Count(func(item int) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, actual)
				})

				// SCENARIO: Every
				t.Run("Every - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Every(func(item int) bool {
						return item < 10
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				t.Run("Every - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Every(func(item int) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: ForEach
				t.Run("ForEach", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					s.ForEach(func(item int) {
						count++
					})
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})
			})

			t.Run("Stack operations work", func(t *testing.T) {
				// SCENARIO: Push
				t.Run("Push", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					s.Push(4)
					// ========= [A]ssert  =========
					must.Eq(t, 4, s.Len())
					must.True(t, s.Peek().IsSome())
					must.Eq(t, 4, s.Peek().Unwrap())
				})

				// SCENARIO: Pop
				t.Run("Pop - non-empty stack", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					popped := s.Pop()
					// ========= [A]ssert  =========
					must.True(t, popped.IsSome())
					must.Eq(t, 3, popped.Unwrap())
					must.Eq(t, 2, s.Len())
				})

				t.Run("Pop - empty stack", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack()
					// ========= [A]ct     =========
					popped := s.Pop()
					// ========= [A]ssert  =========
					must.True(t, popped.IsNone())
				})

				// SCENARIO: Peek
				t.Run("Peek - non-empty stack", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					peeked := s.Peek()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsSome())
					must.Eq(t, 3, peeked.Unwrap())
					must.Eq(t, 3, s.Len()) // Length unchanged
				})

				t.Run("Peek - empty stack", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack()
					// ========= [A]ct     =========
					peeked := s.Peek()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsNone())
				})

				// SCENARIO: LIFO order
				t.Run("LIFO order", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack()
					// ========= [A]ct     =========
					s.Push(1)
					s.Push(2)
					s.Push(3)
					// ========= [A]ssert  =========
					must.Eq(t, 3, s.Pop().Unwrap())
					must.Eq(t, 2, s.Pop().Unwrap())
					must.Eq(t, 1, s.Pop().Unwrap())
					must.True(t, s.Pop().IsNone())
				})
			})

			t.Run("Sequence methods work", func(t *testing.T) {
				// SCENARIO: Values
				t.Run("Values", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range s.Values() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: All
				t.Run("All", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range s.All() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: Find
				t.Run("Find - found", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					found := s.Find(func(item int) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsSome())
					must.Eq(t, 2, found.Unwrap())
				})

				t.Run("Find - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					found := s.Find(func(item int) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsNone())
				})

				// SCENARIO: FindIndex
				t.Run("FindIndex - found", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					idx := s.FindIndex(func(item int) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsSome())
					must.Eq(t, 1, idx.Unwrap())
				})

				t.Run("FindIndex - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					idx := s.FindIndex(func(item int) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsNone())
				})

				// SCENARIO: Get
				t.Run("Get - valid index", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					item := s.Get(1)
					// ========= [A]ssert  =========
					must.True(t, item.IsSome())
					must.Eq(t, 2, item.Unwrap())
				})

				t.Run("Get - invalid index", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					item := s.Get(10)
					// ========= [A]ssert  =========
					must.True(t, item.IsNone())
				})

				// SCENARIO: Retain
				t.Run("Retain", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3, 4, 5)
					// ========= [A]ct     =========
					s.Retain(func(item int) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, s.Len())
					must.True(t, s.Contains(2))
					must.True(t, s.Contains(4))
				})

				// SCENARIO: Sort
				t.Run("Sort", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(3, 1, 4, 1, 5, 9, 2, 6)
					// ========= [A]ct     =========
					s.Sort(func(a, b int) compare.Order {
						if a < b {
							return compare.OrderLess
						} else if a > b {
							return compare.OrderGreater
						}
						return compare.OrderEqual
					})
					// ========= [A]ssert  =========
					must.Eq(t, []int{1, 1, 2, 3, 4, 5, 6, 9}, slices.Collect(s.Values()))
				})
			})
		})
	}
}

func TestStackFromComparable(t *testing.T) {
	testCases := []struct {
		name     string
		backedBy stack.StackBackedBy
	}{
		{"Slice", stack.StackBackedBySlice},
		{"SinglyLinkedList", stack.StackBackedBySinglyLinkedList},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newStack := func(items ...ComparableInt) collection.Stack[ComparableInt] {
				builder := stack.NewComparableBuilder[ComparableInt]()
				builder.BackedBy(tc.backedBy)
				return builder.From(items...).Build()
			}

			t.Run("Can build", func(t *testing.T) {
				// ========= [A]rrange =========
				builder := stack.NewComparableBuilder[ComparableInt]()
				builder.BackedBy(tc.backedBy)
				s := builder.Build()
				// ========= [A]ssert  =========
				must.Eq(t, 0, s.Len())
			})

			t.Run("Can build from items", func(t *testing.T) {
				// ========= [A]rrange =========
				s := newStack(1, 2, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 3, s.Len())
				must.True(t, s.Contains(1))
				must.True(t, s.Contains(2))
				must.True(t, s.Contains(3))
			})

			t.Run("Can build from items with duplicates", func(t *testing.T) {
				// ========= [A]rrange =========
				s := newStack(1, 2, 2, 3, 3, 3)
				// ========= [A]ssert  =========
				must.Eq(t, 6, s.Len())
			})

			t.Run("Collection methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				s := newStack(1, 2, 3)

				// SCENARIO: Len
				t.Run("Len", func(t *testing.T) {
					// ========= [A]ct     =========
					length := s.Len()
					// ========= [A]ssert  =========
					must.Eq(t, 3, length)
				})

				// SCENARIO: Contains
				t.Run("Contains - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Contains(4)
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Contains - true", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Contains(2)
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: IsEmpty
				t.Run("IsEmpty - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.IsEmpty()
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: Clear
				t.Run("Clear", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					s.Clear()
					// ========= [A]ssert  =========
					must.Eq(t, 0, s.Len())
					must.True(t, s.IsEmpty())
				})
			})

			t.Run("Aggregate methods work", func(t *testing.T) {
				// ========= [A]rrange =========
				s := newStack(1, 2, 3, 4, 5)

				// SCENARIO: Any
				t.Run("Any - False", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Any(func(item ComparableInt) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})
				t.Run("Any - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Any(func(item ComparableInt) bool {
						return item > 3
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				// SCENARIO: Count
				t.Run("Count", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Count(func(item ComparableInt) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, actual)
				})

				// SCENARIO: Every
				t.Run("Every - True", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Every(func(item ComparableInt) bool {
						return item < 10
					})
					// ========= [A]ssert  =========
					must.True(t, actual)
				})

				t.Run("Every - false", func(t *testing.T) {
					// ========= [A]ct     =========
					actual := s.Every(func(item ComparableInt) bool {
						return item > 10
					})
					// ========= [A]ssert  =========
					must.False(t, actual)
				})

				// SCENARIO: ForEach
				t.Run("ForEach", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					s.ForEach(func(item ComparableInt) {
						count++
					})
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})
			})

			t.Run("Stack operations work", func(t *testing.T) {
				// SCENARIO: Push
				t.Run("Push", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					s.Push(4)
					// ========= [A]ssert  =========
					must.Eq(t, 4, s.Len())
					must.True(t, s.Peek().IsSome())
					must.Eq(t, ComparableInt(4), s.Peek().Unwrap())
				})

				// SCENARIO: Pop
				t.Run("Pop - non-empty stack", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					popped := s.Pop()
					// ========= [A]ssert  =========
					must.True(t, popped.IsSome())
					must.Eq(t, ComparableInt(3), popped.Unwrap())
					must.Eq(t, 2, s.Len())
				})

				t.Run("Pop - empty stack", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack()
					// ========= [A]ct     =========
					popped := s.Pop()
					// ========= [A]ssert  =========
					must.True(t, popped.IsNone())
				})

				// SCENARIO: Peek
				t.Run("Peek - non-empty stack", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					peeked := s.Peek()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsSome())
					must.Eq(t, ComparableInt(3), peeked.Unwrap())
					must.Eq(t, 3, s.Len()) // Length unchanged
				})

				t.Run("Peek - empty stack", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack()
					// ========= [A]ct     =========
					peeked := s.Peek()
					// ========= [A]ssert  =========
					must.True(t, peeked.IsNone())
				})

				// SCENARIO: LIFO order
				t.Run("LIFO order", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack()
					// ========= [A]ct     =========
					s.Push(1)
					s.Push(2)
					s.Push(3)
					// ========= [A]ssert  =========
					must.Eq(t, ComparableInt(3), s.Pop().Unwrap())
					must.Eq(t, ComparableInt(2), s.Pop().Unwrap())
					must.Eq(t, ComparableInt(1), s.Pop().Unwrap())
					must.True(t, s.Pop().IsNone())
				})
			})

			t.Run("Sequence methods work", func(t *testing.T) {
				// SCENARIO: Values
				t.Run("Values", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range s.Values() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: All
				t.Run("All", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					count := 0
					for range s.All() {
						count++
					}
					// ========= [A]ssert  =========
					must.Eq(t, 3, count)
				})

				// SCENARIO: Find
				t.Run("Find - found", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					found := s.Find(func(item ComparableInt) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsSome())
					must.Eq(t, ComparableInt(2), found.Unwrap())
				})

				t.Run("Find - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					found := s.Find(func(item ComparableInt) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, found.IsNone())
				})

				// SCENARIO: FindIndex
				t.Run("FindIndex - found", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					idx := s.FindIndex(func(item ComparableInt) bool {
						return item == 2
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsSome())
					must.Eq(t, 1, idx.Unwrap())
				})

				t.Run("FindIndex - not found", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					idx := s.FindIndex(func(item ComparableInt) bool {
						return item == 10
					})
					// ========= [A]ssert  =========
					must.True(t, idx.IsNone())
				})

				// SCENARIO: Get
				t.Run("Get - valid index", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					item := s.Get(1)
					// ========= [A]ssert  =========
					must.True(t, item.IsSome())
					must.Eq(t, ComparableInt(2), item.Unwrap())
				})

				t.Run("Get - invalid index", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3)
					// ========= [A]ct     =========
					item := s.Get(10)
					// ========= [A]ssert  =========
					must.True(t, item.IsNone())
				})

				// SCENARIO: Retain
				t.Run("Retain", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(1, 2, 3, 4, 5)
					// ========= [A]ct     =========
					s.Retain(func(item ComparableInt) bool {
						return item%2 == 0
					})
					// ========= [A]ssert  =========
					must.Eq(t, 2, s.Len())
					must.True(t, s.Contains(2))
					must.True(t, s.Contains(4))
				})

				// SCENARIO: Sort
				t.Run("Sort", func(t *testing.T) {
					// ========= [A]rrange =========
					s := newStack(3, 1, 4, 1, 5, 9, 2, 6)
					// ========= [A]ct     =========
					s.Sort(func(a, b ComparableInt) compare.Order {
						if a < b {
							return compare.OrderLess
						} else if a > b {
							return compare.OrderGreater
						}
						return compare.OrderEqual
					})
					// ========= [A]ssert  =========
					must.Eq(t, []ComparableInt{1, 1, 2, 3, 4, 5, 6, 9}, slices.Collect(s.Values()))
				})
			})
		})
	}
}
