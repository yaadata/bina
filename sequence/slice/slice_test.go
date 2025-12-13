package slice_test

import (
	"slices"
	"testing"

	"github.com/shoenig/test/must"

	"github.com/yaadata/bina/sequence/slice"
)

func TestSliceFromBuiltin(t *testing.T) {
	t.Run("Can build", func(t *testing.T) {
		// [A]rrange - With capacity
		sequence := slice.NewBuiltinBuilder[int]().
			Capacity(10).
			Build()
		// [A]ssert
		must.Eq(t, 0, sequence.Len())
		must.Eq(t, 10, cap(sequence.ToSlice()))
	})

	t.Run("Can build without capacity", func(t *testing.T) {
		// [A]rrange
		sequence := slice.NewBuiltinBuilder[int]().
			Build()
		// [A]ssert
		must.Eq(t, 0, sequence.Len())
		must.Eq(t, 0, cap(sequence.ToSlice()))
	})

	t.Run("Can build from items", func(t *testing.T) {
		// [A]rrange
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		// [A]ssert
		must.Eq(t, 3, sequence.Len())
		must.Eq(t, []int{1, 2, 3}, sequence.ToSlice())
	})

	t.Run("Can enumerate", func(t *testing.T) {
		// [A]rrange
		original := []int{1, 2, 3}
		sequence := slice.NewBuiltinBuilder[int]().
			From(original...).
			Build()

		// [A]ssert
		for index, value := range sequence.Enumerate() {
			must.Eq(t, original[index], value)
		}
	})

	t.Run("Can gather all", func(t *testing.T) {
		// [A]rrange
		original := []int{1, 2, 3}
		sequence := slice.NewBuiltinBuilder[int]().
			From(original...).
			Build()
		// [A]ct
		var actual []int
		for value := range sequence.All() {
			actual = append(actual, value)
		}
		// [A]ssert
		must.True(t, slices.Equal(original, actual))
	})

	t.Run("Can Extend", func(t *testing.T) {
		// [A]rrange
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		// [A]ct
		sequence.Extend(4, 5, 6)
		// [A]ssert
		must.True(t, slices.Equal([]int{1, 2, 3, 4, 5, 6}, sequence.ToSlice()))
	})

	t.Run("Can Extend From Sequence", func(t *testing.T) {
		// [A]rrange
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		other := slice.NewBuiltinBuilder[int]().
			From(4, 5, 6).
			Build()
		// [A]ct
		sequence.ExtendFromSequence(other)
		// [A]ssert
		must.True(t, slices.Equal([]int{1, 2, 3, 4, 5, 6}, sequence.ToSlice()))
	})

	t.Run("Find returns Option", func(t *testing.T) {
		// ========= [A]rrange =========
		sequence := slice.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()

		// SCENARIO: no item matches predicate
		t.Run("No match", func(t *testing.T) {
			// ========= [A]ct =========
			actual := sequence.Find(func(item int) bool {
				return item == 4
			})
			// ========= [A]ssert =========
			must.False(t, actual.IsNone())
		})

		// SCENARIO: an item matches the predicate
		t.Run("Returns first match", func(t *testing.T) {
			// ========= [A]ct =========
			actual := sequence.Find(func(item int) bool {
				return item == 1 || item == 3
			})
			// ========= [A]ssert =========
			must.True(t, actual.IsSome())
			must.Eq(t, 1, actual.Unwrap())
		})
	})

}
