package hashmap_test

import (
	"testing"

	"github.com/shoenig/test/must"

	"codeberg.org/yaadata/bina/core/kv"
	"codeberg.org/yaadata/bina/maps/hashmap"
)

func TestHashMapBuiltinBuilder(t *testing.T) {
	t.Run("Builder methods work", func(t *testing.T) {
		t.Run("Can build", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
				Build()
			// ========= [A]ssert  =========
			must.Eq(t, 0, m.Len())
		})

		t.Run("Can build with capacity", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
				Capacity(10).
				Build()
			// ========= [A]ssert  =========
			must.Eq(t, 0, m.Len())
		})

		t.Run("Can build from builtin map", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
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
			m := hashmap.BuiltinBuilder[string, int]().
				Capacity(10).
				From(map[string]int{"a": 1, "b": 2}).
				Build()
			// ========= [A]ssert  =========
			must.Eq(t, 2, m.Len())
		})
	})

	t.Run("Collection methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		m := hashmap.BuiltinBuilder[string, int]().
			From(map[string]int{"a": 1, "b": 2, "c": 3}).
			Build()

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
			emptyMap := hashmap.BuiltinBuilder[string, int]().Build()
			// ========= [A]ct     =========
			actual := emptyMap.IsEmpty()
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: Clear
		t.Run("Clear", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2}).
				Build()
			// ========= [A]ct     =========
			m.Clear()
			// ========= [A]ssert  =========
			must.Eq(t, 0, m.Len())
			must.True(t, m.IsEmpty())
		})
	})

	t.Run("Aggregate methods work", func(t *testing.T) {
		// ========= [A]rrange =========
		m := hashmap.BuiltinBuilder[string, int]().
			From(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}).
			Build()

		// SCENARIO: Any
		t.Run("Any - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Any(func(entry kv.Pair[string, int]) bool {
				return entry.Value() > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})
		t.Run("Any - true", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Any(func(entry kv.Pair[string, int]) bool {
				return entry.Value() > 3
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})

		// SCENARIO: Count
		t.Run("Count", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Count(func(entry kv.Pair[string, int]) bool {
				return entry.Value()%2 == 0
			})
			// ========= [A]ssert  =========
			must.Eq(t, 2, actual)
		})

		// SCENARIO: Every
		t.Run("Every - true", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Every(func(entry kv.Pair[string, int]) bool {
				return entry.Value() < 10
			})
			// ========= [A]ssert  =========
			must.True(t, actual)
		})
		t.Run("Every - false", func(t *testing.T) {
			// ========= [A]ct     =========
			actual := m.Every(func(entry kv.Pair[string, int]) bool {
				return entry.Value() > 10
			})
			// ========= [A]ssert  =========
			must.False(t, actual)
		})

		// SCENARIO: ForEach
		t.Run("ForEach", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
			// ========= [A]ct     =========
			count := 0
			sum := 0
			m.ForEach(func(entry kv.Pair[string, int]) {
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
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
			// ========= [A]ct     =========
			result := m.Get("b")
			// ========= [A]ssert  =========
			must.True(t, result.IsSome())
			must.Eq(t, 2, result.Unwrap())
		})
		t.Run("Get - non-existing key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
			// ========= [A]ct     =========
			result := m.Get("d")
			// ========= [A]ssert  =========
			must.True(t, result.IsNone())
		})

		// SCENARIO: Put
		t.Run("Put - new key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2}).
				Build()
			// ========= [A]ct     =========
			result := m.Put("c", 3)
			// ========= [A]ssert  =========
			must.True(t, result)
			must.Eq(t, 3, m.Len())
			must.Eq(t, 3, m.Get("c").Unwrap())
		})
		t.Run("Put - overwrite existing key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2}).
				Build()
			// ========= [A]ct     =========
			result := m.Put("b", 20)
			// ========= [A]ssert  =========
			must.True(t, result)
			must.Eq(t, 2, m.Len())
			must.Eq(t, 20, m.Get("b").Unwrap())
		})

		// SCENARIO: Delete
		t.Run("Delete - existing key", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
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
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
			// ========= [A]ct     =========
			result := m.Delete("d")
			// ========= [A]ssert  =========
			must.True(t, result.IsNone())
			must.Eq(t, 3, m.Len())
		})

		// SCENARIO: Merge
		t.Run("Merge - no overlapping keys", func(t *testing.T) {
			// ========= [A]rrange =========
			m1 := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2}).
				Build()
			m2 := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"c": 3, "d": 4}).
				Build()
			// ========= [A]ct     =========
			merged := m1.Merge(m2, func(key string, current, incoming int) int {
				return current + incoming
			})
			// ========= [A]ssert  =========
			must.Eq(t, 2, merged.Len())
			must.Eq(t, 3, merged.Get("c").Unwrap())
			must.Eq(t, 4, merged.Get("d").Unwrap())
		})
		t.Run("Merge - with overlapping keys", func(t *testing.T) {
			// ========= [A]rrange =========
			m1 := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
			m2 := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"b": 20, "c": 30, "d": 4}).
				Build()
			// ========= [A]ct     =========
			merged := m1.Merge(m2, func(key string, current, incoming int) int {
				return current + incoming
			})
			// ========= [A]ssert  =========
			must.Eq(t, 3, merged.Len())
			must.Eq(t, 22, merged.Get("b").Unwrap())
			must.Eq(t, 33, merged.Get("c").Unwrap())
			must.Eq(t, 4, merged.Get("d").Unwrap())
		})
	})

	t.Run("Iterator methods work", func(t *testing.T) {
		// SCENARIO: All
		t.Run("All", func(t *testing.T) {
			// ========= [A]rrange =========
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
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
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
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
			m := hashmap.BuiltinBuilder[string, int]().
				From(map[string]int{"a": 1, "b": 2, "c": 3}).
				Build()
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
			m := hashmap.BuiltinBuilder[string, int]().Build()
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
			m := hashmap.BuiltinBuilder[string, int]().Build()
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
			m := hashmap.BuiltinBuilder[string, int]().Build()
			// ========= [A]ct     =========
			count := 0
			for range m.Values() {
				count++
			}
			// ========= [A]ssert  =========
			must.Eq(t, 0, count)
		})
	})
}
