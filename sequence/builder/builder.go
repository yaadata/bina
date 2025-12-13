package builder

import (
	"github.com/yaadata/bina/core/sequence"
)

type Builder[T any] interface {
	From(items ...T) Builder[T]
	Capacity(cap int) Builder[T]
	Build() sequence.Sequence[T]
}
