package generics

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (val T, ok bool) {
	if s.IsEmpty() {
		return val, ok
	}

	index := len(s.values) - 1
	val = s.values[index]
	s.values = s.values[:index]

	return val, true
}
