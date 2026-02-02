package collection

// DefaultSearchTreeTraversalConfiguration returns a configuration with in-order traversal.
func DefaultSearchTreeTraversalConfiguration() *SearchTreeTraversalConfiguration {
	return &SearchTreeTraversalConfiguration{
		strategy: SearchTreeStrategyInOrder,
	}
}

// SearchTreeTraversalConfiguration holds options for traversing a search tree.
type SearchTreeTraversalConfiguration struct {
	strategy SearchTreeStrategy
}

// Strategy returns the configured traversal strategy.
func (s *SearchTreeTraversalConfiguration) Strategy() SearchTreeStrategy {
	return s.strategy
}

// SearchTreeTraversalOption configures search tree traversal behavior.
type SearchTreeTraversalOption func(cfg *SearchTreeTraversalConfiguration)

// WithSearchTreeStrategy returns an option that sets the traversal strategy.
func WithSearchTreeStrategy(strategy SearchTreeStrategy) SearchTreeTraversalOption {
	return func(cfg *SearchTreeTraversalConfiguration) {
		cfg.strategy = strategy
	}
}
