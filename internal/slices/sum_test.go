package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	want := 15

	got := Sum(numbers)

	if want != got {
		t.Errorf("want %d got %d given, %v", want, got, numbers)
	}
}

func TestSumAll(t *testing.T) {
	want := []int{3, 9}
	got := SumAll([]int{1, 2}, []int{0, 9})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Test arrays with enought elements", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 1})
		want := []int{5, 10}
		checkSums(t, got, want)
	})

	t.Run("Test arrays with zero elements", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 1})
		want := []int{0, 10}
		checkSums(t, got, want)
	})
}
