package array_test

import (
	"testing"

	"github.com/shoenig/test/must"

	"codeberg.org/yaadata/bina/core/compare"
	core_range "codeberg.org/yaadata/bina/core/range"
	"codeberg.org/yaadata/bina/sequence/array"
)

func TestArrayFromBuiltin(t *testing.T) {
	// ============================================================================
	// Group 1: Building and Basic Properties
	// ============================================================================

	t.Run("Can build with size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 5, arr.Len())
	})

	t.Run("Len returns fixed size regardless of content", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(10).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 10, arr.Len())
		must.Eq(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, arr.ToSlice())
	})

	t.Run("IsEmpty on new array returns true", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		// ========= [A]ssert  =========
		must.True(t, arr.IsEmpty())
	})

	t.Run("IsEmpty returns false when array has non-zero values", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 2, 3})
		// ========= [A]ssert  =========
		must.False(t, arr.IsEmpty())
	})

	t.Run("Clear resets all elements to zero values", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 2, 3})
		// ========= [A]ct     =========
		arr.Clear()
		// ========= [A]ssert  =========
		must.Eq(t, 3, arr.Len())
		must.True(t, arr.IsEmpty())
		must.Eq(t, []int{0, 0, 0}, arr.ToSlice())
	})

	t.Run("Contains finds existing element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 2, 3})
		// ========= [A]ssert  =========
		must.True(t, arr.Contains(2))
	})

	t.Run("Contains returns false for missing element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 2, 3})
		// ========= [A]ssert  =========
		must.False(t, arr.Contains(4))
	})

	// ============================================================================
	// Group 2: Element Access
	// ============================================================================

	t.Run("Get returns element at valid index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.Get(1)
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, 20, result.Unwrap())
	})

	t.Run("Get returns None for negative index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.Get(-1)
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("Get returns None for index out of bounds", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.Get(5)
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("First returns first element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.First()
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, 10, result.Unwrap())
	})

	t.Run("First returns None on empty array", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(0).
			Build()
		// ========= [A]ct     =========
		result := arr.First()
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("Last returns last element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.Last()
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, 30, result.Unwrap())
	})

	t.Run("Last returns None on empty array", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(0).
			Build()
		// ========= [A]ct     =========
		result := arr.Last()
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	// ============================================================================
	// Group 3: Element Modification and Fixed-Size Behavior
	// ============================================================================

	t.Run("Offer sets element at valid index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		// ========= [A]ct     =========
		success := arr.Offer(42, 1)
		// ========= [A]ssert  =========
		must.True(t, success)
		must.Eq(t, []int{0, 42, 0}, arr.ToSlice())
	})

	t.Run("Offer fails for negative index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		// ========= [A]ct     =========
		success := arr.Offer(42, -1)
		// ========= [A]ssert  =========
		must.False(t, success)
	})

	t.Run("Offer fails for index out of bounds - array is fixed size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		// ========= [A]ct     =========
		success := arr.Offer(42, 5)
		// ========= [A]ssert  =========
		must.False(t, success)
		must.Eq(t, 3, arr.Len()) // Size remains fixed
	})

	t.Run("OfferRange sets elements from start by default", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		// ========= [A]ct     =========
		success := arr.OfferRange([]int{1, 2, 3})
		// ========= [A]ssert  =========
		must.True(t, success)
		must.Eq(t, []int{1, 2, 3, 0, 0}, arr.ToSlice())
		must.Eq(t, 5, arr.Len()) // Size remains fixed
	})

	t.Run("OfferRange fails when elements exceed array size - array is fixed size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		// ========= [A]ct     =========
		success := arr.OfferRange([]int{1, 2, 3, 4, 5})
		// ========= [A]ssert  =========
		must.False(t, success)
		must.Eq(t, 3, arr.Len()) // Size remains fixed
	})

	t.Run("OfferRange fails when negative from index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		// ========= [A]ct     =========
		success := arr.OfferRange([]int{1, 2, 3}, core_range.WithRangeFrom(-1))
		// ========= [A]ssert  =========
		must.False(t, success)
	})

	t.Run("OfferRange fails when end index exceeds array size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		// ========= [A]ct     =========
		success := arr.OfferRange([]int{1, 2, 3}, core_range.WithRangeEnd(10))
		// ========= [A]ssert  =========
		must.False(t, success)
		must.Eq(t, 5, arr.Len()) // Size remains fixed
	})

	t.Run("Len remains constant after all operations - array is fixed size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		initialLen := arr.Len()

		// ========= [A]ct     =========
		arr.OfferRange([]int{1, 2, 3})
		arr.Offer(99, 4)
		arr.Clear()

		// ========= [A]ssert  =========
		must.Eq(t, initialLen, arr.Len())
		must.Eq(t, 5, arr.Len())
	})

	// ============================================================================
	// Group 4: Iteration
	// ============================================================================

	t.Run("Values iterates over all elements", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{10, 20, 30})
		// ========= [A]ct     =========
		var collected []int
		for item := range arr.Values() {
			collected = append(collected, item)
		}
		// ========= [A]ssert  =========
		must.Eq(t, []int{10, 20, 30}, collected)
	})

	t.Run("Values can be stopped early", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		arr.OfferRange([]int{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		var collected []int
		for item := range arr.Values() {
			collected = append(collected, item)
			if item == 3 {
				break
			}
		}
		// ========= [A]ssert  =========
		must.Eq(t, []int{1, 2, 3}, collected)
	})

	t.Run("All iterates with index and value", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{10, 20, 30})
		// ========= [A]ct     =========
		var indices []int
		var values []int
		for idx, val := range arr.All() {
			indices = append(indices, idx)
			values = append(values, val)
		}
		// ========= [A]ssert  =========
		must.Eq(t, []int{0, 1, 2}, indices)
		must.Eq(t, []int{10, 20, 30}, values)
	})

	t.Run("All can be stopped early", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		arr.OfferRange([]int{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		var indices []int
		for idx := range arr.All() {
			indices = append(indices, idx)
			if idx == 2 {
				break
			}
		}
		// ========= [A]ssert  =========
		must.Eq(t, []int{0, 1, 2}, indices)
	})

	t.Run("ForEach executes function for each element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		arr.OfferRange([]int{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		sum := 0
		arr.ForEach(func(item int) {
			sum += item
		})
		// ========= [A]ssert  =========
		must.Eq(t, 15, sum)
	})

	// ============================================================================
	// Group 5: Search Methods
	// ============================================================================

	t.Run("Find returns first matching element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		arr.OfferRange([]int{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		result := arr.Find(func(item int) bool {
			return item > 2
		})
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, 3, result.Unwrap())
	})

	t.Run("Find returns None when no match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 2, 3})
		// ========= [A]ct     =========
		result := arr.Find(func(item int) bool {
			return item > 10
		})
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("FindIndex returns index of first matching element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		arr.OfferRange([]int{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		result := arr.FindIndex(func(item int) bool {
			return item > 2
		})
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, 2, result.Unwrap())
	})

	t.Run("FindIndex returns None when no match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 2, 3})
		// ========= [A]ct     =========
		result := arr.FindIndex(func(item int) bool {
			return item > 10
		})
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("Any returns true when predicate matches", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		arr.OfferRange([]int{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		result := arr.Any(func(item int) bool {
			return item == 3
		})
		// ========= [A]ssert  =========
		must.True(t, result)
	})

	t.Run("Any returns false when no predicate matches", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 2, 3})
		// ========= [A]ct     =========
		result := arr.Any(func(item int) bool {
			return item > 10
		})
		// ========= [A]ssert  =========
		must.False(t, result)
	})

	t.Run("Every returns true when all elements match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		arr.OfferRange([]int{2, 4, 6, 8, 10})
		// ========= [A]ct     =========
		result := arr.Every(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.True(t, result)
	})

	t.Run("Every returns false when not all elements match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		arr.OfferRange([]int{2, 4, 5, 8, 10})
		// ========= [A]ct     =========
		result := arr.Every(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.False(t, result)
	})

	t.Run("Count returns number of matching elements", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(6).
			Build()
		arr.OfferRange([]int{1, 2, 3, 4, 5, 6})
		// ========= [A]ct     =========
		result := arr.Count(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, 3, result)
	})

	t.Run("Count returns zero when no elements match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 3, 5})
		// ========= [A]ct     =========
		result := arr.Count(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, 0, result)
	})

	// ============================================================================
	// Group 6: Modification Methods (Retain, Filter)
	// ============================================================================

	t.Run("Retain keeps only matching elements and zeros others", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(6).
			Build()
		arr.OfferRange([]int{1, 2, 3, 4, 5, 6})
		// ========= [A]ct     =========
		arr.Retain(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, []int{0, 2, 0, 4, 0, 6}, arr.ToSlice())
		must.Eq(t, 6, arr.Len()) // Size remains fixed
	})

	t.Run("Retain zeros all elements when none match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 3, 5})
		// ========= [A]ct     =========
		arr.Retain(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, []int{0, 0, 0}, arr.ToSlice())
		must.True(t, arr.IsEmpty())
	})

	t.Run("Filter returns new array with matching elements", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(6).
			Build()
		arr.OfferRange([]int{1, 2, 3, 4, 5, 6})
		// ========= [A]ct     =========
		filtered := arr.Filter(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, []int{2, 4, 6}, filtered.ToSlice())
		must.Eq(t, 3, filtered.Len())
		// Original array unchanged
		must.Eq(t, []int{1, 2, 3, 4, 5, 6}, arr.ToSlice())
		must.Eq(t, 6, arr.Len())
	})

	t.Run("Filter returns empty array when no match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(3).
			Build()
		arr.OfferRange([]int{1, 3, 5})
		// ========= [A]ct     =========
		filtered := arr.Filter(func(item int) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, 0, filtered.Len())
		must.True(t, filtered.IsEmpty())
	})

	// ============================================================================
	// Group 7: Sorting
	// ============================================================================

	t.Run("Sort orders elements ascending", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(6).
			Build()
		arr.OfferRange([]int{5, 2, 8, 1, 9, 3})
		// ========= [A]ct     =========
		arr.Sort(func(a, b int) compare.Order {
			if a < b {
				return compare.OrderLess
			} else if a > b {
				return compare.OrderGreater
			}
			return compare.OrderEqual
		})
		// ========= [A]ssert  =========
		must.Eq(t, []int{1, 2, 3, 5, 8, 9}, arr.ToSlice())
		must.Eq(t, 6, arr.Len()) // Size remains fixed
	})

	t.Run("Sort orders elements descending", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(6).
			Build()
		arr.OfferRange([]int{5, 2, 8, 1, 9, 3})
		// ========= [A]ct     =========
		arr.Sort(func(a, b int) compare.Order {
			if a > b {
				return compare.OrderLess
			} else if a < b {
				return compare.OrderGreater
			}
			return compare.OrderEqual
		})
		// ========= [A]ssert  =========
		must.Eq(t, []int{9, 8, 5, 3, 2, 1}, arr.ToSlice())
	})

	t.Run("Sort is stable", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewBuiltinBuilder[int]().
			Size(5).
			Build()
		arr.OfferRange([]int{3, 1, 2, 1, 1})
		// ========= [A]ct     =========
		arr.Sort(func(a, b int) compare.Order {
			if a < b {
				return compare.OrderLess
			} else if a > b {
				return compare.OrderGreater
			}
			return compare.OrderEqual
		})
		// ========= [A]ssert  =========
		must.Eq(t, []int{1, 1, 1, 2, 3}, arr.ToSlice())
	})
}

