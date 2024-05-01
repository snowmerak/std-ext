package iterable

type Iterable[T any] interface {
	Reset()
	Next() bool
	Value() T
}
