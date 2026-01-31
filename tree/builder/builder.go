package builder

import (
	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/kv"
)

type BaseBuilder[K any, V any, Target collection.SearchTree[K, V], Self any] interface {
	Build() Target
	// Order for SearchTree
	// Default order is 5
	Order(order int) Self
	From(pairs ...kv.Pair[K, V]) Self
}
