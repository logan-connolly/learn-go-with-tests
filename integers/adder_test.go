package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
