package shared

type Predicate[T any] func(item T) bool
