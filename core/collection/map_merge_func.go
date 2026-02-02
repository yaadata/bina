package collection

// MapMergeFunc resolves conflicts when merging maps.
// It receives the key, the current value, and the incoming value, returning the merged result.
type MapMergeFunc[K comparable, V any] func(key K, current, incoming V) V
