package maps

type MapEntry[K any, V any] interface {
	Key() K
	Value() V
}
