package generics

func IsEven(n int) bool {
	return n%2 == 0
}

// Find an item in a slice based on a predicate.
func Find[T any](items []T, predicate func(T) bool) (value T, found bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	return
}
