package predicate

// Predicate is a function that tests whether an item satisfies a condition.
type Predicate[T any] func(item T) bool
