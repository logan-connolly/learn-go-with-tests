package generics

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("initiated stack is empty", func(t *testing.T) {
		s := new(Stack[int])

		AssertEqual(t, s.IsEmpty(), true)
	})

	t.Run("integer: empty stack return 'zero' value", func(t *testing.T) {
		s := new(Stack[int])

		got, ok := s.Pop()

		AssertEqual(t, ok, false)
		AssertEqual(t, got, 0)
	})

	t.Run("string: empty stack return 'zero' value", func(t *testing.T) {
		s := new(Stack[string])

		got, ok := s.Pop()

		AssertEqual(t, ok, false)
		AssertEqual(t, got, "")
	})

	t.Run("stack returns last value", func(t *testing.T) {
		s := new(Stack[int])
		s.Push(5)
		s.Push(3)

		got, ok := s.Pop()

		AssertEqual(t, ok, true)
		AssertEqual(t, got, 3)
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
