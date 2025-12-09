package slice_test

import (
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
}
