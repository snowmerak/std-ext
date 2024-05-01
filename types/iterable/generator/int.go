package generator

import "golang.org/x/exp/constraints"

func Int[T constraints.Integer](start, end, step T) *Generator[T] {
	return New(start, func(v T) (T, bool) {
		if step > 0 && v >= end || step < 0 && v <= end {
			return 0, false
		}
		return v + step, true
	})
}
