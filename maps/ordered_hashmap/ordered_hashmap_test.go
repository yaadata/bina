package orderedhashmap_test

import (
	"slices"
	"testing"

	"github.com/shoenig/test/must"

	"codeberg.org/yaadata/bina/maps"
	orderedhashmap "codeberg.org/yaadata/bina/maps/ordered_hashmap"
)

func TestOrderedHashMapBuiltinBuilder(t *testing.T) {
	t.Run("Builder methods work", func(t *testing.T) {
		t.Run("Can build", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().
				Build()
			// ========= [A]ssert  =========
			must.Eq(t, 0, m.Len())
		})

		t.Run("Can build with capacity", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().
				Capacity(10).
				Build()
			// ========= [A]ssert  =========
			must.Eq(t, 0, m.Len())
		})

		t.Run("Can build from builtin map", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
			// ========= [A]ssert  =========
			must.Eq(t, 3, m.Len())
			must.True(t, m.Contains("a"))
			must.True(t, m.Contains("b"))
			must.True(t, m.Contains("c"))
		})

		t.Run("Can build from builtin map with capacity", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().
				Capacity(10).
				From(map[string]int{"a": 1, "b": 2}).
				Build()
			// ========= [A]ssert  =========
			must.Eq(t, 2, m.Len())
		})
	})

	t.Run("Collection methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		m := orderedhashmap.BuiltinBuilder[string, int]().Build()
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)

		// SCENARIO: Len
		t.Run("Len", func(t *testing.T) {
			// ========= [A]ct     =========
			length := m.Len()
			// ========= [A]ssert  =========
			must.Eq(t, 3, length)
		})

		// SCENARIO: Contains
		t.Run("Contains - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Contains("d")
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("Contains - true", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Contains("b")
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: IsEmpty
		t.Run("IsEmpty - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.IsEmpty()
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("IsEmpty - true", func(t *testing.T) {
			// ========= [A]rrange =========
			emptyMap := orderedhashmap.BuiltinBuilder[string, int]().Build()
			// ========= [A]ct     =========
			actual := emptyMap.IsEmpty()
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: Clear
		t.Run("Clear", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			// ========= [A]ct     =========
			m.Clear()
			// ========= [A]ssert  =========
			must.Eq(t, 0, m.Len())
			must.True(t, m.IsEmpty())
		})
	})

	t.Run("Aggregate methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		m := orderedhashmap.BuiltinBuilder[string, int]().Build()
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)
		m.Put("d", 4)
		m.Put("e", 5)

		// SCENARIO: Any
		t.Run("Any - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Any(func(entry maps.MapEntry[string, int]) bool {
				return entry.Value() > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("Any - true", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Any(func(entry maps.MapEntry[string, int]) bool {
				return entry.Value() > 3
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: Count
		t.Run("Count", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Count(func(entry maps.MapEntry[string, int]) bool {
				return entry.Value()%2 == 0
			})
			// ========= [A]ssert  =========
			must.Eq(t, 2, actual)
		})

		// SCENARIO: Every
		t.Run("Every - true", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Every(func(entry maps.MapEntry[string, int]) bool {
				return entry.Value() < 10
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})
		t.Run("Every - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Every(func(entry maps.MapEntry[string, int]) bool {
				return entry.Value() > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})

		// SCENARIO: ForEach
		t.Run("ForEach", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			count := 0
			sum := 0
			m.ForEach(func(entry maps.MapEntry[string, int]) {
				count++
				sum += entry.Value()
			})
			// ========= [A]ssert  =========
			must.Eq(t, 3, count)
			must.Eq(t, 6, sum)
		})
	})

	t.Run("Map operations work", func(t *testing.T) {
		// SCENARIO: Get
		t.Run("Get - existing key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			result := m.Get("b")
			// ========= [A]ssert  =========
			must.True(t, result.IsSome())
			must.Eq(t, 2, result.Unwrap())
		})
		t.Run("Get - non-existing key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			result := m.Get("d")
			// ========= [A]ssert  =========
			must.True(t, result.IsNone())
		})

		// SCENARIO: Put
		t.Run("Put - new key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			// ========= [A]ct     =========
			result := m.Put("c", 3)
			// ========= [A]ssert  =========
			must.True(t, result)
			must.Eq(t, 3, m.Len())
			must.Eq(t, 3, m.Get("c").Unwrap())
		})
		t.Run("Put - overwrite existing key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			// ========= [A]ct     =========
			result := m.Put("b", 20)
			// ========= [A]ssert  =========
			must.False(t, result)
			must.Eq(t, 2, m.Len())
			must.Eq(t, 20, m.Get("b").Unwrap())
		})

		// SCENARIO: Delete
		t.Run("Delete - existing key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			result := m.Delete("b")
			// ========= [A]ssert  =========
			must.True(t, result.IsSome())
			must.Eq(t, 2, result.Unwrap())
			must.Eq(t, 2, m.Len())
			must.False(t, m.Contains("b"))
		})
		t.Run("Delete - non-existing key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			result := m.Delete("d")
			// ========= [A]ssert  =========
			must.True(t, result.IsNone())
			must.Eq(t, 3, m.Len())
		})

		// SCENARIO: Merge
		t.Run("Merge - no overlapping keys", func(t *testing.T) {
			// ========= [A]rrange =========
			m1 := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m1.Put("a", 1)
			m1.Put("b", 2)
			m2 := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m2.Put("c", 3)
			m2.Put("d", 4)
			// ========= [A]ct     =========
			merged := m1.Merge(m2, func(key string, current, incoming int) int {
				return current + incoming
			})
			// ========= [A]ssert  =========
			must.Eq(t, 4, merged.Len())
			must.Eq(t, 1, merged.Get("a").Unwrap())
			must.Eq(t, 2, merged.Get("b").Unwrap())
			must.Eq(t, 3, merged.Get("c").Unwrap())
			must.Eq(t, 4, merged.Get("d").Unwrap())
		})
		t.Run("Merge - with overlapping keys", func(t *testing.T) {
			// ========= [A]rrange =========
			m1 := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m1.Put("a", 1)
			m1.Put("b", 2)
			m1.Put("c", 3)
			m2 := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m2.Put("b", 20)
			m2.Put("c", 30)
			m2.Put("d", 4)
			// ========= [A]ct     =========
			merged := m1.Merge(m2, func(key string, current, incoming int) int {
				return current + incoming
			})
			// ========= [A]ssert  =========
			must.Eq(t, 4, merged.Len())
			must.Eq(t, 1, merged.Get("a").Unwrap())
			must.Eq(t, 22, merged.Get("b").Unwrap())
			must.Eq(t, 33, merged.Get("c").Unwrap())
			must.Eq(t, 4, merged.Get("d").Unwrap())
		})
	})

	t.Run("Iterator methods work", func(t *testing.T) {
		// SCENARIO: All
		t.Run("All", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			count := 0
			sum := 0
			for _, v := range m.All() {
				count++
				sum += v
			}
			// ========= [A]ssert  =========
			must.Eq(t, 3, count)
			must.Eq(t, 6, sum)
		})

		// SCENARIO: Keys
		t.Run("Keys", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			keys := make([]string, 0, 3)
			for k := range m.Keys() {
				keys = append(keys, k)
			}
			// ========= [A]ssert  =========
			must.Eq(t, 3, len(keys))
			must.SliceContains(t, keys, "a")
			must.SliceContains(t, keys, "b")
			must.SliceContains(t, keys, "c")
		})

		// SCENARIO: Values
		t.Run("Values", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			values := make([]int, 0, 3)
			for v := range m.Values() {
				values = append(values, v)
			}
			// ========= [A]ssert  =========
			must.Eq(t, 3, len(values))
			must.SliceContains(t, values, 1)
			must.SliceContains(t, values, 2)
			must.SliceContains(t, values, 3)
		})

		// SCENARIO: Empty map iterators
		t.Run("All - empty map", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			// ========= [A]ct     =========
			count := 0
			for range m.All() {
				count++
			}
			// ========= [A]ssert  =========
			must.Eq(t, 0, count)
		})

		t.Run("Keys - empty map", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			// ========= [A]ct     =========
			count := 0
			for range m.Keys() {
				count++
			}
			// ========= [A]ssert  =========
			must.Eq(t, 0, count)
		})

		t.Run("Values - empty map", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			// ========= [A]ct     =========
			count := 0
			for range m.Values() {
				count++
			}
			// ========= [A]ssert  =========
			must.Eq(t, 0, count)
		})
	})

	t.Run("OrderedMap methods work", func(t *testing.T) {
		// SCENARIO: First
		t.Run("First - non-empty map", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			result := m.First()
			// ========= [A]ssert  =========
			must.True(t, result.IsSome())
			must.Eq(t, "a", result.Unwrap().Key())
			must.Eq(t, 1, result.Unwrap().Value())
		})
		t.Run("First - empty map", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			// ========= [A]ct     =========
			result := m.First()
			// ========= [A]ssert  =========
			must.True(t, result.IsNone())
		})

		// SCENARIO: Last
		t.Run("Last - non-empty map", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			result := m.Last()
			// ========= [A]ssert  =========
			must.True(t, result.IsSome())
			must.Eq(t, "c", result.Unwrap().Key())
			must.Eq(t, 3, result.Unwrap().Value())
		})
		t.Run("Last - empty map", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			// ========= [A]ct     =========
			result := m.Last()
			// ========= [A]ssert  =========
			must.True(t, result.IsNone())
		})
	})

	t.Run("Ordering is preserved", func(t *testing.T) {
		t.Run("Insertion order is preserved in iteration", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("c", 3)
			m.Put("a", 1)
			m.Put("b", 2)
			// ========= [A]ct     =========
			keys := make([]string, 0, 3)
			for k := range m.Keys() {
				keys = append(keys, k)
			}
			// ========= [A]ssert  =========
			must.Eq(t, []string{"c", "a", "b"}, keys)
		})

		t.Run("Insertion order is preserved in All iterator", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("c", 3)
			m.Put("a", 1)
			m.Put("b", 2)
			// ========= [A]ct     =========
			keys := make([]string, 0, 3)
			values := make([]int, 0, 3)
			for k, v := range m.All() {
				keys = append(keys, k)
				values = append(values, v)
			}
			// ========= [A]ssert  =========
			must.Eq(t, []string{"c", "a", "b"}, keys)
			must.Eq(t, []int{3, 1, 2}, values)
		})

		t.Run("Insertion order is preserved in Values iterator", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("c", 3)
			m.Put("a", 1)
			m.Put("b", 2)
			// ========= [A]ct     =========
			values := make([]int, 0, 3)
			for v := range m.Values() {
				values = append(values, v)
			}
			// ========= [A]ssert  =========
			must.Eq(t, []int{3, 1, 2}, values)
		})

		t.Run("Updating existing key preserves its position", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			m.Put("b", 20)
			keys := make([]string, 0, 3)
			for k := range m.Keys() {
				keys = append(keys, k)
			}
			// ========= [A]ssert  =========
			must.Eq(t, []string{"a", "b", "c"}, keys)
			must.Eq(t, 20, m.Get("b").Unwrap())
		})

		t.Run("Order is preserved after delete", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			m.Put("d", 4)
			// ========= [A]ct     =========
			m.Delete("b")
			keys := make([]string, 0, 3)
			for k := range m.Keys() {
				keys = append(keys, k)
			}
			// ========= [A]ssert  =========
			must.Eq(t, []string{"a", "c", "d"}, keys)
		})

		t.Run("First and Last are correct after delete", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			m.Delete("a")
			// ========= [A]ssert  =========
			must.Eq(t, "b", m.First().Unwrap().Key())
			must.Eq(t, "c", m.Last().Unwrap().Key())
		})

		t.Run("First and Last are correct after deleting last element", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			m.Delete("c")
			// ========= [A]ssert  =========
			must.Eq(t, "a", m.First().Unwrap().Key())
			must.Eq(t, "b", m.Last().Unwrap().Key())
		})
	})

	t.Run("Compaction works correctly", func(t *testing.T) {
		t.Run("Many deletes followed by iteration", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			for i := range 100 {
				m.Put(string(rune('a'+i)), i)
			}
			// ========= [A]ct     =========
			// Delete every other element
			for i := 0; i < 100; i += 2 {
				m.Delete(string(rune('a' + i)))
			}
			// ========= [A]ssert  =========
			must.Eq(t, 50, m.Len())
			count := 0
			for range m.Keys() {
				count++
			}
			must.Eq(t, 50, count)
		})

		t.Run("Delete and re-add preserves new insertion order", func(t *testing.T) {
			// ========= [A]rrange =========
			m := orderedhashmap.BuiltinBuilder[string, int]().Build()
			m.Put("a", 1)
			m.Put("b", 2)
			m.Put("c", 3)
			// ========= [A]ct     =========
			m.Delete("b")
			m.Put("b", 20)
			keys := slices.Collect(m.Keys())
			// ========= [A]ssert  =========
			must.Eq(t, []string{"a", "c", "b"}, keys)
		})
	})
}
