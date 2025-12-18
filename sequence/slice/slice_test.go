package slice_test

import (
	"slices"
	"testing"

	"github.com/shoenig/test/must"

	"github.com/yaadata/bina/core/compare"
	"github.com/yaadata/bina/sequence/slice"
)

func TestSliceFromBuiltin(t *testing.T) {
	t.Run("Can build", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			Capacity(10).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 0, sequence.Len())
		must.Eq(t, 10, cap(sequence.ToSlice()))
	})

	t.Run("Can build without capacity", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			Build()

		// ========= [A]ssert  =========
		must.Eq(t, 0, sequence.Len())
		must.Eq(t, 0, cap(sequence.ToSlice()))
	})

	t.Run("Can build from items", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 3, sequence.Len())
		must.Eq(t, []int{1, 2, 3}, sequence.ToSlice())
	})

	t.Run("Can enumerate", func(t *testing.T) {
		// ========= [A]rrange =========
		original := []int{1, 2, 3}
		sequence := slice.NewBuiltinBuilder[int]().
			From(original...).
			Build()
		// ========= [A]ssert  =========
		for index, value := range sequence.Enumerate() {
			must.Eq(t, original[index], value)
		}
	})

	t.Run("Can gather all", func(t *testing.T) {
		// ========= [A]rrange =========
		original := []int{1, 2, 3}
		sequence := slice.NewBuiltinBuilder[int]().
			From(original...).
			Build()
		// ========= [A]ct     =========
		var actual []int
		for value := range sequence.All() {
			actual = append(actual, value)
		}
		// ========= [A]ssert  =========
		must.True(t, slices.Equal(original, actual))
	})

	t.Run("Can Extend", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		// ========= [A]ct     =========
		sequence.Extend(4, 5, 6)
		// ========= [A]ssert  =========
		must.True(t, slices.Equal([]int{1, 2, 3, 4, 5, 6}, sequence.ToSlice()))
	})

	t.Run("Can Extend From Sequence", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		other := slice.NewBuiltinBuilder[int]().
			From(4, 5, 6).
			Build()
		// ========= [A]ct     =========
		sequence.ExtendFromSequence(other)
		// ========= [A]ssert  =========
		must.True(t, slices.Equal([]int{1, 2, 3, 4, 5, 6}, sequence.ToSlice()))
	})

	t.Run("Find returns Option", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
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
		sequence := slice.NewBuiltinBuilder[int]().
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

	t.Run("Can Filter", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3, 4, 5, 6).
			Build()
		// ========= [A]ct     =========
		actual := sequence.Filter(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.True(t, slices.Equal([]int{2, 4, 6}, actual.ToSlice()))
	})

	t.Run("Can Get", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
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
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 4).
			Build()
		// ========= [A]ct     =========
		sequence.Insert(2, 3)
		// ========= [A]ssert  =========
		must.True(t, slices.Equal([]int{1, 2, 3, 4}, sequence.ToSlice()))
	})

	t.Run("Can RemoveAt", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3, 4).
			Build()
		// ========= [A]ct     =========
		sequence.RemoveAt(2)
		// ========= [A]ssert  =========
		must.True(t, slices.Equal([]int{1, 2, 4}, sequence.ToSlice()))
	})

	t.Run("Can RemoveAt - panics", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3, 4).
			Build()
		// ========= [A]ssert  =========
		must.Panic(t, func() {
			sequence.RemoveAt(20)
		})
	})

	t.Run("Can Retain", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3, 4, 5, 6).
			Build()
		// ========= [A]ct     =========
		sequence.Retain(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.True(t, slices.Equal([]int{2, 4, 6}, sequence.ToSlice()))
	})

	t.Run("Can Sort", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
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
		must.True(t, slices.Equal([]int{1, 2, 3, 4, 5, 6}, sequence.ToSlice()))
	})

	t.Run("Can Get First and Last", func(t *testing.T) {
		// SCENARIO: Empty sequence
		t.Run("Empty sequence", func(t *testing.T) {
			// ========= [A]rrange =========
			sequence := slice.NewBuiltinBuilder[int]().
				Build()
			// ========= [A]ct     =========
			first := sequence.First()
			last := sequence.Last()
			// ========= [A]ssert  =========
			must.True(t, first.IsNone())
			must.True(t, last.IsNone())
		})

		// SCENARIO: Non-empty sequence
		t.Run("Non-empty sequence", func(t *testing.T) {
			// ========= [A]rrange =========
			sequence := slice.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			first := sequence.First()
			last := sequence.Last()
			// ========= [A]ssert  =========
			must.True(t, first.IsSome())
			must.Eq(t, 1, first.Unwrap())
			must.True(t, last.IsSome())
			must.Eq(t, 3, last.Unwrap())
		})
	})

	t.Run("Collection methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
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
		sequence := slice.NewBuiltinBuilder[int]().
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
}
