package generics

import (
	"testing"

	"github.com/logan-connolly/learn-go-with-tests/internal/utils"
)

func TestFind(t *testing.T) {
	t.Run("find even number in slice", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}

		val, found := Find(items, IsEven)

		utils.AssertEqual(t, val, 2)
		utils.AssertEqual(t, found, true)
	})

	t.Run("no even number found", func(t *testing.T) {
		items := []int{1, 3, 5}

		val, found := Find(items, IsEven)

		utils.AssertEqual(t, val, 0)
		utils.AssertEqual(t, found, false)
	})
}
