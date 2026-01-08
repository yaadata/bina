package hashset_test

import (
	"testing"

	"github.com/shoenig/test/must"

	hashset "codeberg.org/yaadata/bina/set/hashset"
)

type HashableInt int

func (h HashableInt) Hash() int {
	return int(h)
}

func TestHashSetFromBuiltin(t *testing.T) {
	t.Run("Can build", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewBuiltinBuilder[int]().
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 0, set.Len())
	})

	t.Run("Can build with capacity", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewBuiltinBuilder[int]().
			Capacity(10).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 0, set.Len())
	})

	t.Run("Can build from items", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 3, set.Len())
		must.True(t, set.Contains(1))
		must.True(t, set.Contains(2))
		must.True(t, set.Contains(3))
	})

	t.Run("Can build from items with duplicates", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewBuiltinBuilder[int]().
			From(1, 2, 2, 3, 3, 3).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 3, set.Len())
	})

	t.Run("Collection methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewBuiltinBuilder[int]().
			From(1, 2, 3).
			Build()

		// SCENARIO: Len
		t.Run("Len", func(t *testing.T) {
			// ========= [A]ct     =========
			length := set.Len()
			// ========= [A]ssert  =========
			must.Eq(t, 3, length)
		})

		// SCENARIO: Contains
		t.Run("Contains - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Contains(4)
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("Contains - true", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Contains(2)
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: IsEmpty
		t.Run("IsEmpty - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.IsEmpty()
			// ========= [A]ssert  =========
			must.False(t, actual)
		})

		// SCENARIO: Clear
		t.Run("Clear", func(t *testing.T) {
			// ========= [A]ct     =========
			set.Clear()
			// ========= [A]ssert  =========
			must.Eq(t, 0, set.Len())
			must.True(t, set.IsEmpty())
		})
	})

	t.Run("Aggregate methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewBuiltinBuilder[int]().
			From(1, 2, 3, 4, 5).
			Build()

		// SCENARIO: Any
		t.Run("Any - False", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Any(func(item int) bool {
				return item > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("Any - True", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Any(func(item int) bool {
				return item > 3
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: Count
		t.Run("Count", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Count(func(item int) bool {
				return item%2 == 0
			})
			// ========= [A]ssert  =========
			must.Eq(t, 2, actual)
		})

		// SCENARIO: Every
		t.Run("Every - True", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Every(func(item int) bool {
				return item < 10
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		t.Run("Every - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Every(func(item int) bool {
				return item > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})

		// SCENARIO: ForEach
		t.Run("ForEach", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 2, 3, 3, 3).
				Build()
			// ========= [A]ct     =========
			count := 0
			set.ForEach(func(item int) {
				count++
			})
			// ========= [A]ssert  =========
			must.Eq(t, 3, count)
		})
	})

	t.Run("Set mutation methods work", func(t *testing.T) {
		// SCENARIO: Add
		t.Run("Add - new element", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			added := set.Add(4)
			// ========= [A]ssert  =========
			must.True(t, added)
			must.Eq(t, 4, set.Len())
			must.True(t, set.Contains(4))
		})

		t.Run("Add - duplicate element", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			added := set.Add(2)
			// ========= [A]ssert  =========
			must.False(t, added)
			must.Eq(t, 3, set.Len())
		})

		// SCENARIO: All
		t.Run("All", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 2, 3, 3, 3).
				Build()
			// ========= [A]ct     =========
			count := 0
			for range set.Values() {
				count++
			}
			// ========= [A]ssert  =========
			must.Eq(t, 3, count)
		})

		// SCENARIO: Extend
		t.Run("Extend", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			set.Extend(4, 5, 6)
			// ========= [A]ssert  =========
			must.Eq(t, 6, set.Len())
			must.True(t, set.Contains(4))
			must.True(t, set.Contains(5))
			must.True(t, set.Contains(6))
		})

		t.Run("Extend - with duplicates", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			set.Extend(2, 3, 4)
			// ========= [A]ssert  =========
			must.Eq(t, 4, set.Len())
		})

		// SCENARIO: Remove
		t.Run("Remove - existing element", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			removed := set.Remove(2)
			// ========= [A]ssert  =========
			must.True(t, removed)
			must.Eq(t, 2, set.Len())
			must.False(t, set.Contains(2))
		})

		t.Run("Remove - non-existing element", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			removed := set.Remove(4)
			// ========= [A]ssert  =========
			must.False(t, removed)
			must.Eq(t, 3, set.Len())
		})
	})

	t.Run("Set operations work", func(t *testing.T) {
		// SCENARIO: Union
		t.Run("Union", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewBuiltinBuilder[int]().
				From(3, 4, 5).
				Build()
			// ========= [A]ct     =========
			union := set1.Union(set2)
			// ========= [A]ssert  =========
			must.Eq(t, 5, union.Len())
			must.True(t, union.Contains(1))
			must.True(t, union.Contains(2))
			must.True(t, union.Contains(3))
			must.True(t, union.Contains(4))
			must.True(t, union.Contains(5))
		})

		// SCENARIO: Intersect
		t.Run("Intersect - has common elements", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewBuiltinBuilder[int]().
				From(2, 3, 4).
				Build()
			// ========= [A]ct     =========
			intersect := set1.Intersect(set2)
			// ========= [A]ssert  =========
			must.True(t, intersect.IsSome())
			must.Eq(t, 2, intersect.Unwrap().Len())
			must.True(t, intersect.Unwrap().Contains(2))
			must.True(t, intersect.Unwrap().Contains(3))
		})

		t.Run("Intersect - no common elements", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewBuiltinBuilder[int]().
				From(4, 5, 6).
				Build()
			// ========= [A]ct     =========
			intersect := set1.Intersect(set2)
			// ========= [A]ssert  =========
			must.True(t, intersect.IsNone())
		})

		// SCENARIO: IsSubsetOf
		t.Run("IsSubsetOf - true", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewBuiltinBuilder[int]().
				From(1, 2).
				Build()
			set2 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3, 4).
				Build()
			// ========= [A]ct     =========
			result := set1.IsSubsetOf(set2)
			// ========= [A]ssert  =========
			must.True(t, result)
		})

		t.Run("IsSubsetOf - false", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 5).
				Build()
			set2 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3, 4).
				Build()
			// ========= [A]ct     =========
			result := set1.IsSubsetOf(set2)
			// ========= [A]ssert  =========
			must.False(t, result)
		})

		// SCENARIO: IsSupersetOf
		t.Run("IsSupersetOf - true", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3, 4).
				Build()
			set2 := hashset.NewBuiltinBuilder[int]().
				From(1, 2).
				Build()
			// ========= [A]ct     =========
			result := set1.IsSupersetOf(set2)
			// ========= [A]ssert  =========
			must.True(t, result)
		})

		t.Run("IsSupersetOf - false", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3, 4).
				Build()
			set2 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 5).
				Build()
			// ========= [A]ct     =========
			result := set1.IsSupersetOf(set2)
			// ========= [A]ssert  =========
			must.False(t, result)
		})

		// SCENARIO: SymmetricDifference
		t.Run("SymmetricDifference - has difference", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewBuiltinBuilder[int]().
				From(2, 3, 4).
				Build()
			// ========= [A]ct     =========
			symDiff := set1.SymmetricDifference(set2)
			// ========= [A]ssert  =========
			must.True(t, symDiff.IsSome())
			must.Eq(t, 2, symDiff.Unwrap().Len())
			must.True(t, symDiff.Unwrap().Contains(1))
			must.True(t, symDiff.Unwrap().Contains(4))
		})

		t.Run("SymmetricDifference - identical sets", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewBuiltinBuilder[int]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			symDiff := set1.SymmetricDifference(set2)
			// ========= [A]ssert  =========
			must.True(t, symDiff.IsNone())
		})
	})
}

