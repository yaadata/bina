package collection

func DefaultSearchTreeTraversalConfiguration() *SearchTreeTraversalConfiguration {
	return &SearchTreeTraversalConfiguration{
		strategy: SearchTreeStrategyInOrder,
	}
}

type SearchTreeTraversalConfiguration struct {
	strategy SearchTreeStrategy
}

func (s *SearchTreeTraversalConfiguration) Strategy() SearchTreeStrategy {
	return s.strategy
}

type SearchTreeTraversalOption func(cfg *SearchTreeTraversalConfiguration)

func WithSearchTreeStrategy(strategy SearchTreeStrategy) SearchTreeTraversalOption {
	return func(cfg *SearchTreeTraversalConfiguration) {
		cfg.strategy = strategy
	}
}
