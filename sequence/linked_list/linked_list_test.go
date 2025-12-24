package linkedlist_test

import (
	"testing"

	"github.com/shoenig/test/must"

	"codeberg.org/yaadata/bina/core/compare"
	linkedlist "codeberg.org/yaadata/bina/sequence/linked_list"
)

func TestLinkedListFromBuiltin(t *testing.T) {
	t.Run("Can build", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 0, sequence.Len())
	})

	t.Run("Can build with capacity (no-op for linked list)", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			Capacity(10).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 0, sequence.Len())
	})

	t.Run("Can build from items", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 3, sequence.Len())
		must.Eq(t, []int{1, 2, 3}, sequence.ToSlice())
	})

	t.Run("Can enumerate", func(t *testing.T) {
		// ========= [A]rrange =========
		original := []int{1, 2, 3}
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(original...).
			Build()
		// ========= [A]ct     =========
		var actual []int
		for _, value := range sequence.Enumerate() {
			actual = append(actual, value)
		}
		// ========= [A]ssert  =========
		must.Eq(t, original, actual)
	})

	t.Run("Can gather all", func(t *testing.T) {
		// ========= [A]rrange =========
		original := []int{1, 2, 3}
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(original...).
			Build()
		// ========= [A]ct     =========
		var actual []int
		for value := range sequence.All() {
			actual = append(actual, value)
		}
		// ========= [A]ssert  =========
		must.Eq(t, original, actual)
	})

	t.Run("Can Extend", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		// ========= [A]ct     =========
		sequence.Extend(4, 5, 6)
		// ========= [A]ssert  =========
		must.Eq(t, 6, sequence.Len())
		must.Eq(t, []int{1, 2, 3, 4, 5, 6}, sequence.ToSlice())
	})

	t.Run("Can Extend From Sequence", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		other := linkedlist.NewBuiltinBuilder[int]().
			From(4, 5, 6).
			Build()
		// ========= [A]ct     =========
		sequence.ExtendFromSequence(other)
		// ========= [A]ssert  =========
		must.Eq(t, 6, sequence.Len())
		must.Eq(t, []int{1, 2, 3, 4, 5, 6}, sequence.ToSlice())
	})

	t.Run("Find returns Option", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()

		// SCENARIO: no item matches predicate
		t.Run("No match", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Find(func(item int) bool {
				return item == 4
			})
			// ========= [A]ssert  =========
			must.True(t, actual.IsNone())
		})

		// SCENARIO: an item matches the predicate
		t.Run("Returns first match", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Find(func(item int) bool {
				return item == 1 || item == 3
			})
			// ========= [A]ssert  =========
			must.True(t, actual.IsSome())
			must.Eq(t, 1, actual.Unwrap())
		})
	})

	t.Run("FindIndex returns Option", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()

		// SCENARIO: no item matches predicate
		t.Run("No match", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.FindIndex(func(item int) bool {
				return item == 4
			})
			// ========= [A]ssert  =========
			must.True(t, actual.IsNone())
		})

		// SCENARIO: an item matches the predicate
		t.Run("Returns first match", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.FindIndex(func(item int) bool {
				return item == 1 || item == 3
			})
			// ========= [A]ssert  =========
			must.True(t, actual.IsSome())
			must.Eq(t, 0, actual.Unwrap())
		})
	})

	t.Run("Can Get", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()

		// SCENARIO: no item at index
		t.Run("No item at index", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Get(3)
			// ========= [A]ssert  =========
			must.True(t, actual.IsNone())
		})

		// SCENARIO: item at index
		t.Run("Item at index", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Get(1)
			// ========= [A]ssert  =========
			must.True(t, actual.IsSome())
			must.Eq(t, 2, actual.Unwrap())
		})
	})

	t.Run("Can Insert", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 4).
			Build()
		// ========= [A]ct     =========
		sequence.Insert(2, 3)
		// ========= [A]ssert  =========
		must.Eq(t, 4, sequence.Len())
		must.Eq(t, []int{1, 2, 3, 4}, sequence.ToSlice())
	})

	t.Run("Can RemoveAt", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3, 4).
			Build()
		// ========= [A]ct     =========
		removed := sequence.RemoveAt(2)
		// ========= [A]ssert  =========
		must.True(t, removed.IsSome())
		must.Eq(t, 3, removed.Unwrap())
		must.Eq(t, 3, sequence.Len())
		must.Eq(t, []int{1, 2, 4}, sequence.ToSlice())
	})

	t.Run("Can RemoveAt - returns None", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3, 4).
			Build()
		// ========= [A]ct     =========
		removed := sequence.RemoveAt(20)
		// ========= [A]ssert  =========
		must.True(t, removed.IsNone())
		must.Eq(t, 4, sequence.Len())
	})

	t.Run("Can Retain", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3, 4, 5, 6).
			Build()
		// ========= [A]ct     =========
		sequence.Retain(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, 3, sequence.Len())
		must.Eq(t, []int{2, 4, 6}, sequence.ToSlice())
	})

	t.Run("Can Sort", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(5, 2, 6, 1, 4, 3).
			Build()
		// ========= [A]ct     =========
		sequence.Sort(func(a, b int) compare.Order {
			if a < b {
				return compare.OrderLess
			} else if a > b {
				return compare.OrderGreater
			} else {
				return compare.OrderEqual
			}
		})
		// ========= [A]ssert  =========
		must.Eq(t, 6, sequence.Len())
		must.Eq(t, []int{1, 2, 3, 4, 5, 6}, sequence.ToSlice())
	})

	t.Run("Collection methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()

		// SCENARIO: Len
		t.Run("Len", func(t *testing.T) {
			// ========= [A]ct     =========
			length := sequence.Len()
			// ========= [A]ssert  =========
			must.Eq(t, 3, length)
		})

		// SCENARIO: Contains
		t.Run("Contains - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Contains(4)
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("Contains - true", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Contains(2)
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: IsEmpty
		t.Run("IsEmpty - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.IsEmpty()
			// ========= [A]ssert  =========
			must.False(t, actual)
		})

		// SCENARIO: Clear
		t.Run("Clear", func(t *testing.T) {
			// ========= [A]ct     =========
			sequence.Clear()
			// ========= [A]ssert  =========
			must.Eq(t, 0, sequence.Len())
			must.True(t, sequence.IsEmpty())
		})
	})

	t.Run("Aggregate methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := linkedlist.NewBuiltinBuilder[int]().
			From(1, 2, 3, 4, 5).
			Build()

		// SCENARIO: Any
		t.Run("Any - False", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Any(func(item int) bool {
				return item > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("Any - True", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Any(func(item int) bool {
				return item > 3
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: Count
		t.Run("Count", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Count(func(item int) bool {
				return item%2 == 0
			})
			// ========= [A]ssert  =========
			must.Eq(t, 2, actual)
		})

		// SCENARIO: Every
		t.Run("Every - True", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Every(func(item int) bool {
				return item < 10
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		t.Run("Every - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := sequence.Every(func(item int) bool {
				return item > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})

		// SCENARIO: For each
		t.Run("ForEach", func(t *testing.T) {
			// ========= [A]ct     =========
			summation := 0
			sequence.ForEach(func(item int) {
				summation += item
			})
			// ========= [A]ssert  =========
			must.Eq(t, 15, summation)
		})
	})

	t.Run("LinkedList specific methods", func(t *testing.T) {
		t.Run("Can Prepend", func(t *testing.T) {
			// ========= [A]rrange =========
			sequence := linkedlist.NewBuiltinBuilder[int]().
				From(2, 3, 4).
				Build()
			// ========= [A]ct     =========
			sequence.Prepend(1)
			// ========= [A]ssert  =========
			must.Eq(t, 4, sequence.Len())
			must.Eq(t, []int{1, 2, 3, 4}, sequence.ToSlice())
			head := sequence.Head()
			must.Eq(t, 1, head.Unwrap().Value())
		})

		t.Run("Can Prepend to empty list", func(t *testing.T) {
			// ========= [A]rrange =========
			sequence := linkedlist.NewBuiltinBuilder[int]().
				Build()
			// ========= [A]ct     =========
			sequence.Prepend(1)
			// ========= [A]ssert  =========
			must.Eq(t, 1, sequence.Len())
			must.Eq(t, []int{1}, sequence.ToSlice())
		})

		t.Run("Can GetNodeAt", func(t *testing.T) {
			// ========= [A]rrange =========
			sequence := linkedlist.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()

			// SCENARIO: node at index
			t.Run("Node at index", func(t *testing.T) {
				// ========= [A]ct     =========
				node := sequence.GetNodeAt(1)
				// ========= [A]ssert  =========
				must.True(t, node.IsSome())
				must.Eq(t, 2, node.Unwrap().Value())
			})

			// SCENARIO: no node at index
			t.Run("No node at index", func(t *testing.T) {
				// ========= [A]ct     =========
				node := sequence.GetNodeAt(10)
				// ========= [A]ssert  =========
				must.True(t, node.IsNone())
			})
		})

		t.Run("Can modify node value via SetValue", func(t *testing.T) {
			// ========= [A]rrange =========
			sequence := linkedlist.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			node := sequence.GetNodeAt(1)
			node.Unwrap().SetValue(20)
			// ========= [A]ssert  =========
			must.Eq(t, []int{1, 20, 3}, sequence.ToSlice())
		})

		t.Run("Can Append", func(t *testing.T) {
			// ========= [A]rrange =========
			sequence := linkedlist.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			sequence.Append(4)
			// ========= [A]ssert  =========
			must.Eq(t, 4, sequence.Len())
			must.Eq(t, []int{1, 2, 3, 4}, sequence.ToSlice())
			tail := sequence.Tail()
			must.True(t, tail.IsSome())
			must.Eq(t, 4, tail.Unwrap().Value())
		})
	})
}
