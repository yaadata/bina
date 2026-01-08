package tree

type BTree[K any, V any] interface {
	SearchTree[K, V]
	Order() int
}
