package stream

import (
	"github.com/snowmerak/std-ext/types/iterable"
	"sync"
)

type Stream[T any] struct {
	iterator iterable.Iterable[T]
	values   []T
	actions  []func(T) (T, bool)

	locker sync.Mutex
}

func New[T any](iterator iterable.Iterable[T]) *Stream[T] {
	return &Stream[T]{
		iterator: iterator,
		values:   make([]T, 0),
		actions:  make([]func(T) (T, bool), 0),

		locker: sync.Mutex{},
	}
}

func (s *Stream[T]) Filter(cond func(T) bool) *Stream[T] {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.actions = append(s.actions, func(v T) (T, bool) {
		if cond(v) {
			return v, true
		}

		return v, false
	})
	return s
}

func (s *Stream[T]) Map(f func(T) T) *Stream[T] {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.actions = append(s.actions, func(v T) (T, bool) {
		return f(v), true
	})
	return s
}

func (s *Stream[T]) Take(n int) []T {
	s.locker.Lock()
	defer s.locker.Unlock()

	values := make([]T, 0, n)

	count := 0
	for count < n {
		if !s.iterator.Next() {
			break
		}
		values = append(s.values, s.iterator.Value())
		count++
	}

	return values
}

func (s *Stream[T]) Collect() []T {
	s.locker.Lock()
	defer s.locker.Unlock()

	for s.iterator.Next() {
		v, ok := s.iterator.Value(), true
		for _, action := range s.actions {
			v, ok = action(v)
			if !ok {
				continue
			}
		}
		s.values = append(s.values, v)
	}

	return s.values
}
