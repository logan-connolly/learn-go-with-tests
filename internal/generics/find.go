package generics

func IsEven(n int) bool {
	return n%2 == 0
}

// Find an item in a slice based on a predicate.
func Find[T any](items []T, pred func(T) bool) (val T, ok bool) {
	for _, item := range items {
		if pred(item) {
			return item, true
		}
	}
	return
}
