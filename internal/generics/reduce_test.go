package generics

import (
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("reduce sum integers", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		AssertEqual(t, Reduce(items, Combine[int], 0), 15)
	})

	t.Run("reduce concatenate strings", func(t *testing.T) {
		items := []string{"h", "e", "l", "l", "o"}
		AssertEqual(t, Reduce(items, Combine[string], ""), "hello")
	})

	t.Run("reduce multiply integers", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		AssertEqual(t, Reduce(items, Multiply[int], 1), 120)
	})
}
