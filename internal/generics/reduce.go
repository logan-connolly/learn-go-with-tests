package generics

type number interface {
	int | float64
}

type combinable interface {
	number | string
}

func Combine[T combinable](x, y T) T {
	return x + y
}

func Multiply[T number](x, y T) T {
	return x * y
}

// Reduce fold an array applying a function.
func Reduce[T any](items []T, fn func(T, T) T, val T) T {
	if len(items) == 1 {
		return fn(val, items[0])
	}

	first, rest := items[0], items[1:]

	return Reduce(rest, fn, fn(val, first))
}
