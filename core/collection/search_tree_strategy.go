package collection

type SearchTreeStrategy int

const (
	SearchTreeStrategyInOrder SearchTreeStrategy = iota
	SearchTreeStrategyPreOrder
	SearchTreeStrategyPostOrder
)
