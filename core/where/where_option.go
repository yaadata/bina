package where

import (
	. "codeberg.org/yaadata/opt"
)

type WhereOption[K any] func(ranger *Where[K])

func From[K any](from K) WhereOption[K] {
	return func(ranger *Where[K]) {
		ranger.from = Some(from)
	}
}

func To[K any](end K) WhereOption[K] {
	return func(w *Where[K]) {
		w.to = Some(end)
	}
}
