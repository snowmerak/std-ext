package generator

import "github.com/snowmerak/std-ext/types/iterable"

var _ iterable.Iterable[int] = &Generator[int]{}

type Generator[T any] struct {
	initValue T
	prev      T
	gen       func(T) (T, bool)
}

func New[T any](initValue T, gen func(T) (T, bool)) *Generator[T] {
	return &Generator[T]{
		initValue: initValue,
		prev:      initValue,
		gen:       gen,
	}
}

func (g *Generator[T]) Reset() {
	g.prev = g.initValue
}

func (g *Generator[T]) Next() bool {
	v, ok := g.gen(g.prev)
	if !ok {
		return false
	}

	g.prev = v
	return true
}

func (g *Generator[T]) Value() T {
	return g.prev
}
