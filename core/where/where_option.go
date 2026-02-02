package where

import (
	. "codeberg.org/yaadata/opt"
)

// WhereOption configures a Where range specification.
type WhereOption[K any] func(ranger *Where[K])

// From sets the start point of the range.
func From[K any](from K) WhereOption[K] {
	return func(ranger *Where[K]) {
		ranger.from = Some(from)
	}
}

// To sets the end point of the range.
func To[K any](end K) WhereOption[K] {
	return func(w *Where[K]) {
		w.to = Some(end)
	}
}