func TestHashSetFromHashable(t *testing.T) {
	t.Run("Can build", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewHashableBuilder[int, HashableInt]().
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 0, set.Len())
	})

	t.Run("Can build with capacity", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewHashableBuilder[int, HashableInt]().
			Capacity(10).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 0, set.Len())
	})

	t.Run("Can build from items", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewHashableBuilder[int, HashableInt]().
			From(1, 2, 3).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 3, set.Len())
		must.True(t, set.Contains(1))
		must.True(t, set.Contains(2))
		must.True(t, set.Contains(3))
	})

	t.Run("Can build from items with duplicates", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewHashableBuilder[int, HashableInt]().
			From(1, 2, 2, 3, 3, 3).
			Build()
		// ========= [A]ssert  =========
		must.Eq(t, 3, set.Len())
	})

	t.Run("Collection methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewHashableBuilder[int, HashableInt]().
			From(1, 2, 3).
			Build()

		// SCENARIO: Len
		t.Run("Len", func(t *testing.T) {
			// ========= [A]ct     =========
			length := set.Len()
			// ========= [A]ssert  =========
			must.Eq(t, 3, length)
		})

		// SCENARIO: Contains
		t.Run("Contains - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Contains(4)
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("Contains - true", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Contains(2)
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: IsEmpty
		t.Run("IsEmpty - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.IsEmpty()
			// ========= [A]ssert  =========
			must.False(t, actual)
		})

		// SCENARIO: Clear
		t.Run("Clear", func(t *testing.T) {
			// ========= [A]ct     =========
			set.Clear()
			// ========= [A]ssert  =========
			must.Eq(t, 0, set.Len())
			must.True(t, set.IsEmpty())
		})
	})

	t.Run("Aggregate methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		set := hashset.NewHashableBuilder[int, HashableInt]().
			From(1, 2, 3, 4, 5).
			Build()

		// SCENARIO: Any
		t.Run("Any - False", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Any(func(item HashableInt) bool {
				return item > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("Any - True", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Any(func(item HashableInt) bool {
				return item > 3
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: Count
		t.Run("Count", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Count(func(item HashableInt) bool {
				return item%2 == 0
			})
			// ========= [A]ssert  =========
			must.Eq(t, 2, actual)
		})

		// SCENARIO: Every
		t.Run("Every - True", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Every(func(item HashableInt) bool {
				return item < 10
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		t.Run("Every - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := set.Every(func(item HashableInt) bool {
				return item > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})

		// SCENARIO: ForEach
		t.Run("ForEach", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 2, 3, 3, 3).
				Build()
			// ========= [A]ct     =========
			count := 0
			set.ForEach(func(item HashableInt) {
				count++
			})
			// ========= [A]ssert  =========
			must.Eq(t, 3, count)
		})
	})

	t.Run("Set mutation methods work", func(t *testing.T) {
		// SCENARIO: Add
		t.Run("Add - new element", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			added := set.Add(4)
			// ========= [A]ssert  =========
			must.True(t, added)
			must.Eq(t, 4, set.Len())
			must.True(t, set.Contains(4))
		})

		t.Run("Add - duplicate element", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			added := set.Add(2)
			// ========= [A]ssert  =========
			must.False(t, added)
			must.Eq(t, 3, set.Len())
		})

		// SCENARIO: All
		t.Run("All", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 2, 3, 3, 3).
				Build()
			// ========= [A]ct     =========
			count := 0
			for range set.Values() {
				count++
			}
			// ========= [A]ssert  =========
			must.Eq(t, 3, count)
		})

		// SCENARIO: Extend
		t.Run("Extend", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			set.Extend(4, 5, 6)
			// ========= [A]ssert  =========
			must.Eq(t, 6, set.Len())
			must.True(t, set.Contains(4))
			must.True(t, set.Contains(5))
			must.True(t, set.Contains(6))
		})

		t.Run("Extend - with duplicates", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			set.Extend(2, 3, 4)
			// ========= [A]ssert  =========
			must.Eq(t, 4, set.Len())
		})

		// SCENARIO: Remove
		t.Run("Remove - existing element", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			removed := set.Remove(2)
			// ========= [A]ssert  =========
			must.True(t, removed)
			must.Eq(t, 2, set.Len())
			must.False(t, set.Contains(2))
		})

		t.Run("Remove - non-existing element", func(t *testing.T) {
			// ========= [A]rrange =========
			set := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			removed := set.Remove(4)
			// ========= [A]ssert  =========
			must.False(t, removed)
			must.Eq(t, 3, set.Len())
		})
	})

	t.Run("Set operations work", func(t *testing.T) {
		// SCENARIO: Union
		t.Run("Union", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewHashableBuilder[int, HashableInt]().
				From(3, 4, 5).
				Build()
			// ========= [A]ct     =========
			union := set1.Union(set2)
			// ========= [A]ssert  =========
			must.Eq(t, 5, union.Len())
			must.True(t, union.Contains(1))
			must.True(t, union.Contains(2))
			must.True(t, union.Contains(3))
			must.True(t, union.Contains(4))
			must.True(t, union.Contains(5))
		})

		// SCENARIO: Intersect
		t.Run("Intersect - has common elements", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewHashableBuilder[int, HashableInt]().
				From(2, 3, 4).
				Build()
			// ========= [A]ct     =========
			intersect := set1.Intersect(set2)
			// ========= [A]ssert  =========
			must.True(t, intersect.IsSome())
			must.Eq(t, 2, intersect.Unwrap().Len())
			must.True(t, intersect.Unwrap().Contains(2))
			must.True(t, intersect.Unwrap().Contains(3))
		})

		t.Run("Intersect - no common elements", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewHashableBuilder[int, HashableInt]().
				From(4, 5, 6).
				Build()
			// ========= [A]ct     =========
			intersect := set1.Intersect(set2)
			// ========= [A]ssert  =========
			must.True(t, intersect.IsNone())
		})

		// SCENARIO: IsSubsetOf
		t.Run("IsSubsetOf - true", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2).
				Build()
			set2 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3, 4).
				Build()
			// ========= [A]ct     =========
			result := set1.IsSubsetOf(set2)
			// ========= [A]ssert  =========
			must.True(t, result)
		})

		t.Run("IsSubsetOf - false", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 5).
				Build()
			set2 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3, 4).
				Build()
			// ========= [A]ct     =========
			result := set1.IsSubsetOf(set2)
			// ========= [A]ssert  =========
			must.False(t, result)
		})

		// SCENARIO: IsSupersetOf
		t.Run("IsSupersetOf - true", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3, 4).
				Build()
			set2 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2).
				Build()
			// ========= [A]ct     =========
			result := set1.IsSupersetOf(set2)
			// ========= [A]ssert  =========
			must.True(t, result)
		})

		t.Run("IsSupersetOf - false", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3, 4).
				Build()
			set2 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 5).
				Build()
			// ========= [A]ct     =========
			result := set1.IsSupersetOf(set2)
			// ========= [A]ssert  =========
			must.False(t, result)
		})

		// SCENARIO: SymmetricDifference
		t.Run("SymmetricDifference - has difference", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewHashableBuilder[int, HashableInt]().
				From(2, 3, 4).
				Build()
			// ========= [A]ct     =========
			symDiff := set1.SymmetricDifference(set2)
			// ========= [A]ssert  =========
			must.True(t, symDiff.IsSome())
			must.Eq(t, 2, symDiff.Unwrap().Len())
			must.True(t, symDiff.Unwrap().Contains(1))
			must.True(t, symDiff.Unwrap().Contains(4))
		})

		t.Run("SymmetricDifference - identical sets", func(t *testing.T) {
			// ========= [A]rrange =========
			set1 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			set2 := hashset.NewHashableBuilder[int, HashableInt]().
				From(1, 2, 3).
				Build()
			// ========= [A]ct     =========
			symDiff := set1.SymmetricDifference(set2)
			// ========= [A]ssert  =========
			must.True(t, symDiff.IsNone())
		})
	})
}
