package collection

type MapMergeFunc[K comparable, V any] func(key K, current, incoming V) V
