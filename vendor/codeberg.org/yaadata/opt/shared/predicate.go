package shared

type Predicate[T any] func(val T) bool
