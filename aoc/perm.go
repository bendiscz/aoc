package aoc

// TODO(mbenda) channel?
func Permutations[T any](data []T, fn func([]T)) {
	if len(data) <= 1 {
		fn(data)
		return
	}

	s := make([]T, len(data))
	copy(s, data)
	for i := 0; i < len(data); i++ {
		if i > 0 {
			data[0], data[i] = data[i], data[0]
		}
		Permutations(data[1:], func(_ []T) { fn(data) })
	}
	copy(data, s)
}
