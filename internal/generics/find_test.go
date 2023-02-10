package generics

import (
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find even number in slice", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}

		val, found := Find(items, IsEven)

		AssertEqual(t, val, 2)
		AssertEqual(t, found, true)
	})

	t.Run("no even number found", func(t *testing.T) {
		items := []int{1, 3, 5}

		val, found := Find(items, IsEven)

		AssertEqual(t, val, 0)
		AssertEqual(t, found, false)
	})
}
