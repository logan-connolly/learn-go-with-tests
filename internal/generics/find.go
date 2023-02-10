package generics

func IsEven(n int) bool {
	return n%2 == 0
}

// Find an item in a slice based on a predicate.
func Find[T any](items []T, pred func(T) bool) (val T, ok bool) {
	if len(items) == 0 {
		return val, ok
	}

	first, rest := items[0], items[1:]

	if pred(first) {
		return first, true
	}

	return Find(rest, pred)
}
