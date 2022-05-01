package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1,2,3,4,5}
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
