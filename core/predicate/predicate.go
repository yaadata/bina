package predicate

type Predicate[T any] func(item T) bool
