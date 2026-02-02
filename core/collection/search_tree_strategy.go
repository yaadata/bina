package collection

// SearchTreeStrategy specifies the traversal order for a search tree.
type SearchTreeStrategy int

const (
	// SearchTreeStrategyInOrder visits left subtree, node, then right subtree.
	SearchTreeStrategyInOrder SearchTreeStrategy = iota

	// SearchTreeStrategyPreOrder visits node, left subtree, then right subtree.
	SearchTreeStrategyPreOrder

	// SearchTreeStrategyPostOrder visits left subtree, right subtree, then node.
	SearchTreeStrategyPostOrder
)
