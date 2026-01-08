package core_range

import (
	. "github.com/yaadata/optionsgo"
)

type Range[K any] struct {
	from Option[K]
	end  Option[K]
}

func New[K any]() *Range[K] {
	return &Range[K]{
		from: None[K](),
		end:  None[K](),
	}
}

func (c *Range[K]) From() Option[K] {
	return c.from
}

func (c *Range[K]) End() Option[K] {
	return c.end
}

type RangeConfig[K any] func(ranger *Range[K])

func WithRangeFrom[K any](from K) RangeConfig[K] {
	return func(ranger *Range[K]) {
		ranger.from = Some(from)
	}
}

func WithRangeEnd[K any](end K) RangeConfig[K] {
	return func(ranger *Range[K]) {
		ranger.end = Some(end)
	}
}
