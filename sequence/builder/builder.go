package builder

import (
	"github.com/yaadata/bina/core/sequential"
)

type Builder[T any] interface {
	From(items ...T) Builder[T]
	Capacity(cap int) Builder[T]
	Build() sequential.Sequence[T]
}