type ComparableInt int

func (c ComparableInt) Equal(other ComparableInt) bool {
	return c == other
}

func (c ComparableInt) Compare(other ComparableInt) compare.Order {
	if c < other {
		return compare.OrderLess
	} else if c > other {
		return compare.OrderGreater
	}
	return compare.OrderEqual
}

func TestArrayFromComparable(t *testing.T) {
	// ============================================================================
	// Group 1: Building and Basic Properties
	// ============================================================================

	t.Run("Can build with size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 5, arr.Len())
	})

	t.Run("Len returns fixed size regardless of content", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(10).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 10, arr.Len())
	})

	t.Run("IsEmpty on new array returns true", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		// ========= [A]ssert  =========
		must.True(t, arr.IsEmpty())
	})

	t.Run("IsEmpty returns false when array has non-zero values", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3})
		// ========= [A]ssert  =========
		must.False(t, arr.IsEmpty())
	})

	t.Run("Clear resets all elements to zero values", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3})
		// ========= [A]ct     =========
		arr.Clear()
		// ========= [A]ssert  =========
		must.Eq(t, 3, arr.Len())
		must.True(t, arr.IsEmpty())
	})

	t.Run("Contains finds existing element using Equal", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3})
		// ========= [A]ssert  =========
		must.True(t, arr.Contains(2))
	})

	t.Run("Contains returns false for missing element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3})
		// ========= [A]ssert  =========
		must.False(t, arr.Contains(4))
	})

	// ============================================================================
	// Group 2: Element Access
	// ============================================================================

	t.Run("Get returns element at valid index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.Get(1)
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, ComparableInt(20), result.Unwrap())
	})

	t.Run("Get returns None for negative index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.Get(-1)
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("Get returns None for index out of bounds", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.Get(5)
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("First returns first element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.First()
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, ComparableInt(10), result.Unwrap())
	})

	t.Run("First returns None on empty array", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(0).
			Build()
		// ========= [A]ct     =========
		result := arr.First()
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("Last returns last element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{10, 20, 30})
		// ========= [A]ct     =========
		result := arr.Last()
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, ComparableInt(30), result.Unwrap())
	})

	t.Run("Last returns None on empty array", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(0).
			Build()
		// ========= [A]ct     =========
		result := arr.Last()
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	// ============================================================================
	// Group 3: Element Modification and Fixed-Size Behavior
	// ============================================================================

	t.Run("Offer sets element at valid index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		// ========= [A]ct     =========
		success := arr.Offer(42, 1)
		// ========= [A]ssert  =========
		must.True(t, success)
		must.Eq(t, []ComparableInt{0, 42, 0}, arr.ToSlice())
	})

	t.Run("Offer fails for negative index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		// ========= [A]ct     =========
		success := arr.Offer(42, -1)
		// ========= [A]ssert  =========
		must.False(t, success)
	})

	t.Run("Offer fails for index out of bounds - array is fixed size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		// ========= [A]ct     =========
		success := arr.Offer(42, 5)
		// ========= [A]ssert  =========
		must.False(t, success)
		must.Eq(t, 3, arr.Len()) // Size remains fixed
	})

	t.Run("OfferRange sets elements from start by default", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		// ========= [A]ct     =========
		success := arr.OfferRange([]ComparableInt{1, 2, 3})
		// ========= [A]ssert  =========
		must.True(t, success)
		must.Eq(t, []ComparableInt{1, 2, 3, 0, 0}, arr.ToSlice())
		must.Eq(t, 5, arr.Len()) // Size remains fixed
	})

	t.Run("OfferRange fails when elements exceed array size - array is fixed size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		// ========= [A]ct     =========
		success := arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5})
		// ========= [A]ssert  =========
		must.False(t, success)
		must.Eq(t, 3, arr.Len()) // Size remains fixed
	})

	t.Run("OfferRange fails when negative from index", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		// ========= [A]ct     =========
		success := arr.OfferRange([]ComparableInt{1, 2, 3}, core_range.WithRangeFrom(-1))
		// ========= [A]ssert  =========
		must.False(t, success)
	})

	t.Run("OfferRange fails when end index exceeds array size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		// ========= [A]ct     =========
		success := arr.OfferRange([]ComparableInt{1, 2, 3}, core_range.WithRangeEnd(10))
		// ========= [A]ssert  =========
		must.False(t, success)
		must.Eq(t, 5, arr.Len()) // Size remains fixed
	})

	t.Run("Len remains constant after all operations - array is fixed size", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		initialLen := arr.Len()

		// ========= [A]ct     =========
		arr.OfferRange([]ComparableInt{1, 2, 3})
		arr.Offer(99, 4)
		arr.Clear()

		// ========= [A]ssert  =========
		must.Eq(t, initialLen, arr.Len())
		must.Eq(t, 5, arr.Len())
	})

	// ============================================================================
	// Group 4: Iteration
	// ============================================================================

	t.Run("Values iterates over all elements", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{10, 20, 30})
		// ========= [A]ct     =========
		var collected []ComparableInt
		for item := range arr.Values() {
			collected = append(collected, item)
		}
		// ========= [A]ssert  =========
		must.Eq(t, []ComparableInt{10, 20, 30}, collected)
	})

	t.Run("Values can be stopped early", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		var collected []ComparableInt
		for item := range arr.Values() {
			collected = append(collected, item)
			if item == 3 {
				break
			}
		}
		// ========= [A]ssert  =========
		must.Eq(t, []ComparableInt{1, 2, 3}, collected)
	})

	t.Run("All iterates with index and value", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{10, 20, 30})
		// ========= [A]ct     =========
		var indices []int
		var values []ComparableInt
		for idx, val := range arr.All() {
			indices = append(indices, idx)
			values = append(values, val)
		}
		// ========= [A]ssert  =========
		must.Eq(t, []int{0, 1, 2}, indices)
		must.Eq(t, []ComparableInt{10, 20, 30}, values)
	})

	t.Run("All can be stopped early", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		var indices []int
		for idx := range arr.All() {
			indices = append(indices, idx)
			if idx == 2 {
				break
			}
		}
		// ========= [A]ssert  =========
		must.Eq(t, []int{0, 1, 2}, indices)
	})

	t.Run("ForEach executes function for each element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		var sum ComparableInt = 0
		arr.ForEach(func(item ComparableInt) {
			sum += item
		})
		// ========= [A]ssert  =========
		must.Eq(t, ComparableInt(15), sum)
	})

	// ============================================================================
	// Group 5: Search Methods
	// ============================================================================

	t.Run("Find returns first matching element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		result := arr.Find(func(item ComparableInt) bool {
			return item > 2
		})
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, ComparableInt(3), result.Unwrap())
	})

	t.Run("Find returns None when no match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3})
		// ========= [A]ct     =========
		result := arr.Find(func(item ComparableInt) bool {
			return item > 10
		})
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("FindIndex returns index of first matching element", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		result := arr.FindIndex(func(item ComparableInt) bool {
			return item > 2
		})
		// ========= [A]ssert  =========
		must.True(t, result.IsSome())
		must.Eq(t, 2, result.Unwrap())
	})

	t.Run("FindIndex returns None when no match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3})
		// ========= [A]ct     =========
		result := arr.FindIndex(func(item ComparableInt) bool {
			return item > 10
		})
		// ========= [A]ssert  =========
		must.True(t, result.IsNone())
	})

	t.Run("Any returns true when predicate matches", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5})
		// ========= [A]ct     =========
		result := arr.Any(func(item ComparableInt) bool {
			return item == 3
		})
		// ========= [A]ssert  =========
		must.True(t, result)
	})

	t.Run("Any returns false when no predicate matches", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3})
		// ========= [A]ct     =========
		result := arr.Any(func(item ComparableInt) bool {
			return item > 10
		})
		// ========= [A]ssert  =========
		must.False(t, result)
	})

	t.Run("Every returns true when all elements match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		arr.OfferRange([]ComparableInt{2, 4, 6, 8, 10})
		// ========= [A]ct     =========
		result := arr.Every(func(item ComparableInt) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.True(t, result)
	})

	t.Run("Every returns false when not all elements match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		arr.OfferRange([]ComparableInt{2, 4, 5, 8, 10})
		// ========= [A]ct     =========
		result := arr.Every(func(item ComparableInt) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.False(t, result)
	})

	t.Run("Count returns number of matching elements", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(6).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5, 6})
		// ========= [A]ct     =========
		result := arr.Count(func(item ComparableInt) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, 3, result)
	})

	t.Run("Count returns zero when no elements match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 3, 5})
		// ========= [A]ct     =========
		result := arr.Count(func(item ComparableInt) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, 0, result)
	})

	// ============================================================================
	// Group 6: Modification Methods (Retain, Filter)
	// ============================================================================

	t.Run("Retain keeps only matching elements and zeros others", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(6).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5, 6})
		// ========= [A]ct     =========
		arr.Retain(func(item ComparableInt) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, []ComparableInt{0, 2, 0, 4, 0, 6}, arr.ToSlice())
		must.Eq(t, 6, arr.Len()) // Size remains fixed
	})

	t.Run("Retain zeros all elements when none match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 3, 5})
		// ========= [A]ct     =========
		arr.Retain(func(item ComparableInt) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, []ComparableInt{0, 0, 0}, arr.ToSlice())
		must.True(t, arr.IsEmpty())
	})

	t.Run("Filter returns new array with matching elements", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(6).
			Build()
		arr.OfferRange([]ComparableInt{1, 2, 3, 4, 5, 6})
		// ========= [A]ct     =========
		filtered := arr.Filter(func(item ComparableInt) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, []ComparableInt{2, 4, 6}, filtered.ToSlice())
		must.Eq(t, 3, filtered.Len())
		// Original array unchanged
		must.Eq(t, []ComparableInt{1, 2, 3, 4, 5, 6}, arr.ToSlice())
		must.Eq(t, 6, arr.Len())
	})

	t.Run("Filter returns empty array when no match", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(3).
			Build()
		arr.OfferRange([]ComparableInt{1, 3, 5})
		// ========= [A]ct     =========
		filtered := arr.Filter(func(item ComparableInt) bool {
			return item%2 == 0
		})
		// ========= [A]ssert  =========
		must.Eq(t, 0, filtered.Len())
		must.True(t, filtered.IsEmpty())
	})

	// ============================================================================
	// Group 7: Sorting
	// ============================================================================

	t.Run("Sort orders elements ascending", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(6).
			Build()
		arr.OfferRange([]ComparableInt{5, 2, 8, 1, 9, 3})
		// ========= [A]ct     =========
		arr.Sort(func(a, b ComparableInt) compare.Order {
			return a.Compare(b)
		})
		// ========= [A]ssert  =========
		must.Eq(t, []ComparableInt{1, 2, 3, 5, 8, 9}, arr.ToSlice())
		must.Eq(t, 6, arr.Len()) // Size remains fixed
	})

	t.Run("Sort orders elements descending", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(6).
			Build()
		arr.OfferRange([]ComparableInt{5, 2, 8, 1, 9, 3})
		// ========= [A]ct     =========
		arr.Sort(func(a, b ComparableInt) compare.Order {
			return b.Compare(a) // Reversed for descending
		})
		// ========= [A]ssert  =========
		must.Eq(t, []ComparableInt{9, 8, 5, 3, 2, 1}, arr.ToSlice())
	})

	t.Run("Sort is stable", func(t *testing.T) {
		// ========= [A]rrange =========
		arr := array.NewComparableInterfaceBuilder[ComparableInt]().
			Size(5).
			Build()
		arr.OfferRange([]ComparableInt{3, 1, 2, 1, 1})
		// ========= [A]ct     =========
		arr.Sort(func(a, b ComparableInt) compare.Order {
			return a.Compare(b)
		})
		// ========= [A]ssert  =========
		must.Eq(t, []ComparableInt{1, 1, 1, 2, 3}, arr.ToSlice())
	})
}
