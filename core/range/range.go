package core_range

import (
	. "github.com/yaadata/optionsgo"
)

type Range struct {
	from Option[int]
	end  Option[int]
}

func New() *Range {
	return &Range{
		from: None[int](),
		end:  None[int](),
	}
}

func (c *Range) From() Option[int] {
	return c.from
}

func (c *Range) End() Option[int] {
	return c.end
}

type CoreRangeConfig func(ranger *Range)

func WithRangeFrom(from int) CoreRangeConfig {
	return func(ranger *Range) {
		ranger.from = Some(from)
	}
}

func WithRangeEnd(end int) CoreRangeConfig {
	return func(ranger *Range) {
		ranger.end = Some(end)
	}
}
