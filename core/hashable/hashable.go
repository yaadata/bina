package hashable

type Hashable[T comparable] interface {
	Hash() T
}
