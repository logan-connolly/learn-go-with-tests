package generics

import (
	"testing"

	"github.com/logan-connolly/learn-go-with-tests/internal/utils"
)

func TestStack(t *testing.T) {
	t.Run("initiated stack is empty", func(t *testing.T) {
		s := new(Stack[int])

		utils.AssertEqual(t, s.IsEmpty(), true)
	})

	t.Run("integer: empty stack return 'zero' value", func(t *testing.T) {
		s := new(Stack[int])

		got, ok := s.Pop()

		utils.AssertEqual(t, ok, false)
		utils.AssertEqual(t, got, 0)
	})

	t.Run("string: empty stack return 'zero' value", func(t *testing.T) {
		s := new(Stack[string])

		got, ok := s.Pop()

		utils.AssertEqual(t, ok, false)
		utils.AssertEqual(t, got, "")
	})

	t.Run("stack returns last value", func(t *testing.T) {
		s := new(Stack[int])
		s.Push(5)
		s.Push(3)

		got, ok := s.Pop()

		utils.AssertEqual(t, ok, true)
		utils.AssertEqual(t, got, 3)
	})
}
