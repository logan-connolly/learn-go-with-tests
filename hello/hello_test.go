package hello

import (
	"fmt"
	"testing"
)

func ExampleHello() {
	hello := Hello("Pierre", french)
	fmt.Println(hello)
	// Output: Bonjour, Pierre
}

func ExampleHello_default() {
	hello := Hello("", "")
	fmt.Println(hello)
	// Output: Hello, World
}

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say 'Hello, World' when no name provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say in English (default)", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say in Spanish", func(t *testing.T) {
		got := Hello("Hugo", spanish)
		want := "Hola, Hugo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say in French", func(t *testing.T) {
		got := Hello("Pierre", french)
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})
}
