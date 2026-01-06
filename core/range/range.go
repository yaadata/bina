package core_range

import (
	. "github.com/yaadata/optionsgo"
)

type CoreRange struct {
	from Option[int]
	end  Option[int]
}

func New() *CoreRange {
	return &CoreRange{
		from: None[int](),
		end:  None[int](),
	}
}

func (c *CoreRange) From() Option[int] {
	return c.from
}

func (c *CoreRange) End() Option[int] {
	return c.end
}

type CoreRangeConfig func(ranger *CoreRange)

func WithRangeFrom(from int) CoreRangeConfig {
	return func(ranger *CoreRange) {
		ranger.from = Some(from)
	}
}

func WithRangeEnd(end int) CoreRangeConfig {
	return func(ranger *CoreRange) {
		ranger.end = Some(end)
	}
}
