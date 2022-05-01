package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1,2,3,4,5}
	want := 15

	got := Sum(numbers)

	if want != got {
		t.Errorf("want %d got %d given, %v", want, got, numbers)
	}
}
