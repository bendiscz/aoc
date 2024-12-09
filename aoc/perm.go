package aoc

import "iter"

func Permutations[T any](data []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		genPerm[T](data, data, yield)
	}
}

func genPerm[T any](data, base []T, yield func([]T) bool) bool {
	if len(data) <= 1 {
		return yield(base)
	}

	s := make([]T, len(data))
	copy(s, data)
	for i := 0; i < len(data); i++ {
		if i > 0 {
			data[0], data[i] = data[i], data[0]
		}
		if !genPerm[T](data[1:], base, yield) {
			return false
		}
	}
	copy(data, s)
	return true
}
