package cloneable

type Cloneable[T any] interface {
	Clone() T
}
